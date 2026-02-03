package memberships

import "time"

type (
	SignUpRequest struct {
		Email string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
)

type (
	LoginResponse struct {
		AccessToken string `json:"accessToken"`
	}
)

type (
	UserModel struct {
		ID int `db:"id"`
		Email string `db:"email"`
		Username string `db:"username"`
		Password string `db:"password"`
		CreatedAt time.Time `db:"createdAt"`
		UpdatedAt time.Time `db:"updatedAt"`
		CreatedBy string `db:"createdBy"`
		UpdatedBy string `db:"updatedBy"`
	}
)