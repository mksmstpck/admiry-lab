package cache

import (
	"context"
	"encoding/json"

	"github.com/mkskstpck/to-rename/user-service/models"
	"github.com/redis/go-redis/v9"
)

func (c *UserCache) Get(key string, ctx context.Context) (interface{}, int32, error) {
	user := models.User{}
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, 404, nil
	}
	if err != nil {
		return nil, 500, err
	}
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, 500, err
	}
	return user, 200, nil
}

func (c *UserCache) Set(key string, value interface{}, ctx context.Context) (int32, error) {
	val, err := json.Marshal(value)
	if err != nil {
		return 500, err
	}
	err = c.client.Set(ctx, key, val, c.expires).Err()
	if err != nil {
		return 500, err
	}
	return 200, nil
}

func (c *UserCache) Delete(key string, ctx context.Context) (int32, error) {
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		return 500, err
	}
	return 200, nil
}
