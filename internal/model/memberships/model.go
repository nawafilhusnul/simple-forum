package memberships

import "time"

type (
	SignUpRequest struct {
		Name     string `json:"name" binding:"required"`
		UserName string `json:"user_name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	SignInRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	SignInResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token,omitempty"`
	}

	UserModel struct {
		ID        int64     `db:"id"`
		Name      string    `db:"name"`
		UserName  string    `db:"user_name"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy int64     `db:"created_by"`
		UpdatedBy int64     `db:"updated_by"`
	}

	RefreshTokenModel struct {
		ID           int64     `db:"id"`
		UserID       int64     `db:"user_id"`
		RefreshToken string    `db:"refresh_token"`
		IssuedAt     time.Time `db:"issued_at"`
		ExpiredAt    time.Time `db:"expired_at"`
		CreatedAt    time.Time `db:"created_at"`
		CreatedBy    int64     `db:"created_by"`
		UpdatedAt    time.Time `db:"updated_at"`
		UpdatedBy    int64     `db:"updated_by"`
	}

	ValidateRefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	ValidateRefreshTokenResponse struct {
		AccessToken string `json:"access_token"`
	}
)
