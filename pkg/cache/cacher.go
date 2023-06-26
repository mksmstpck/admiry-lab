package cache

import (
	"context"
	"encoding/json"

	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/redis/go-redis/v9"
)

func (c *Cacher) GetCompany(key string, ctx context.Context) (interface{}, int32, error) {
	var res models.Company
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, 404, nil
	}
	if err != nil {
		return nil, 500, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		return nil, 500, err
	}
	return res, 200, nil
}

func (c *Cacher) GetUser(key string, ctx context.Context) (interface{}, int32, error) {
	var res models.User
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, 404, nil
	}
	if err != nil {
		return nil, 500, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		return nil, 500, err
	}
	return res, 200, nil
}

func (c *Cacher) Set(key string, value interface{}, ctx context.Context) (int32, error) {
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

func (c *Cacher) Delete(key string, ctx context.Context) (int32, error) {
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		return 500, err
	}
	return 204, nil
}
