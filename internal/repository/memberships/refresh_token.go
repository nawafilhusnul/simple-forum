package memberships

import (
	"context"
	"database/sql"
	"errors"

	membershipsModel "github.com/nawafilhusnul/forum/internal/model/memberships"
)

func (r *repository) InsertRefreshToken(ctx context.Context, refreshTokenModel membershipsModel.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, issued_at, expired_at, created_at, created_by, updated_at, updated_by) VALUES (?, ?, NOW(), ?, NOW(), ?, NOW(), ?)`
	_, err := r.db.ExecContext(ctx, query,
		refreshTokenModel.UserID,
		refreshTokenModel.RefreshToken,
		refreshTokenModel.ExpiredAt,
		refreshTokenModel.CreatedBy,
		refreshTokenModel.UpdatedBy,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetRefreshTokenByUserID(ctx context.Context, userID int64) (*membershipsModel.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, issued_at, expired_at, created_at, created_by, updated_at, updated_by FROM refresh_tokens WHERE user_id = ? AND expired_at > NOW() ORDER BY created_at DESC LIMIT 1`
	var refreshToken membershipsModel.RefreshTokenModel
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&refreshToken.ID,
		&refreshToken.UserID,
		&refreshToken.RefreshToken,
		&refreshToken.IssuedAt,
		&refreshToken.ExpiredAt,
		&refreshToken.CreatedAt,
		&refreshToken.CreatedBy,
		&refreshToken.UpdatedAt,
		&refreshToken.UpdatedBy,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &refreshToken, nil
}
