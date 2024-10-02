package redisSource

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type Redis struct {
	rdb *redis.Client
}

func NewRedis(rdb *redis.Client) *Redis {
	return &Redis{rdb: rdb}
}

func (r Redis) PerformQuery(ammo int, ctx context.Context) error {
	ammoStr := strconv.Itoa(ammo)

	_, err := r.rdb.HGetAll(ctx, ammoStr).Result()
	if err != nil {
		return err
	}
	return nil
}
