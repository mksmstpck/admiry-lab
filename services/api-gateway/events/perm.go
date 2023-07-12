package events

import (
	"errors"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/pborman/uuid"
)

func (p Perm) PermGetById(id uuid.UUID) (models.Permission, int32, error) {
	var res models.Response[models.Permission]
	err := p.conn.Request("perm-get-by-id", id, &res, time.Second)
	if err != nil {
		log.Error("events.PermFindOneById: ", err)
		return models.Permission{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.PermFindOneById: ", res.Error)
		return models.Permission{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.PermFindOneById: permission found")
	return res.Message, res.Status, nil
}

func (p Perm) PermGetByName(name string) (models.Permission, int32, error) {
	var res models.Response[models.Permission]
	err := p.conn.Request("perm-get-by-name", name, &res, time.Second)
	if err != nil {
		log.Error("events.PermFindOneByName: ", err)
		return models.Permission{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.PermFindOneByName: ", res.Error)
		return models.Permission{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.PermFindOneByName: permission found")
	return res.Message, res.Status, nil
}

func (p Perm) PermGetAll() ([]models.Permission, int32, error) {
	var res models.Response[[]models.Permission]
	err := p.conn.Request("perms-get-all", nil, &res, time.Second)
	if err != nil {
		log.Error("events.PermFindAll: ", err)
		return []models.Permission{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.PermFindAll: ", res.Error)
		return []models.Permission{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.PermFindAll: permissions found")
	return res.Message, res.Status, nil
}

func (p Perm) PermPost(perm *models.Permission) (models.Permission, int32, error) {
	var res models.Response[models.Permission]
	err := p.conn.Request("perm-create", perm, &res, time.Second)
	if err != nil {
		log.Error("events.PermCreateOne: ", err)
		return models.Permission{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.PermCreateOne: ", res.Error)
		return models.Permission{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.PermCreateOne: permission created")
	return res.Message, res.Status, nil
}

func (p Perm) PermPut(perm *models.Permission) (int32, error) {
	var res models.Response[string]
	err := p.conn.Request("perm-update", perm, &res, time.Second)
	if err != nil {
		log.Error("events.PermUpdateOne: ", err)
		return 500, err
	}
	if res.Error != "" {
		log.Error("events.PermUpdateOne: ", res.Error)
		return res.Status, errors.New(res.Error)
	}
	log.Info("events.PermUpdateOne: permission updated")
	return res.Status, nil
}

func (p Perm) PermDelete(id uuid.UUID) (int32, error) {
	var res models.Response[string]
	err := p.conn.Request("perm-delete", id, &res, time.Second)
	if err != nil {
		log.Error("events.PermDeleteOne: ", err)
		return 500, err
	}
	if res.Error != "" {
		log.Error("events.PermDeleteOne: ", res.Error)
		return res.Status, errors.New(res.Error)
	}
	log.Info("events.PermDeleteOne: permission deleted")
	return res.Status, nil
}
