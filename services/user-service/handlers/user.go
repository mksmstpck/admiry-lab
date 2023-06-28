package handlers

import (
	"context"

	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (h *Handler) userReadById() {
	h.conn.Subscribe("users-id-get", func(_, reply string, id uuid.UUID) {
		user, code, err := h.cache.GetUser(id.String(), context.Background())
		if err == nil && user != nil {
			res := models.Response[models.User]{Status: code, Message: user.(models.User)}
			h.conn.Publish(reply, res)
		}
		user, code, err = h.user.UserFindOneById(id)
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		if user.(models.User).ID == nil {
			res := models.Response[models.User]{Status: 404, Error: "user not found"}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Set(id.String(), user, context.Background())
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[models.User]{Status: 200, Message: user.(models.User)}
		h.conn.Publish(reply, res)
	})
}

func (h *Handler) userReadByUsername() {
	h.conn.Subscribe("users-username-get", func(_, reply string, username string) {
		user, code, err := h.cache.GetUser(username, context.Background())
		if err == nil && user != nil {
			res := models.Response[models.User]{Status: code, Message: user.(models.User)}
			h.conn.Publish(reply, res)
		}
		user, code, err = h.user.UserFindOneByUsername(username)
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		if user.(models.User).ID == nil {
			res := models.Response[models.User]{Status: 404, Error: "user not found"}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Set(username, user, context.Background())
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[models.User]{Status: 200, Message: user.(models.User)}
		h.conn.Publish(reply, res)
	})
}

func (h *Handler) userReadByEmail() {
	h.conn.Subscribe("users-email-get", func(_, reply string, email string) {
		user, code, err := h.cache.GetUser(email, context.Background())
		if err == nil && user != nil {
			res := models.Response[models.User]{Status: code, Message: user.(models.User)}
			h.conn.Publish(reply, res)
		}
		user, code, err = h.user.UserFindOneByEmail(email)
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		if user.(models.User).ID == nil {
			res := models.Response[models.User]{Status: 404, Error: "user not found"}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Set(email, user, context.Background())
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[models.User]{Status: 200, Message: user.(models.User)}
		h.conn.Publish(reply, res)
	})
}

func (h *Handler) userCreate() {
	h.conn.Subscribe("users-post", func(_, reply string, user models.User) {
		userExistUsername, code, err := h.user.UserFindOneByUsername(user.Username)
		if err != nil {
			if err.Error() != "user not found" {
				res := models.Response[models.User]{Status: code, Error: err.Error()}
				h.conn.Publish(reply, res)
			}
		}
		if userExistUsername.ID != nil {
			res := models.Response[models.User]{Status: 409, Error: "user with this username already exists"}
			h.conn.Publish(reply, res)
			return
		}
		userExistEmail, code, err := h.user.UserFindOneByEmail(user.Email)
		if err != nil {
			if err.Error() != "user not found" {
				res := models.Response[models.User]{Status: code, Error: err.Error()}
				h.conn.Publish(reply, res)
			}
		}
		if userExistEmail.ID != nil {
			res := models.Response[models.User]{Status: 409, Error: "user with this email already exists"}
			h.conn.Publish(reply, res)
			return
		}
		code, err = h.user.UserCreateOne(user)
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		user, code, err = h.user.UserFindOneByEmail(user.Email)
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Set(user.ID.String(), user, context.Background())
		if err != nil {
			res := models.Response[models.User]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[models.User]{Status: 201, Message: user}
		h.conn.Publish(reply, res)
	})
}

func (h *Handler) userUpdate() {
	h.conn.Subscribe("users-put", func(_, reply string, user models.User) {
		code, err := h.user.UserUpdateOne(user)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Set(user.ID.String(), user, context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[string]{Status: code, Message: "updated"}
		h.conn.Publish(reply, res)
	})
}

func (h *Handler) userDelete() {
	h.conn.Subscribe("users-delete", func(_, reply string, id uuid.UUID) {
		code, err := h.user.UserDeleteOne(id)
		if err != nil {
			res := models.Response[string]{Status: code, Error: err.Error()}
			h.conn.Publish(reply, res)
		}
		code, err = h.cache.Delete(id.String(), context.Background())
		if err != nil {
			res := models.Response[string]{Status: code, Message: err.Error()}
			h.conn.Publish(reply, res)
		}
		res := models.Response[string]{Status: code, Message: "deleted"}
		h.conn.Publish(reply, res)
	})
}
