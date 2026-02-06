package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/kautsarhasby/situs-forum/internal/model/memberships"
	"github.com/kautsarhasby/situs-forum/pkg/jwt"
	tokenUtil "github.com/kautsarhasby/situs-forum/pkg/token"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, request memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, request.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", "", err
	}
	if user == nil {
		return "", "", errors.New("email not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", "", errors.New("email or passord invalid")
	}

	token, err := jwt.CreateToken(int64(user.ID), user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", "", err
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, int64(user.ID), time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get latest refreshToken")
		return "", "", err
	}
	// jika refershToken nya ada , maka kembalikan refresh yang sudah ada (pakai lagi)
	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	// generate refresh token baru
	refreshToken := tokenUtil.GenerateRefreshToken()
	if refreshToken == "" {
		return token, "", errors.New("failed to get refresh token")
	}
	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:       int64(user.ID),
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(10 * 24 * time.Hour),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(int64(user.ID), 10),
		UpdatedBy:    strconv.FormatInt(int64(user.ID), 10),
	})
	if err != nil {
		log.Error().Err(err).Msg("error inserting  refresh token to database")
		return "", "", err
	}
	return token, refreshToken, nil
}
