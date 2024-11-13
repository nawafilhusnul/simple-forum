package memberships

import (
	"context"
	"database/sql"
	"errors"

	membershipsModel "github.com/nawafilhusnul/forum/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, userName string, userID int64) (*membershipsModel.UserModel, error) {
	query := `SELECT id, name, user_name, email, password, created_at, updated_at, created_by, updated_by FROM users WHERE email = ? OR user_name = ? OR id = ?`

	row := r.db.QueryRowContext(ctx, query, email, userName, userID)

	var user membershipsModel.UserModel
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *repository) CreateUser(ctx context.Context, user *membershipsModel.UserModel) error {
	query := `INSERT INTO users (name, user_name, email, password, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, NOW(), NOW(), ?, ?)`

	_, err := r.db.ExecContext(ctx, query, user.Name, user.UserName, user.Email, user.Password, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
