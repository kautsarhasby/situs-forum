package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/kautsarhasby/situs-forum/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, request memberships.SignUpRequest) error {
	user,err := s.membershiprepo.GetUser(ctx, request.Email, request.Username)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("username or email alerady exists")
	}

	pass, err:= bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now:= time.Now()
	model := memberships.UserModel {
		Email : request.Email,
		Username : request.Username,
		Password : string(pass),
		CreatedAt: now,
		UpdatedAt : now,
		CreatedBy : request.Email,
		UpdatedBy: request.Email,
	}
	return s.membershiprepo.CreateUser(ctx, model)
}
