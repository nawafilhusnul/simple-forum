package posts

import "time"

type (
	CreatePostCommentRequest struct {
		Comment string `json:"comment" binding:"required"`
	}

	PostCommentModel struct {
		ID        int64     `db:"id"`
		PostID    int64     `db:"post_id"`
		Comment   string    `db:"comment"`
		CreatedAt time.Time `db:"created_at"`
		CreatedBy int64     `db:"created_by"`
		UpdatedAt time.Time `db:"updated_at"`
		UpdatedBy int64     `db:"updated_by"`
	}
)
