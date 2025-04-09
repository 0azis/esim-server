package redis

import (
	"context"
	"esim/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis interface {
	SetCode(email string, code int) error
	GetCode(email string) (int, error)
}

type redisCli struct {
	cli *redis.Client
}

func New(cfg config.Config) redisCli {
	r := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr(),
		Password: cfg.Redis.Password(),
	})

	return redisCli{r}
}

func (r redisCli) SetCode(email string, code int) error {
	s := r.cli.Set(context.Background(), email, code, time.Second*35)
	return s.Err()
}

func (r redisCli) GetCode(email string) (int, error) {
	s := r.cli.Get(context.Background(), email)
	return s.Int()
}
