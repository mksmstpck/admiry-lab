package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/pborman/uuid"
)

func (h *Handlers) roleReadById(c *gin.Context) {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		log.Info("handlers.roleReadById: role not found")
		c.JSON(http.StatusNotFound, models.Message{Message: "role not found"})
		return
	}
	r, code, err := h.role.RoleGetById(UUID)
	if err != nil {
		log.Error("handlers.roleReadById: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.roleReadById: role found")
	c.JSON(http.StatusOK, r)
}

func (h *Handlers) roleReadByName(c *gin.Context) {
	name := c.Param("name")
	r, code, err := h.role.RoleGetByName(name)
	if err != nil {
		log.Error("handlers.roleReadByName: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.roleReadByName: role found")
	c.JSON(http.StatusOK, r)
}

func (h *Handlers) roleCreate(c *gin.Context) {
	role := new(models.Role)
	if err := c.ShouldBind(&role); err != nil {
		log.Error("handlers.roleCreate: ", err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		return
	}
	if role.PermissionIDs == nil {
		role.PermissionIDs = []uuid.UUID{uuid.NIL}
	}
	r, code, err := h.role.RolePost(role)
	if err != nil {
		log.Error("handlers.roleCreate: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.roleCreate: role created")
	c.JSON(http.StatusOK, r)
}

func (h *Handlers) roleUpdate(c *gin.Context) {
	role := new(models.Role)
	if err := c.ShouldBind(&role); err != nil {
		log.Error("handlers.roleUpdate: ", err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		return
	}
	code, err := h.role.RolePut(role)
	if err != nil {
		log.Error("handlers.roleUpdate: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.roleUpdate: role updated")
	c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) roleDelete(c *gin.Context) {
	UUID := uuid.Parse(c.Param("id"))
	code, err := h.role.RoleDelete(UUID)
	if err != nil {
		log.Error("handlers.roleDelete: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.roleDelete: role deleted")
	c.JSON(http.StatusNoContent, nil)
}
