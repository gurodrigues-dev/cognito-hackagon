package repository

import (
	"context"
	"gin/config"
	"gin/types"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	conn *redis.Client
}

func NewRedisClient() (*Redis, error) {

	conf := config.Get()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Cache.Address,
		Password: conf.Cache.Password,
		DB:       0,
	})

	repo := &Redis{
		conn: client,
	}

	return repo, nil

}

func (r *Redis) CreateUser(ctx context.Context, user *types.User) error {

	err := r.conn.Set("username", user.Username, 6*time.Hour).Err()

	if err != nil {
		return err
	}

	err = r.conn.Set("password", user.Password, 6*time.Hour).Err()

	if err != nil {
		return err
	}

	return nil

}

func (r *Redis) ReadUser(ctx context.Context) (*types.User, bool) {

	username, err := r.conn.Get("username").Result()

	if err != nil {
		log.Printf("error while searching username")
		return nil, false
	}

	password, err := r.conn.Get("password").Result()

	if err != nil {
		log.Printf("error while searching password")
		return nil, false
	}

	user := &types.User{
		Username: username,
		Password: password,
	}

	return user, true

}
