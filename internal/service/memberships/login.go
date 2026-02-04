package memberships

import (
	"context"
	"errors"
	"fmt"

	"github.com/kautsarhasby/situs-forum/internal/model/memberships"
	"github.com/kautsarhasby/situs-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login (ctx context.Context, request memberships.LoginRequest) (string, error){
	user, err := s.membershiprepo.GetUser(ctx, request.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "",err
		}	
		
		if user == nil {
			return "", errors.New("email not exist")
		}
		
		err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password))
	fmt.Println("user : ",user.Password)
	fmt.Println("req :",request.Password)
	fmt.Println("byte user: ",[]byte(user.Password))
	fmt.Println("byte req: ",[]byte(request.Password))
	if err != nil {
		return "", errors.New("email or passord invalid")
	}

	token, err := jwt.CreateToken(int64(user.ID),user.Username,s.cfg.Service.SecretJWT)
	if err !=nil {
		return "",err
	}

	return token, nil
}