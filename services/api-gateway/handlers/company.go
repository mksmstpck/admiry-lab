package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/pborman/uuid"
)

func (h *Handlers) companyReadById(c *gin.Context) {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		log.Info("handlers.companyReadById: company not found")
		c.JSON(http.StatusNotFound, models.Message{Message: "company not found"})
		return
	}
	co, code, err := h.company.CompanyGetById(UUID)
	if err != nil {
		log.Error("handlers.companyReadById: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.companyReadById: company found")
	c.JSON(http.StatusOK, co)
}

func (h *Handlers) companyReadByName(c *gin.Context) {
	name := c.Param("name")
	co, code, err := h.company.CompanyGetByName(name)
	if err != nil {
		log.Error("handlers.companyReadByName: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.companyReadByName: company found")
	c.JSON(http.StatusOK, co)
}

func (h *Handlers) companyReadAll(c *gin.Context) {
	u, code, err := h.company.CompanyGetAll()
	if err != nil {
		log.Error("handlers.companyReadAll: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.companyReadAll: company found")
	c.JSON(http.StatusOK, u)
}

func (h *Handlers) companyCreate(c *gin.Context) {
	co := new(models.Company)
	if err := c.ShouldBind(&co); err != nil {
		log.Error("handlers.companyCreate: ", err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		return
	}
	if co.UserIDs == nil {
		co.UserIDs = []uuid.UUID{uuid.NIL}
	}
	company, code, err := h.company.CompanyPost(co)
	if err != nil {
		log.Error("handlers.companyCreate: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.companyCreate: company created")
	c.JSON(http.StatusCreated, company)
}

func (h *Handlers) companyUpdate(c *gin.Context) {
	co := new(models.Company)
	if err := c.ShouldBind(&co); err != nil {
		log.Error("handlers.companyUpdate: ", err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	code, err := h.company.CompanyPut(co)
	if err != nil {
		log.Error("handlers.companyUpdate: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.companyUpdate: company updated")
	c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) companyDelete(c *gin.Context) {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		log.Info("handlers.companyDelete: company not found")
		c.JSON(http.StatusNotFound, models.Message{Message: "company not found"})
		return
	}
	code, err := h.company.CompanyDelete(UUID)
	if err != nil {
		log.Error("handlers.companyDelete: ", err)
		c.JSON(int(code), models.Message{Message: err.Error()})
		return
	}
	log.Info("handlers.companyDelete: company deleted")
	c.JSON(http.StatusNoContent, nil)
}
