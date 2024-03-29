package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"notesBackend/folders"
	"notesBackend/list_elements"
	"notesBackend/lists"
	"notesBackend/models"
	"notesBackend/notes"
	"notesBackend/utility"
	"os"
	"strings"
	"time"

	auth0 "github.com/b-venter/auth0-go-jwt-middleware"
	jwt "github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var aud = os.Getenv("AUTH0_AUD") //As set in Auth0 API > Settings > Identifier

var iss = os.Getenv("AUTH0_ISS") //Your Auth0 tenant domain

var jwksUrl = os.Getenv("AUTH0_JWKS") //https://YOUR-AUTH0-TENANT-DOMAIN/.well-known/jwks.json

/* end */

var jwtClaims []string

type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

//getPemCert: function to get certificate in PEM format.
//Used by verify() func
func getPemCert(token *jwt.Token) (string, error) {
	godotenv.Load(".env")
	jwksUrl := os.Getenv("AUTH0_JWKS") //https://YOUR-AUTH0-TENANT-DOMAIN/.well-known/jwks.json

	cert := ""

	resp, err := http.Get(jwksUrl)

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

//A jwt.Keyfunc type function, required for JWTMiddleware struct
//It processes the aud, iss validity of the jwt.Token. and gets certificate which CheckJWT will parse and verify.
func verify(token *jwt.Token) (interface{}, error) {
	godotenv.Load(".env")
	aud := os.Getenv("AUTH0_AUD")

	// Verify 'aud' claim
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	if !checkAud {
		return token, errors.New("invalid audience")
	}

	godotenv.Load(".env")
	iss := os.Getenv("AUTH0_ISS") //Your Auth0 tenant domain

	// Verify 'iss' claim
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
	if !checkIss {
		return token, errors.New("invalid issuer")
	}

	//Store "scope" so long. Use it or lose it later.
	claim := fmt.Sprintf("%v", token.Claims.(jwt.MapClaims)["scope"])
	jwtClaims = strings.Fields(string(claim))

	cert, err := getPemCert(token)
	if err != nil {
		//panic(err.Error())
		return nil, err
	}

	result, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	if err != nil {
		return nil, err
	}

	return result, nil

}

//Type for unpacking json
// type ujson map[string]interface{}

//Function to call Auth0 /userinfo API endpoint
func getUser(tok string) models.User {
	godotenv.Load(".env")
	iss := os.Getenv("AUTH0_ISS") //Your Auth0 tenant domain

	url := iss + "userinfo"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", tok)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//Get as json
	var bodyj models.User
	if err := json.Unmarshal(body, &bodyj); err != nil {
		fmt.Println("Error providing json: ", err)
	}

	return bodyj
}

/* MIDDLEWARE */

//Custom middleware

//Middleware that verifies token
func middleJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// c holds Response and Request
		// t has the methods to process the token
		var t *auth0.JWTMiddleware = auth0.New(auth0.Options{
			ValidationKeyGetter: verify,
			SigningMethod:       jwt.SigningMethodRS256,
		})
		vr := t.CheckJWT(c.Response().Writer, c.Request())
		if vr != nil {
			return echo.ErrUnauthorized //Not allowed to proceed. Might be nice to break it down further to "no token", etc
		}
		token, _, _ := new(jwt.Parser).ParseUnverified(c.Request().Context().Value("user").(*jwt.Token).Raw, jwt.MapClaims{})
		newRequest := c.Request().WithContext(context.WithValue(c.Request().Context(), "currentUser", token.Claims.(jwt.MapClaims)))
		*c.Request() = *newRequest

		return next(c) //Proceed to next.
	}
}

func middleUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Retrieve user info
		//token := c.Request().Header.Get("Authorization")
		//user := getUser(token)
		//fmt.Println("User info: ", userInfo)

		//newRequest := c.Request().WithContext(context.WithValue(c.Request().Context(), "user", user))
		//*c.Request() = *newRequest

		//TODO: Replace this with a database lookup
		// if userInfo["email"] != "my-email@gmail.com" {
		// 	return echo.ErrForbidden //Note: https://stackoverflow.com/questions/3297048/403-forbidden-vs-401-unauthorized-http-responses
		// }

		return next(c) //Proceed to next.
	}
}

func CORSMiddlewareWrapper(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()
		dynamicCORSConfig := middleware.CORSConfig{
			AllowOrigins: []string{req.Header.Get("Origin")},
			AllowHeaders: []string{"*"},
		}
		CORSMiddleware := middleware.CORSWithConfig(dynamicCORSConfig)
		CORSHandler := CORSMiddleware(next)
		return CORSHandler(ctx)
	}
}

var dbClient *sql.DB

func startup() {

	dbClient = utility.NewDbClient()
	for utility.Migrate(dbClient) != nil {
		fmt.Printf("Verbindung Fehlgeschlagen, %s", utility.Migrate(dbClient))
		time.Sleep(20 * time.Second)
	}
}

func main() {

	startup()
	NoteRepository := notes.NewRepository(dbClient)
	NoteService := notes.NewService(NoteRepository)
	NoteDelivery := notes.NewDelivery(NoteService)

	ListElementsRepository := list_elements.NewRepository(dbClient)
	ListElementsService := list_elements.NewService(ListElementsRepository)

	ListRepository := lists.NewRepository(dbClient)
	ListService := lists.NewService(ListRepository, ListElementsService)
	ListDelivery := lists.NewDelivery(ListService)

	FolderRepository := folders.NewRepository(dbClient)
	FolderService := folders.NewService(FolderRepository, ListService, NoteService)
	FolderDelivery := folders.NewDelivery(FolderService)

	// UserRepository := users.NewRepository(dbClient)
	// UserService := users.NewService(UserRepository)
	// UserDelivery := users.NewDelivery(UserService)

	e := echo.New()

	//e.Use(CORSMiddlewareWrapper)

	r := e.Group("/api")
	r.Use(middleJWT)

	r.GET("/notes", NoteDelivery.GetAll, middleUser)
	r.GET("/note/:id", NoteDelivery.GetById, middleUser)
	r.GET("/notes/foldernotes", NoteDelivery.GetByFolder, middleUser)
	r.POST("/notes", NoteDelivery.Post, middleUser)
	r.PUT("/note/:id", NoteDelivery.Update, middleUser)
	r.DELETE("/note/:id", NoteDelivery.Delete, middleUser)

	r.GET("/lists", ListDelivery.GetAll, middleUser)
	r.GET("/list/:id", ListDelivery.GetById, middleUser)
	r.GET("/lists/folderlists", ListDelivery.GetByFolder, middleUser)
	r.POST("/lists", ListDelivery.Post, middleUser)
	r.PUT("/list/:id", ListDelivery.Update, middleUser)
	r.DELETE("/list/:id", ListDelivery.Delete, middleUser)

	r.GET("/folders", FolderDelivery.GetAll, middleUser)
	r.GET("/folder/:id", FolderDelivery.GetById, middleUser)
	r.POST("/folders", FolderDelivery.Post, middleUser)
	r.PUT("/folder/:id", FolderDelivery.Update, middleUser)
	r.DELETE("/folder/:id", FolderDelivery.Delete, middleUser)

	// e.GET("/api/users", UserDelivery.GetAll)
	// e.POST("/api/users", UserDelivery.Post)

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "1323"
	// }
	port := "8081"
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
