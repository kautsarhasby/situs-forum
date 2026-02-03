package memberships

import (
	"context"

	"github.com/kautsarhasby/situs-forum/internal/configs"
	"github.com/kautsarhasby/situs-forum/internal/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	cfg *configs.Config
	membershiprepo membershipRepository
}

func NewService(membershipRepo membershipRepository, cfg *configs.Config) *service {
	return &service{
		cfg : cfg,
		membershiprepo:  membershipRepo,
	}
}