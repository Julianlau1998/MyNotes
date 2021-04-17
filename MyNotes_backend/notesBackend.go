package main

import (
	"net/http"
	"notesBackend/lists"
	"notesBackend/notes"
	"notesBackend/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

func main() {
	NoteRepository := notes.NewRepository()
	NoteService := notes.NewService(NoteRepository)
	NoteDelivery := notes.NewDelivery(NoteService)

	ListRepository := lists.NewRepository()
	ListService := lists.NewService(ListRepository)
	ListDelivery := lists.NewDelivery(ListService)

	UserRepository := users.NewRepository()
	UserService := users.NewService(UserRepository)
	UserDelivery := users.NewDelivery(UserService)

	e := echo.New()

	e.Use(CORSMiddlewareWrapper)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/api/notes", NoteDelivery.GetAll)
	e.GET("/api/note/:id", NoteDelivery.GetById)
	e.POST("/api/notes", NoteDelivery.Post)
	e.PUT("/api/note/:id", NoteDelivery.Update)
	e.DELETE("/api/note/:id", NoteDelivery.Delete)

	e.GET("/api/lists", ListDelivery.GetAll)
	e.GET("/api/list/:id", ListDelivery.GetById)
	e.POST("/api/lists", ListDelivery.Post)
	e.PUT("/api/list/:id", ListDelivery.Update)
	e.DELETE("/api/list/:id", ListDelivery.Delete)

	e.GET("/api/users", UserDelivery.GetAll)
	e.POST("/api/users", UserDelivery.Post)

	e.Logger.Fatal(e.Start(":1323"))
}
