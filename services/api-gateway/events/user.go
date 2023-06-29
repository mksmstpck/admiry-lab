package events

import (
	"errors"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (u User) UserGetByEmail(email string) (models.User, int32, error) {
	var res models.Response[models.User]
	err := u.conn.Request("users-email-get", email, &res, time.Second)
	if err != nil {
		log.Error("events.UserGetByEmail: ", err)
		return models.User{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.UserGetByEmail: ", res.Error)
		return models.User{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.UserGetByEmail: user found")
	return res.Message, res.Status, nil
}

func (u User) UserGetByUsername(username string) (models.User, int32, error) {
	var res models.Response[models.User]
	err := u.conn.Request("users-username-get", username, &res, time.Second)
	if err != nil {
		log.Error("events.UserGetByUsername: ", err)
		return models.User{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.UserGetByUsername: ", res.Error)
		return models.User{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.UserGetByUsername: user found")
	return res.Message, res.Status, nil
}

func (u User) UserGetById(id uuid.UUID) (models.User, int32, error) {
	var res models.Response[models.User]
	err := u.conn.Request("users-id-get", id, &res, time.Second)
	if err != nil {
		log.Error("events.UserGetById: ", err)
		return models.User{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.UserGetById: ", res.Error)
		return models.User{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.UserGetById: user found")
	return res.Message, res.Status, nil
}

func (u User) UserPost(user *models.User) (models.User, int32, error) {
	var res models.Response[models.User]
	err := u.conn.Request("users-post", user, &res, time.Second)
	if err != nil {
		log.Error("events.UserPost: ", err)
		return models.User{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.UserPost: ", res.Error)
		return models.User{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.UserPost: user created")
	return res.Message, res.Status, nil
}

func (u User) UserPut(user *models.User) (int32, error) {
	var res models.Response[string]
	err := u.conn.Request("users-put", user, &res, time.Second)
	if err != nil {
		log.Error("events.UserPut: ", err)
		return 500, err
	}
	if res.Error != "" {
		log.Error("events.UserPut: ", res.Error)
		return res.Status, errors.New(res.Error)
	}
	log.Info("events.UserPut: user updated")
	return res.Status, nil
}

func (u User) UserDelete(id uuid.UUID) (int32, error) {
	var res models.Response[string]
	err := u.conn.Request("users-delete", id, &res, time.Second)
	if err != nil {
		log.Error("events.UserDelete: ", err)
		return 500, err
	}
	if res.Error != "" {
		log.Error("events.UserDelete: ", res.Error)
		return res.Status, errors.New(res.Error)
	}
	log.Info("events.UserDelete: user deleted")
	return res.Status, nil
}
