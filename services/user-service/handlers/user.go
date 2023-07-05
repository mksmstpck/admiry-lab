package handlers

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/mkskstpck/to-rename/pkg/utils"
	"github.com/pborman/uuid"
)

func (h *Handler) userReadById() {
	_, err := h.conn.Subscribe("users-id-get", func(_, reply string, id uuid.UUID) {
		user, code, err := h.cache.GetUser(id.String(), context.Background())
		if err == nil && user != nil {
			res := models.Response[models.User]{Status: code, Message: user.(models.User)}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: user found by id")
		}
		user, code, err = h.user.UserFindOneById(id)
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		if user.(models.User).ID == nil {
			res := models.Response[models.User]{Status: 404, Error: "user not found"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: user not found by id")
		}
		code, err = h.cache.Set(id.String(), user, context.Background())
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		res := models.Response[models.User]{Status: 200, Message: user.(models.User)}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: user found by id")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) userReadByUsername() {
	_, err := h.conn.Subscribe("users-username-get", func(_, reply string, username string) {
		user, code, err := h.cache.GetUser(username, context.Background())
		if err == nil && user != nil {
			res := models.Response[models.User]{Status: code, Message: user.(models.User)}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: user found by username")
		}
		user, code, err = h.user.UserFindOneByUsername(username)
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		if user.(models.User).ID == nil {
			res := models.Response[models.User]{Status: 404, Error: "user not found"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: user not found by username")
		}
		code, err = h.cache.Set(username, user, context.Background())
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		res := models.Response[models.User]{Status: 200, Message: user.(models.User)}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: user found by username")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) userReadByEmail() {
	_, err := h.conn.Subscribe("users-email-get", func(_, reply string, email string) {
		user, code, err := h.cache.GetUser(email, context.Background())
		if err == nil && user != nil {
			res := models.Response[models.User]{Status: code, Message: user.(models.User)}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: user found by email")
		}
		user, code, err = h.user.UserFindOneByEmail(email)
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		if user.(models.User).ID == nil {
			res := models.Response[models.User]{Status: 404, Error: "user not found"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: user not found by email")
		}
		code, err = h.cache.Set(email, user, context.Background())
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		res := models.Response[models.User]{Status: 200, Message: user.(models.User)}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: user found by email")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) userCreate() {
	_, err := h.conn.Subscribe("users-post", func(_, reply string, user models.User) {
		userExistUsername, code, err := h.user.UserFindOneByUsername(user.Username)
		if err != nil {
			if err.Error() != "user not found" {
				res := models.Response[models.User]{Status: code, Error: err.Error()}
				utils.NatsPublishError(h.conn.Publish(reply, res))
				log.Error("handlers: ", err)
			}
		}
		if userExistUsername.ID != nil {
			res := models.Response[models.User]{Status: 409, Error: "user with this username already exists"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: user with this username already exists")
			return
		}
		userExistEmail, code, err := h.user.UserFindOneByEmail(user.Email)
		if err != nil {
			if err.Error() != "user not found" {
				res := models.Response[models.User]{Status: code, Error: err.Error()}
				utils.NatsPublishError(h.conn.Publish(reply, res))
				log.Error("handlers: ", err)
			}
		}
		if userExistEmail.ID != nil {
			res := models.Response[models.User]{Status: 409, Error: "user with this email already exists"}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Info("handlers: user with this email already exists")
			return
		}
		code, err = h.user.UserCreateOne(user)
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		user, code, err = h.user.UserFindOneByEmail(user.Email)
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		code, err = h.cache.Set(user.ID.String(), user, context.Background())
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		res := models.Response[models.User]{Status: 201, Message: user}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: user created")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) userUpdate() {
	_, err := h.conn.Subscribe("users-put", func(_, reply string, user models.User) {
		code, err := h.user.UserUpdateOne(user)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		code, err = h.cache.Set(user.ID.String(), user, context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
			log.Error("handlers: ", err)
		}
		res := models.Response[string]{Status: code, Message: "updated"}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: user updated")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}

func (h *Handler) userDelete() {
	_, err := h.conn.Subscribe("users-delete", func(_, reply string, id uuid.UUID) {
		code, err := h.user.UserDeleteOne(id)
		if err != nil {
			log.Error("handlers: ", err)
			res := models.Response[string]{Status: code, Error: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
		}
		code, err = h.cache.Delete(id.String(), context.Background())
		if err != nil {
			log.Error("handlers: ", err)
			res := models.Response[string]{Status: code, Message: err.Error()}
			utils.NatsPublishError(h.conn.Publish(reply, res))
		}
		res := models.Response[string]{Status: code, Message: "deleted"}
		utils.NatsPublishError(h.conn.Publish(reply, res))
		log.Info("handlers: user deleted")
	})
	if err != nil {
		log.Error("handlers: ", err)
	}
}
