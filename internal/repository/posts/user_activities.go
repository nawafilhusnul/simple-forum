package posts

import (
	"context"
	"database/sql"
	"errors"

	postsModel "github.com/nawafilhusnul/forum/internal/model/posts"
)

func (r *repository) CreateUserPostActivity(ctx context.Context, userID int64, postID int64, req postsModel.UserPostActivityModel) error {
	query := `INSERT INTO user_post_activities (user_id, post_id, is_liked, created_by, updated_by) VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, userID, postID, req.IsLiked, req.CreatedBy, req.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUserPostActivity(ctx context.Context, userID int64, postID int64) (*postsModel.UserPostActivityModel, error) {
	query := `SELECT id, user_id, post_id, is_liked, created_at, created_by, updated_at, updated_by FROM user_post_activities WHERE user_id = ? AND post_id = ?`

	var res postsModel.UserPostActivityModel
	if err := r.db.QueryRowContext(ctx, query, userID, postID).Scan(
		&res.ID,
		&res.UserID,
		&res.PostID,
		&res.IsLiked,
		&res.CreatedAt,
		&res.CreatedBy,
		&res.UpdatedAt,
		&res.UpdatedBy,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

func (r *repository) UpdateUserPostActivity(ctx context.Context, userID int64, postID int64, req postsModel.UserPostActivityModel) error {
	query := `UPDATE user_post_activities SET is_liked = ?, updated_by = ? WHERE user_id = ? AND post_id = ?`

	_, err := r.db.ExecContext(ctx, query, req.IsLiked, req.UpdatedBy, userID, postID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CountLikesByPostID(ctx context.Context, postID int64) (int64, error) {
	query := `SELECT COUNT(1) FROM user_post_activities WHERE post_id = ? AND is_liked = TRUE`

	var count int64
	if err := r.db.QueryRowContext(ctx, query, postID).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
