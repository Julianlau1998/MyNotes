package notes

import (
	"fmt"
	"net/http"
	"notesBackend/models"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	noteService Service
}

func NewDelivery(noteService Service) Delivery {
	return Delivery{noteService: noteService}
}

func (d *Delivery) GetAll(c echo.Context) error {
	fmt.Print("test")
	userID := c.Request().Header.Get("userId")
	notes, err := d.noteService.GetNotes(userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, notes)
}

func (d *Delivery) GetById(c echo.Context) error {
	userID := c.Request().Header.Get("userID")
	id := c.Param("id")
	note, err := d.noteService.GetNoteById(id, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if note.ID == "" {
		return c.String(http.StatusForbidden, "Not Authorized")
	}
	return c.JSON(http.StatusOK, note)
}

func (d *Delivery) Post(c echo.Context) error {
	requestBody := new(models.Note)
	userID := c.Request().Header.Get("userID")
	if err := c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	note, err := d.noteService.Post(requestBody, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, note.ID)
}

func (d *Delivery) Update(c echo.Context) (err error) {
	ID := c.Param("id")
	userID := c.Request().Header.Get("userID")
	requestBody := new(models.Note)
	if err = c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	note, err := d.noteService.updateNote(ID, requestBody, userID)
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
	note, err := d.noteService.deleteNote(id, userID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	if note.ID == "unauthorized" {
		return c.String(http.StatusForbidden, "Not Authorized")
	}
	return c.JSON(http.StatusOK, note)
}
