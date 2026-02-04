package posts

import "time"

type (
	CreateCommentRequest struct{
		CommentContent	string `json:"commentContent"`
	}
)

 type (
	CommentModel struct {
		ID int64 `db:"id"`
		UserID int64 `db:"user_id"`
		PostID int64	`db:"post_id"`
		CommentContent string `db:"postContent"`
		CreatedAt time.Time `db:"createdAt"`
		UpdatedAt time.Time `db:"updatedAt"`
		CreatedBy string `db:"createdBy"`
		UpdatedBy string `db:"updatedBy"`
	}
 )
    