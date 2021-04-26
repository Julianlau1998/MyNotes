package lists

import (
	"fmt"
	"net/http"
	"notesBackend/models"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	listService Service
}

func NewDelivery(listService Service) Delivery {
	return Delivery{listService: listService}
}

func (d *Delivery) GetAll(c echo.Context) error {
	userID := c.Request().Header.Get("userId")
	lists, err := d.listService.GetLists(userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, lists)
}

func (d *Delivery) GetById(c echo.Context) error {
	userID := c.Request().Header.Get("userID")
	id := c.Param("id")
	list, err := d.listService.GetListById(id, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if list.ID == "" {
		return c.String(http.StatusForbidden, "Not Authorized")
	}
	return c.JSON(http.StatusOK, list)
}

func (d *Delivery) GetByFolder(c echo.Context) error {
	userID := c.Request().Header.Get("userID")
	folderID := c.Request().Header.Get("folderID")
	lists, err := d.listService.GetByFolder(folderID, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, lists)
}

func (d *Delivery) Post(c echo.Context) error {
	requestBody := new(models.List)
	userID := c.Request().Header.Get("userID")
	if err := c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	list, err := d.listService.Post(requestBody, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, list.ID)
}

func (d *Delivery) Update(c echo.Context) (err error) {
	ID := c.Param("id")
	userID := c.Request().Header.Get("userID")
	requestBody := new(models.List)
	if err = c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	list, err := d.listService.UpdateList(ID, requestBody, userID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	if list.ID == "" {
		return c.String(http.StatusForbidden, "Not Authorized")
	}
	return c.JSON(http.StatusOK, list)
}

func (d *Delivery) Delete(c echo.Context) (err error) {
	id := c.Param("id")
	userID := c.Request().Header.Get("userID")
	list, err := d.listService.DeleteList(id, userID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	if list.ID == "unauthorized" {
		return c.String(http.StatusForbidden, "Not Authorized")
	}
	return c.JSON(http.StatusOK, list)
}
