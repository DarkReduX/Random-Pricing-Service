// Package repository - contains implementation of redis repository with pricing
package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"main/src/internal/data"
	"os"
	"time"
)

// PriceRepository - repository data type
type PriceRepository struct {
	Client *redis.Client
}

const (
	envRedisURI      = "REDIS_URI"
	envRedisUserName = "REDIS_USERNAME"
	envRedisPassword = "REDIS_PASS"
)

// ConnectRepository - connect to cloud repository using environment variables
// REDIS_URI - url of redis database
// REDIS_USERNAME - database username
// REDIS_PASS - database user password
func (r *PriceRepository) ConnectRepository() error {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv(envRedisURI),
		Username: os.Getenv(envRedisUserName),
		Password: os.Getenv(envRedisPassword),
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	if _, err := r.Client.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}

func (r PriceRepository) SendNewPrice(price data.SymbolPrice) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	r.Client.Set(ctx, price.Uuid, price, 0)
}
