package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/pborman/uuid"
)

func (h *Handlers) permReadById(c *gin.Context) {
	UUID := uuid.Parse(c.Param("id"))
	if UUID == nil {
		c.JSON(http.StatusNotFound, models.Message{Message: "role not found"})
		log.Info("handlers.permReadById: role not found")
		return
	}
	p, code, err := h.perm.PermGetById(UUID)
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.permReadById: ", err)
		return
	}
	c.JSON(http.StatusOK, p)
	log.Info("handlers.permReadById: perm found")
}

func (h *Handlers) permReadByName(c *gin.Context) {
	p, code, err := h.perm.PermGetByName(c.Param("name"))
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.permReadByName: ", err)
		return
	}
	c.JSON(http.StatusOK, p)
	log.Info("handlers.permReadByName: perm found")
}

func (h *Handlers) permReadAll(c *gin.Context) {
	p, code, err := h.perm.PermGetAll()
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.permReadAll: ", err)
		return
	}
	c.JSON(http.StatusOK, p)
	log.Info("handlers.permReadAll: perm found")
}

func (h *Handlers) permCreate(c *gin.Context) {
	perm := new(models.Permission)
	if err := c.ShouldBind(&perm); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.roleCreate: ", err)
		return
	}
	p, code, err := h.perm.PermPost(perm)
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.permCreate: ", err)
		return
	}
	c.JSON(http.StatusOK, p)
	log.Info("handlers.permCreate: perm created")
}

func (h *Handlers) permUpdate(c *gin.Context) {
	perm := new(models.Permission)
	if err := c.ShouldBind(&perm); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.permUpdate: ", err)
		return
	}
	code, err := h.perm.PermPut(perm)
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.permUpdate: ", err)
		return
	}
	c.JSON(http.StatusNoContent, models.Message{Message: "perm updated"})
	log.Info("handlers.permUpdate: perm updated")
}

func (h *Handlers) permDelete(c *gin.Context) {
	UUID := uuid.Parse(c.Param("id"))
	if UUID == nil {
		c.JSON(http.StatusNotFound, models.Message{Message: "role not found"})
		log.Info("handlers.permDelete: role not found")
		return
	}
	code, err := h.perm.PermDelete(UUID)
	if err != nil {
		c.JSON(int(code), models.Message{Message: err.Error()})
		log.Error("handlers.permDelete: ", err)
		return
	}
	c.JSON(http.StatusNoContent, models.Message{Message: "perm deleted"})
	log.Info("handlers.permDelete: perm deleted")
}
