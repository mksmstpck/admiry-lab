package handlers

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/mkskstpck/to-rename/pkg/utils"
	"github.com/pborman/uuid"
)

func (h *Handler) companyReadById() {
	_, err := h.conn.Subscribe("companies-id-get", func(_, reply string, id uuid.UUID) {
		company, code, err := h.cache.GetCompany(id.String(), context.Background())
		if err == nil && company != nil {
			res := models.Response[models.Company]{Status: code, Message: company.(models.Company)}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: company found by id")
		}
		company, code, err = h.company.CompanyFindOneById(id)
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
		}
		if company.(models.Company).ID == nil {
			res := models.Response[models.Company]{Status: 404, Error: "company not found"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
		}
		code, err = h.cache.Set(id.String(), company, context.Background())
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
		}
		res := models.Response[models.Company]{Status: 200, Message: company.(models.Company)}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: company found by id")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) companyReadByName() {
	_, err := h.conn.Subscribe("companies-name-get", func(_, reply string, name string) {
		company, code, err := h.cache.GetCompany(name, context.Background())
		if err == nil && company != nil {
			res := models.Response[models.Company]{Status: code, Message: company.(models.Company)}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: company found by name")
		}
		company, code, err = h.company.CompanyFindOneByName(name)
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		if company.(models.Company).ID == nil {
			res := models.Response[models.Company]{Status: 404, Error: "company not found"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: company not found by name")
		}
		code, err = h.cache.Set(name, company, context.Background())
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		res := models.Response[models.Company]{Status: 200, Message: company.(models.Company)}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: company found by name")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}
func (h *Handler) companyReadAll() {
	_, err := h.conn.Subscribe("companies-get", func(_, reply string, id uuid.UUID) {
		company, code, err := h.company.CompanyFindAll()
		if err != nil && company != nil {
			res := models.Response[[]models.Company]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: companies found")
		}
		if company == nil {
			res := models.Response[[]models.Company]{Status: 404, Error: "company not found"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: companies not found")
		}
		res := models.Response[[]models.Company]{Status: 200, Message: company}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: companies found")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) companyCreate() {
	_, err := h.conn.Subscribe("companies-post", func(_, reply string, company models.Company) {
		companyExist, code, err := h.company.CompanyFindOneByName(company.Name)
		if err != nil {
			if err.Error() != "company not found" {
				res := models.Response[models.Company]{Status: code, Error: err.Error()}
				utils.NatsPublishError(h.conn.Publish(reply, res))
				log.Error("handlers: ", err)
			}
		}
		if companyExist.ID != nil {
			res := models.Response[models.Company]{Status: 409, Error: "company with this name already exists"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: company with this name already exists")
			return
		}
		code, err = h.company.CompanyCreateOne(company)
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		company, code, err = h.company.CompanyFindOneByName(company.Name)
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		code, err = h.cache.Set(company.ID.String(), company, context.Background())
		if err != nil {
			res := models.Response[models.Company]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		res := models.Response[models.Company]{Status: 201, Message: company}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: company created")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) companyUpdate() {
	_, err := h.conn.Subscribe("companies-put", func(_, reply string, company models.Company) {
		code, err := h.company.CompanyUpdateOne(company)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		code, err = h.cache.Set(company.ID.String(), company, context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		res := models.Response[string]{Status: 204, Message: "company updated"}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: company updated")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) companyDelete() {
	_, err := h.conn.Subscribe("companies-delete", func(_, reply string, id uuid.UUID) {
		code, err := h.company.CompanyDeleteOne(id)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		code, err = h.cache.Delete(id.String(), context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		res := models.Response[string]{Status: code, Message: "company deleted"}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: company deleted")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}
