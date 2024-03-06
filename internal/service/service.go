package service

import (
	"context"
	"gin/repository"
	"gin/types"
	"math/rand"
	"strings"
)

type Service struct {
	repository repository.Repository
	cache      repository.Cache
}

func New(repo repository.Repository, cache repository.Cache) *Service {
	return &Service{
		repository: repo,
		cache:      cache,
	}
}

func (s *Service) CreateUser(ctx context.Context, user *types.User) error {
	return s.cache.CreateUser(ctx, user)
}

func (s *Service) ReadUser(ctx context.Context) (*types.User, bool) {
	return s.cache.ReadUser(ctx)
}

func (s *Service) SaveUser(ctx context.Context, user *types.User) error {
	return s.repository.SaveUser(ctx, user)
}

func (s *Service) GenerateRandomLogin() types.User {

	var user types.User

	var usernameBuilder strings.Builder
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 8; i++ {
		randomIndex := rand.Intn(len(charset))
		usernameBuilder.WriteByte(charset[randomIndex])
	}
	user.Username = usernameBuilder.String()

	var passwordBuilder strings.Builder
	for i := 0; i < 10; i++ {
		randomIndex := rand.Intn(len(charset))
		passwordBuilder.WriteByte(charset[randomIndex])
	}
	user.Password = passwordBuilder.String()

	return user

}
