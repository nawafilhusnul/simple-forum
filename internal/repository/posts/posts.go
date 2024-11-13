package posts

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	postsModel "github.com/nawafilhusnul/forum/internal/model/posts"
)

func (r *repository) CreatePost(ctx context.Context, post postsModel.PostModel) error {
	query := `INSERT INTO posts (title, user_id, content, hashtags, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, NOW(), NOW(), ?, ?)`

	_, err := r.db.ExecContext(ctx, query, post.Title, post.UserID, post.Content, post.Hashtags, post.CreatedBy, post.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllPosts(ctx context.Context, limit, offset int) (postsModel.GetAllPostsResponse, error) {
	query := `SELECT p.id, p.title, p.user_id, u.name, p.content, p.hashtags, p.created_at, p.created_by, p.updated_at, p.updated_by FROM posts p INNER JOIN users u ON p.user_id = u.id ORDER BY p.created_at DESC LIMIT ? OFFSET ?`

	var res postsModel.GetAllPostsResponse
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return postsModel.GetAllPostsResponse{}, err
	}
	defer rows.Close()

	data := make([]postsModel.PostResponse, 0)
	for rows.Next() {
		var (
			post     postsModel.PostModel
			userName string
		)

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.UserID,
			&userName,
			&post.Content,
			&post.Hashtags,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.UpdatedAt,
			&post.UpdatedBy,
		); err != nil {
			return postsModel.GetAllPostsResponse{}, err
		}

		data = append(data, postsModel.PostResponse{
			ID:       post.ID,
			UserID:   post.UserID,
			UserName: userName,
			Title:    post.Title,
			Content:  fmt.Sprintf("%s...", post.Content[:75]),
			Hashtags: strings.Split(post.Hashtags, ","),
		})
	}

	res.Data = data
	res.Pagination.Limit = limit
	res.Pagination.Offset = offset

	return res, nil
}

func (r *repository) GetPostByID(ctx context.Context, userID int64, postID int64) (*postsModel.PostResponse, error) {
	query := `SELECT 
    p.id, 
    p.title, 
    p.user_id, 
    u.name, 
    p.content, 
    p.hashtags, 
    p.created_at, 
    p.created_by, 
    p.updated_at, 
    p.updated_by,
    upa.is_liked
  FROM posts p 
  INNER JOIN users u ON p.user_id = u.id 
  INNER JOIN user_post_activities upa ON p.id = upa.post_id AND upa.user_id = ?
  WHERE p.id = ?`

	var (
		model    postsModel.PostModel
		userName string
		isLiked  bool
	)
	if err := r.db.QueryRowContext(ctx, query, userID, postID).Scan(
		&model.ID,
		&model.Title,
		&model.UserID,
		&userName,
		&model.Content,
		&model.Hashtags,
		&model.CreatedAt,
		&model.CreatedBy,
		&model.UpdatedAt,
		&model.UpdatedBy,
		&isLiked,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &postsModel.PostResponse{
		ID:       model.ID,
		UserID:   model.UserID,
		UserName: userName,
		Title:    model.Title,
		Content:  model.Content,
		Hashtags: strings.Split(model.Hashtags, ","),
		IsLiked:  isLiked,
	}, nil
}
