package events

import (
	"errors"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/pborman/uuid"
)

func (r Role) RoleGetById(id uuid.UUID) (models.Role, int32, error) {
	var res models.Response[models.Role]
	err := r.conn.Request("roles-get-by-id", id, &res, time.Second)
	if err != nil {
		log.Error("events.RoleGetById: ", err)
		return models.Role{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.RoleGetById: ", res.Error)
		return models.Role{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.RoleGetById: role found")
	return res.Message, res.Status, nil
}

func (r Role) RoleGetByName(name string) (models.Role, int32, error) {
	var res models.Response[models.Role]
	err := r.conn.Request("roles-get-by-name", name, &res, time.Second)
	if err != nil {
		log.Error("events.RoleGetByName: ", err)
		return models.Role{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.RoleGetByName: ", res.Error)
		return models.Role{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.RoleGetByName: role found")
	return res.Message, res.Status, nil
}

func (r Role) RolePost(role *models.Role) (models.Role, int32, error) {
	var res models.Response[models.Role]
	err := r.conn.Request("roles-post", role, &res, time.Second)
	if err != nil {
		log.Error("events.RolePost: ", err)
		return models.Role{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.RolePost: ", res.Error)
		return models.Role{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.RolePost: role found")
	return res.Message, res.Status, nil
}

func (r Role) RolePut(role *models.Role) (int32, error) {
	var res models.Response[string]
	err := r.conn.Request("roles-put", role, &res, time.Second)
	if err != nil {
		log.Error("events.RolePut: ", err)
		return 500, err
	}
	if res.Error != "" {
		log.Error("events.RolePut: ", res.Error)
		return res.Status, errors.New(res.Error)
	}
	log.Info("events.RolePut: role found")
	return res.Status, nil
}

func (r Role) RoleDelete(id uuid.UUID) (int32, error) {
	var res models.Response[string]
	err := r.conn.Request("roles-delete", id, &res, time.Second)
	if err != nil {
		log.Error("events.RoleDelete: ", err)
		return 500, err
	}
	if res.Error != "" {
		log.Error("events.RoleDelete: ", res.Error)
		return res.Status, errors.New(res.Error)
	}
	log.Info("events.RoleDelete: role found")
	return res.Status, nil
}
