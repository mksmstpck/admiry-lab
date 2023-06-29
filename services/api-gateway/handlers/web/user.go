package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (h *Handlers) userReadById(c echo.Context) error {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		log.Info("handlers.userReadById: user not found")
		return c.JSON(http.StatusNotFound, models.Message{Message: "user not found"})
	}
	u, code, err := h.user.UserGetById(UUID)
	if err != nil {
		log.Error("handlers.userReadById: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.userReadById: user found")
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) userReadByUsername(c echo.Context) error {
	username := c.Param("username")
	u, code, err := h.user.UserGetByUsername(username)
	if err != nil {
		log.Error("handlers.userReadByUsername: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.userReadByUsername: user found")
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) userReadByEmail(c echo.Context) error {
	email := c.Param("email")
	u, code, err := h.user.UserGetByEmail(email)
	if err != nil {
		log.Error("handlers.userReadByEmail: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.userReadByEmail: user found")
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) userCreate(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		log.Error("handlers.userCreate: ", err)
		return err
	}
	if err := c.Validate(u); err != nil {
		log.Error("handlers.userCreate: ", err)
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	user, code, err := h.user.UserPost(u)
	if err != nil {
		log.Error("handlers.userCreate: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.userCreate: user created")
	return c.JSON(http.StatusCreated, user)
}

func (h *Handlers) userUpdate(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		log.Error("handlers.userUpdate: ", err)
		return err
	}
	if err := c.Validate(u); err != nil {
		log.Error("handlers.userUpdate: ", err)
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	code, err := h.user.UserPut(u)
	if err != nil {
		log.Error("handlers.userUpdate: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})

	}
	log.Info("handlers.userUpdate: user updated")
	return c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) userDelete(c echo.Context) error {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		log.Info("handlers.userDelete: user not found")
		return c.JSON(http.StatusNotFound, models.Message{Message: "User not found"})
	}
	code, err := h.user.UserDelete(UUID)
	if err != nil {
		log.Error("handlers.userDelete: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.userDelete: user deleted")
	return c.JSON(http.StatusNoContent, nil)
}
