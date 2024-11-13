package posts

import "time"

type (
	UpsertUserPostActivityRequest struct {
		IsLiked bool `json:"is_liked"`
	}

	UserPostActivityModel struct {
		ID        int64     `json:"id"`
		UserID    int64     `json:"user_id"`
		PostID    int64     `json:"post_id"`
		IsLiked   bool      `json:"is_liked"`
		CreatedAt time.Time `json:"created_at"`
		CreatedBy int64     `json:"created_by"`
		UpdatedAt time.Time `json:"updated_at"`
		UpdatedBy int64     `json:"updated_by"`
	}
)
