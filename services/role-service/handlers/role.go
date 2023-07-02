package handlers

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (h *Handler) roleReadById() {
	_, err := h.conn.Subscribe("roles-get-by-id", func(_, reply string, id uuid.UUID) {
		role, code, err := h.cache.GetRole(id.String(), context.Background())
		if err == nil && role != nil {
			res := models.Response[models.Role]{Status: code, Message: role.(models.Role)}
			h.conn.Publish(reply, res)
			log.Info("handlers: company found by id")
		}
		role, code, err = h.role.RoleFindOneById(id)
		if err != nil {
			res := models.Response[models.Role]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
			log.Error("handlers: ", err)
		}
		if role.(models.Role).ID == nil {
			res := models.Response[models.Role]{Status: 404, Error: "role not found"}
			h.conn.Publish(reply, res)
			log.Info("handlers: ", err)
		}
		code, err = h.cache.Set(id.String(), role, context.Background())
		if err != nil {
			res := models.Response[models.Role]{Status: 200, Error: err.Error()}
			h.conn.Publish(reply, res)
			log.Info("handlers: ", err)
		}
		res := models.Response[models.Role]{Status: code, Message: role.(models.Role)}
		h.conn.Publish(reply, res)
		log.Info("handlers: role found")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) roleReadByName() {
	_, err := h.conn.Subscribe("roles-get-by-name", func(_, reply string, name string) {
		role, code, err := h.cache.GetRole(name, context.Background())
		if err == nil && role != nil {
			res := models.Response[models.Role]{Status: code, Message: role.(models.Role)}
			h.conn.Publish(reply, res)
			log.Info("handlers: company found by name")
		}
		role, code, err = h.role.RoleFindOneByName(name)
		if err != nil {
			res := models.Response[models.Role]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
			log.Error("handlers: ", err)
		}
		if role.(models.Role).ID == nil {
			res := models.Response[models.Role]{Status: 404, Error: "role not found"}
			h.conn.Publish(reply, res)
			log.Info("handlers: ", err)
		}
		code, err = h.cache.Set(name, role, context.Background())
		if err != nil {
			res := models.Response[models.Role]{Status: 200, Error: err.Error()}
			h.conn.Publish(reply, res)
			log.Info("handlers: ", err)
		}
		res := models.Response[models.Role]{Status: code, Message: role.(models.Role)}
		h.conn.Publish(reply, res)
		log.Info("handlers: role found by name")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) roleCreate() {
	_, err := h.conn.Subscribe("roles-create", func(_, reply string, role models.Role) {
		code, err := h.role.RoleCreateOne(role)
		if err != nil {
			res := models.Response[models.Role]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
			log.Error("handlers: ", err)
		}
		code, err = h.cache.Set(role.ID.String(), role, context.Background())
		if err != nil {
			res := models.Response[models.Role]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
			log.Error("handlers: ", err)
		}
		res := models.Response[models.Role]{Status: code, Message: role}
		h.conn.Publish(reply, res)
		log.Info("handlers: role created")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) roleUpdate() {
	_, err := h.conn.Subscribe("roles-put", func(_, reply string, role models.Role) {
		code, err := h.role.RoleUpdateOne(role)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
			log.Error("handlers: ", err)
		}
		code, err = h.cache.Set(role.ID.String(), role, context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
			log.Error("handlers: ", err)
		}
		res := models.Response[string]{Status: code, Message: "role updated"}
		h.conn.Publish(reply, res)
		log.Info("handlers: role updated")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) roleDelete() {
	_, err := h.conn.Subscribe("roles-delete", func(_, reply string, id uuid.UUID) {
		code, err := h.role.RoleDeleteOne(id)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
			log.Error("handlers: ", err)
		}
		code, err = h.cache.Delete(id.String(), context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
			log.Error("handlers: ", err)
		}
		res := models.Response[string]{Status: code, Message: "role deleted"}
		h.conn.Publish(reply, res)
		log.Info("handlers: role deleted")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}
