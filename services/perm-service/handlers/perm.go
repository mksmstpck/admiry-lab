package handlers

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/mkskstpck/admiry-lab/pkg/utils"
	"github.com/pborman/uuid"
)

func (h *Handler) permReadById() {
	_, err := h.conn.Subscribe("perm-get-by-id", func(_, reply string, id uuid.UUID) {
		perm, code, err := h.cache.GetPermission(id.String(), context.Background())
		if err == nil && perm != nil {
			res := models.Response[models.Permission]{Status: code, Message: perm.(models.Permission)}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: perm found by id")
			return
		}
		perm, code, err = h.perm.PermFindOneById(id)
		if err != nil {
			res := models.Response[models.Permission]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
			return
		}
		if perm.(models.Permission).ID == nil {
			res := models.Response[models.Permission]{Status: 404, Error: "permission not found"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: perm not found by id")
			return
		}
		code, err = h.cache.Set(id.String(), perm, context.Background())
		if err != nil {
			res := models.Response[models.Permission]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
			return
		}
		res := models.Response[models.Permission]{Status: code, Message: perm.(models.Permission)}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: perm found by id")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) permReadByName() {
	_, err := h.conn.Subscribe("perm-get-by-name", func(_, reply string, name string) {
		perm, code, err := h.cache.GetPermission(name, context.Background())
		if err == nil && perm != nil {
			res := models.Response[models.Permission]{Status: code, Message: perm.(models.Permission)}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: perm found by name")
			return
		}
		perm, code, err = h.perm.PermFindOneByName(name)
		if err != nil {
			res := models.Response[models.Permission]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
			return
		}
		if perm.(models.Permission).ID == nil {
			res := models.Response[models.Permission]{Status: code, Error: "permission not found"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: perm not found by name")
			return
		}
		code, err = h.cache.Set(name, perm, context.Background())
		if err != nil {
			res := models.Response[models.Permission]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
			return
		}
		res := models.Response[models.Permission]{Status: code, Message: perm.(models.Permission)}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: perm found by name")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) permCreate() {
	_, err := h.conn.Subscribe("perm-create", func(_, reply string, perm models.Permission) {
		permExist, code, err := h.perm.PermFindOneByName(perm.Name)
		if err != nil {
			if err.Error() != "permission not found" {
				res := models.Response[models.Permission]{Status: code, Error: err.Error()}
				utils.NatsPublishError(h.conn.Publish(reply, res))
				log.Info("handlers: ", err)
				return
			}
		}
		if permExist.ID != nil {
			res := models.Response[models.Permission]{Status: 409, Error: "permission already exists"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: perm already exists")
			return
		}
		code, err = h.perm.PermCreateOne(perm)
		if err != nil {
			res := models.Response[models.Permission]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
			return
		}
		perm, code, err = h.perm.PermFindOneByName(perm.Name)
		if err != nil {
			res := models.Response[models.Permission]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
			return
		}
		code, err = h.cache.Set(perm.ID.String(), perm, context.Background())
		if err != nil {
			res := models.Response[models.Permission]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
			return
		}
		res := models.Response[models.Permission]{Status: code, Message: perm}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: perm created")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) permUpdate() {
	_, err := h.conn.Subscribe("perm-update", func(_, reply string, perm models.Permission) {
		code, err := h.perm.PermUpdateOne(perm)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
			return
		}
		code, err = h.cache.Set(perm.ID.String(), perm, context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
			return
		}
		res := models.Response[string]{Status: code, Message: "permission updated"}
		utils.NatsPublishError(h.conn.Publish(reply, res))
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) permDelete() {
	_, err := h.conn.Subscribe("perm-delete", func(_, reply string, id uuid.UUID) {
		code, err := h.perm.PermDeleteOne(id)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: ", err)
			return
		}
		code, err = h.cache.Delete(id.String(), context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		res := models.Response[string]{Status: code, Message: "permission deleted"}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: perm deleted")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}
