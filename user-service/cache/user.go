package cache

import (
	"context"
	"encoding/json"

	"github.com/mkskstpck/to-rename/user-service/models"
)

func (c *UserCache) Get(key string, ctx context.Context) (interface{}, error) {
	user := models.User{}
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *UserCache) Set(key string, value interface{}, ctx context.Context) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = c.client.Set(ctx, key, val, c.expires).Err()
	if err != nil {
		return err
	}
	return nil
}
