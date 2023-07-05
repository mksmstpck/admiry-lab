package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (h *Handlers) userReadById(c *gin.Context) {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		log.Info("handlers.userReadById: user not found")
		c.JSON(http.StatusNotFound, models.Message{Message: "user not found"})
		return
	}
	u, code, err := h.user.UserGetById(UUID)
	if err != nil {
		log.Error("handlers.userReadById: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.userReadById: user found")
	c.JSON(http.StatusOK, u)
}

func (h *Handlers) userReadByUsername(c *gin.Context) {
	username := c.Param("username")
	u, code, err := h.user.UserGetByUsername(username)
	if err != nil {
		log.Error("handlers.userReadByUsername: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.userReadByUsername: user found")
	c.JSON(http.StatusOK, u)
	return
}

func (h *Handlers) userReadByEmail(c *gin.Context) {
	email := c.Param("email")
	u, code, err := h.user.UserGetByEmail(email)
	if err != nil {
		log.Error("handlers.userReadByEmail: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.userReadByEmail: user found")
	c.JSON(http.StatusOK, u)
}

func (h *Handlers) userCreate(c *gin.Context) {
	u := new(models.User)
	if err := c.ShouldBind(u); err != nil {
		log.Error("handlers.userCreate: ", err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		return
	}
	user, code, err := h.user.UserPost(u)
	if err != nil {
		log.Error("handlers.userCreate: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.userCreate: user created")
	c.JSON(http.StatusCreated, user)
}

func (h *Handlers) userUpdate(c *gin.Context) {
	u := new(models.User)
	if err := c.ShouldBind(u); err != nil {
		log.Error("handlers.userUpdate: ", err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	code, err := h.user.UserPut(u)
	if err != nil {
		log.Error("handlers.userUpdate: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.userUpdate: user updated")
	c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) userDelete(c *gin.Context) {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		log.Info("handlers.userDelete: user not found")
		c.JSON(http.StatusNotFound, models.Message{Message: "User not found"})
		return
	}
	code, err := h.user.UserDelete(UUID)
	if err != nil {
		log.Error("handlers.userDelete: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.userDelete: user deleted")
	c.JSON(http.StatusNoContent, nil)
}
