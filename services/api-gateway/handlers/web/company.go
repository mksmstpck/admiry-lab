package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (h *Handlers) companyReadById(c echo.Context) error {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		log.Info("handlers.companyReadById: company not found")
		return c.JSON(http.StatusNotFound, models.Message{Message: "company not found"})
	}
	co, code, err := h.company.CompanyGetById(UUID)
	if err != nil {
		log.Error("handlers.companyReadById: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.companyReadById: company found")
	return c.JSON(http.StatusOK, co)
}

func (h *Handlers) companyReadByName(c echo.Context) error {
	name := c.Param("name")
	co, code, err := h.company.CompanyGetByName(name)
	if err != nil {
		log.Error("handlers.companyReadByName: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.companyReadByName: company found")
	return c.JSON(http.StatusOK, co)
}

func (h *Handlers) companyReadAll(c echo.Context) error {
	u, code, err := h.company.CompanyGetAll()
	if err != nil {
		log.Error("handlers.companyReadAll: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.companyReadAll: company found")
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) companyCreate(c echo.Context) error {
	co := new(models.Company)
	if err := c.Bind(co); err != nil {
		log.Error("handlers.companyCreate: ", err)
		return err
	}
	if err := c.Validate(co); err != nil {
		log.Error("handlers.companyCreate: ", err)
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	if co.UserIDs == nil {
		co.UserIDs = []uuid.UUID{uuid.NIL}
	}
	company, code, err := h.company.CompanyPost(co)
	if err != nil {
		log.Error("handlers.companyCreate: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.companyCreate: company created")
	return c.JSON(http.StatusCreated, company)
}

func (h *Handlers) companyUpdate(c echo.Context) error {
	co := new(models.Company)
	if err := c.Bind(co); err != nil {
		log.Error("handlers.companyUpdate: ", err)
		return err
	}
	if err := c.Validate(co); err != nil {
		log.Error("handlers.companyUpdate: ", err)
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	code, err := h.company.CompanyPut(co)
	if err != nil {
		log.Error("handlers.companyUpdate: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.companyUpdate: company updated")
	return c.JSON(http.StatusNoContent, "updated")
}

func (h *Handlers) companyDelete(c echo.Context) error {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		log.Info("handlers.companyDelete: company not found")
		return c.JSON(http.StatusNotFound, models.Message{Message: "company not found"})
	}
	code, err := h.company.CompanyDelete(UUID)
	if err != nil {
		log.Error("handlers.companyDelete: ", err)
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	log.Info("handlers.companyDelete: company deleted")
	return c.JSON(http.StatusNoContent, "deleted")
}
