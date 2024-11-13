package posts

import (
	"context"

	postModel "github.com/nawafilhusnul/forum/internal/model/posts"
)

func (r *repository) CreatePostComment(ctx context.Context, req *postModel.PostCommentModel) error {
	query := `INSERT INTO post_comments (post_id, comment, created_at, updated_at, created_by, updated_by) VALUES (?, ?, NOW(), NOW(), ?, ?)`
	_, err := r.db.ExecContext(ctx, query, req.PostID, req.Comment, req.CreatedBy, req.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetPostComments(ctx context.Context, postID int64) ([]postModel.CommentResponse, error) {
	query := `SELECT pc.comment, pc.created_by, u.name FROM post_comments pc INNER JOIN users u ON pc.created_by = u.id WHERE pc.post_id = ?`

	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res = make([]postModel.CommentResponse, 0)
	for rows.Next() {
		var (
			comment  postModel.PostCommentModel
			userName string
		)
		if err := rows.Scan(&comment.Comment, &comment.CreatedBy, &userName); err != nil {
			return nil, err
		}
		res = append(res, postModel.CommentResponse{
			UserID:   comment.CreatedBy,
			UserName: userName,
			Content:  comment.Comment,
		})
	}
	return res, nil
}
