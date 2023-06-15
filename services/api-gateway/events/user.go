package events

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mkskstpck/to-rename/pkg/models"
)

func (u User) UserEmailGet(email string) (models.User, int32, error) {
	var res models.Response[models.User]
	err := u.conn.Request("users-email-get", email, &res, time.Second)
	if err != nil {
		return models.User{}, 500, err
	}
	if res.Error != "" {
		return models.User{}, res.Status, errors.New(res.Error)
	}
	return res.Message, res.Status, nil
}

func (u User) UserUsernameGet(username string) (models.User, int32, error) {
	var res models.Response[models.User]
	err := u.conn.Request("users-username-get", username, &res, time.Second)
	if err != nil {
		return models.User{}, 500, err
	}
	if res.Error != "" {
		return models.User{}, res.Status, errors.New(res.Error)
	}
	return res.Message, res.Status, nil
}

func (u User) UserIdGet(id uuid.UUID) (models.User, int32, error) {
	var res models.Response[models.User]
	err := u.conn.Request("users-id-get", id, &res, time.Second)
	if err != nil {
		return models.User{}, 500, err
	}
	if res.Error != "" {
		return models.User{}, res.Status, errors.New(res.Error)
	}
	return res.Message, res.Status, nil
}

func (u User) UserPost(user *models.User) (models.User, int32, error) {
	var res models.Response[models.User]
	err := u.conn.Request("users-post", user, &res, time.Second)
	if err != nil {
		return models.User{}, 500, err
	}
	if res.Error != "" {
		return models.User{}, res.Status, errors.New(res.Error)
	}
	return res.Message, res.Status, nil
}

func (u User) UserPut(user *models.User) (int32, error) {
	var res models.Response[string]
	err := u.conn.Request("users-put", user, &res, time.Second)
	if err != nil {
		return 500, err
	}
	if res.Error != "" {
		return res.Status, errors.New(res.Error)
	}
	return res.Status, nil
}

func (u User) UserDelete(id uuid.UUID) (int32, error) {
	var res models.Response[string]
	err := u.conn.Request("users-delete", id, &res, time.Second)
	if err != nil {
		return 500, err
	}
	if res.Error != "" {
		return res.Status, errors.New(res.Error)
	}
	return res.Status, nil
}
