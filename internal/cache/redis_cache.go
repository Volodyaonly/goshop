package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"goshop/internal/models"

	"github.com/redis/go-redis/v9"
)

type ProductCache struct {
	rdb *redis.Client
}

func NewProductCache(addr string) *ProductCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &ProductCache{
		rdb: rdb,
	}
}

func (c *ProductCache) Get(
	ctx context.Context,
	id int,
) (*models.Product, error) {

	value, err := c.rdb.Get(
		ctx,
		key(id),
	).Result()

	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var product models.Product

	if err := json.Unmarshal(
		[]byte(value),
		&product,
	); err != nil {
		return nil, err
	}

	return &product, nil
}

func (c *ProductCache) Set(
	ctx context.Context,
	product *models.Product,
) error {

	data, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return c.rdb.Set(
		ctx,
		key(product.ID),
		data,
		5*time.Minute,
	).Err()
}

func key(id int) string {
	return "product:" + strconv.Itoa(id)
}
