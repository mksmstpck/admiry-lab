package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mkskstpck/to-rename/pkg/models"
)

func (h *Handlers) userReadById(c echo.Context) error {
	id := c.Param("id")
	uuID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	u, code, err := h.user.UserGetById(uuID)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) userReadByUsername(c echo.Context) error {
	username := c.Param("username")
	u, code, err := h.user.UserGetByUsername(username)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) userReadByEmail(c echo.Context) error {
	email := c.Param("email")
	u, code, err := h.user.UserGetByEmail(email)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) userCreate(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	user, code, err := h.user.UserPost(u)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *Handlers) userUpdate(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	code, err := h.user.UserPut(u)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})

	}
	return c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) userDelete(c echo.Context) error {
	id := c.Param("id")
	uuID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	code, err := h.user.UserDelete(uuID)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}
