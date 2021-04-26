package folders

import (
	"fmt"
	"net/http"
	"notesBackend/models"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	folderService Service
}

func NewDelivery(folderService Service) Delivery {
	return Delivery{folderService: folderService}
}

func (d *Delivery) GetAll(c echo.Context) error {
	userID := c.Request().Header.Get("userId")
	folders, err := d.folderService.GetFolders(userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, folders)
}

func (d *Delivery) GetById(c echo.Context) error {
	userID := c.Request().Header.Get("userID")
	id := c.Param("id")
	folder, err := d.folderService.GetById(id, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if folder.ID == "" {
		return c.String(http.StatusForbidden, "Not Authorized")
	}
	return c.JSON(http.StatusOK, folder)
}

func (d *Delivery) Post(c echo.Context) error {
	requestBody := new(models.Folder)
	userID := c.Request().Header.Get("userID")
	if err := c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	folder, err := d.folderService.Post(requestBody, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, folder.ID)
}

func (d *Delivery) Update(c echo.Context) (err error) {
	ID := c.Param("id")
	userID := c.Request().Header.Get("userID")
	requestBody := new(models.Folder)
	if err = c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	note, err := d.folderService.Update(ID, requestBody, userID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	if note.ID == "" {
		return c.String(http.StatusForbidden, "Not Authorized")
	}
	return c.JSON(http.StatusOK, note)
}

func (d *Delivery) Delete(c echo.Context) (err error) {
	id := c.Param("id")
	userID := c.Request().Header.Get("userID")
	folder, err := d.folderService.Delete(id, userID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	if folder.ID == "unauthorized" {
		return c.String(http.StatusForbidden, "Not Authorized")
	}
	return c.JSON(http.StatusOK, folder)
}
