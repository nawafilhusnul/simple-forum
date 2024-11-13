package posts

import "time"

type (
	CreatePostRequest struct {
		Title    string   `json:"title" binding:"required"`
		Content  string   `json:"content" binding:"required"`
		Hashtags []string `json:"hashtags" binding:"required"`
	}

	PostModel struct {
		ID        int64     `db:"id"`
		UserID    int64     `db:"user_id"`
		Title     string    `db:"title"`
		Content   string    `db:"content"`
		Hashtags  string    `db:"hashtags"`
		CreatedAt time.Time `db:"created_at"`
		CreatedBy int64     `db:"created_by"`
		UpdatedAt time.Time `db:"updated_at"`
		UpdatedBy int64     `db:"updated_by"`
	}

	GetAllPostsResponse struct {
		Data       []PostResponse `json:"data"`
		Pagination Pagination     `json:"pagination"`
	}

	PostResponse struct {
		ID       int64    `json:"id"`
		UserID   int64    `json:"user_id"`
		UserName string   `json:"user_name"`
		Title    string   `json:"title"`
		Content  string   `json:"content"`
		Hashtags []string `json:"hashtags"`
		IsLiked  bool     `json:"is_liked"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}

	GetPostByIDResponse struct {
		PostDetail PostResponse      `json:"post_detail"`
		LikeCount  int64             `json:"like_count"`
		Comments   []CommentResponse `json:"comments"`
	}

	CommentResponse struct {
		UserID   int64  `json:"user_id"`
		UserName string `json:"user_name"`
		Content  string `json:"content"`
	}
)
