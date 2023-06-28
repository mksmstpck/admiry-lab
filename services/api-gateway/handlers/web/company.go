package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (h *Handlers) companyReadById(c echo.Context) error {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: "company not found"})
	}
	co, code, err := h.company.CompanyGetById(UUID)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, co)
}

func (h *Handlers) companyReadByName(c echo.Context) error {
	name := c.Param("name")
	co, code, err := h.company.CompanyGetByName(name)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, co)
}

func (h *Handlers) companyReadAll(c echo.Context) error {
	u, code, err := h.company.CompanyGetAll()
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) companyCreate(c echo.Context) error {
	co := new(models.Company)
	if err := c.Bind(co); err != nil {
		return err
	}
	if err := c.Validate(co); err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	if co.UserIDs == nil {
		co.UserIDs = []uuid.UUID{uuid.NIL}
	}
	company, code, err := h.company.CompanyPost(co)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, company)
}

func (h *Handlers) companyUpdate(c echo.Context) error {
	co := new(models.Company)
	if err := c.Bind(co); err != nil {
		return err
	}
	if err := c.Validate(co); err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	code, err := h.company.CompanyPut(co)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, "updated")
}

func (h *Handlers) companyDelete(c echo.Context) error {
	id := c.Param("id")
	UUID := uuid.Parse(id)
	if UUID == nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: "company not found"})
	}
	code, err := h.company.CompanyDelete(UUID)
	if err != nil {
		return c.JSON(int(code), models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, "deleted")
}
