package handlers

import (
	"context"

	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (h *Handler) companyReadById() {
	h.conn.Subscribe("companies-id-get", func(_, reply string, id uuid.UUID) {
		company, code, err := h.cache.GetCompany(id.String(), context.Background())
		if err == nil && company != nil {
			res := models.Response[models.Company]{Status: code, Message: company.(models.Company)}
			h.conn.Publish(reply, res)
		}
		company, code, err = h.company.CompanyFindOneById(id)
		if err != nil && company != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		if company.(models.Company).ID == nil {
			res := models.Response[models.Company]{Status: 404, Error: "company not found"}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Set(id.String(), company, context.Background())
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[models.Company]{Status: 200, Message: company.(models.Company)}
		h.conn.Publish(reply, res)
	})
}

func (h *Handler) companyReadByName() {
	h.conn.Subscribe("companies-name-get", func(_, reply string, name string) {
		company, code, err := h.cache.GetCompany(name, context.Background())
		if err == nil && company != nil {
			res := models.Response[models.Company]{Status: code, Message: company.(models.Company)}
			h.conn.Publish(reply, res)
		}
		company, code, err = h.company.CompanyFindOneByName(name)
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		if company.(models.Company).ID == nil {
			res := models.Response[models.Company]{Status: 404, Error: "company not found"}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Set(name, company, context.Background())
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[models.Company]{Status: 200, Message: company.(models.Company)}
		h.conn.Publish(reply, res)
	})
}
func (h *Handler) companyReadAll() {
	h.conn.Subscribe("companies-get", func(_, reply string, id uuid.UUID) {
		company, code, err := h.company.CompanyFindAll()
		if err != nil && company != nil {
			res := models.Response[[]models.Company]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		if company == nil {
			res := models.Response[[]models.Company]{Status: 404, Error: "company not found"}
			h.conn.Publish(reply, res)
		}
		res := models.Response[[]models.Company]{Status: 200, Message: company}
		h.conn.Publish(reply, res)
	})
}

func (h *Handler) companyCreate() {
	h.conn.Subscribe("companies-post", func(_, reply string, company models.Company) {
		companyExist, code, err := h.company.CompanyFindOneByName(company.Name)
		if err != nil {
			if err.Error() != "company not found" {
				res := models.Response[models.Company]{Status: code, Error: err.Error()}
				h.conn.Publish(reply, res)
			}
		}
		if companyExist.ID != nil {
			res := models.Response[models.Company]{Status: 409, Error: "company with this name already exists"}
			h.conn.Publish(reply, res)
			return
		}
		code, err = h.company.CompanyCreateOne(company)
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		company, code, err = h.company.CompanyFindOneByName(company.Name)
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Set(company.ID.String(), company, context.Background())
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[models.Company]{Status: 201, Message: company}
		h.conn.Publish(reply, res)
	})
}

func (h *Handler) companyUpdate() {
	h.conn.Subscribe("companies-update", func(_, reply string, company models.Company) {
		code, err := h.company.CompanyUpdateOne(company)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Set(company.ID.String(), company, context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[string]{Status: 204, Message: "company updated"}
		h.conn.Publish(reply, res)
	})
}

func (h *Handler) companyDelete() {
	h.conn.Subscribe("companies-delete", func(_, reply string, id uuid.UUID) {
		code, err := h.company.CompanyDeleteOne(id)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Delete(id.String(), context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[string]{Status: code, Message: "company deleted"}
		h.conn.Publish(reply, res)
	})
}
