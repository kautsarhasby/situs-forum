package posts

import "time"

type (
	CreatePostRequest struct{
		PostTitle	string `json:"postTitle"`
		PostContent	string `json:"postContent"`
		PostHashtags	[]string `json:"postHashtags"`

	}
)

 type (
	PostModel struct {
		ID int64 `db:"id"`
		UserID int64 `db:"user_id"`
		PostTitle string `db:"postTitle"`
		PostContent string `db:"postContent"`
		PostHashtags string `db:"postHashtags"`
		CreatedAt time.Time `db:"createdAt"`
		UpdatedAt time.Time `db:"updatedAt"`
		CreatedBy string `db:"createdBy"`
		UpdatedBy string `db:"updatedBy"`
	}
 )
    