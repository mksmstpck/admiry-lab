package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/mkskstpck/admiry-lab/pkg/utils"
	"github.com/mkskstpck/admiry-lab/services/api-gateway/jwt"
)

func (h *Handlers) logInByEmail(c *gin.Context) {
	var creds *models.AuthByEmail

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info("handlers.LogInByEmail: ", err)
		return
	}

	user, code, err := h.user.UserGetByEmail(creds.Email)
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.LogInByEmail: ", err)
		return
	}

	password, code, err := h.user.UserGetPasswordById(user.ID)
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}

	if ok := utils.CheckPasswordHash(creds.Password, password); ok != true {
		c.JSON(http.StatusBadRequest, models.Message{Message: "invalid password"})
		log.Info("handlers.LogInByEmail: invalid password")
		return
	}

	access, err := jwt.CreateJWT(h.access_exp, h.access_secret, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByEmail: ", err)
		return
	}

	refresh, err := jwt.CreateJWT(h.refresh_exp, h.refresh_secret, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByEmail: ", err)
		return
	}
	c.Header("access", access)
	c.Header("refresh", refresh)

	c.JSON(http.StatusNoContent, nil)
	log.Info("handlers.LogInByEmail: user logged in")
}

func (h *Handlers) logInByUsername(c *gin.Context) {
	var creds *models.AuthByUsername

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info("handlers.LogInByUsername: ", err)
		return
	}

	user, code, err := h.user.UserGetByUsername(creds.Username)
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}

	password, code, err := h.user.UserGetPasswordById(user.ID)
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}

	if ok := utils.CheckPasswordHash(creds.Password, password); ok != true {
		c.JSON(http.StatusBadRequest, models.Message{Message: "invalid password"})
		log.Info("handlers.LogInByUsername: invalid password")
		return
	}

	access, err := jwt.CreateJWT(h.access_exp, h.access_secret, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}

	refresh, err := jwt.CreateJWT(h.refresh_exp, h.refresh_secret, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}
	c.Header("access", access)
	c.Header("refresh", refresh)

	c.JSON(http.StatusNoContent, nil)
	log.Info("handlers.LogInByUsername: user logged in")
}

func (h *Handlers) refresh(c *gin.Context) {
	user_id, err := jwt.ValidateJWT(c.Request.Header.Get("refresh"), h.refresh_secret)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.Refresh: ", err)
		return
	}
	user, code, err := h.user.UserGetById(user_id)
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.Refresh: ", err)
		return
	}

	access, err := jwt.CreateJWT(h.access_exp, h.access_secret, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}

	refresh, err := jwt.CreateJWT(h.refresh_exp, h.refresh_secret, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}
	c.Header("access", access)
	c.Header("refresh", refresh)

	c.JSON(http.StatusNoContent, nil)
	log.Info("handlers.Refresh: user logged in")
}
