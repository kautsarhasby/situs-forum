package posts

import (
	"context"

	"github.com/kautsarhasby/situs-forum/internal/model/posts"
)

func (r *repository) GetComments(ctx context.Context, )


func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments (user_id, post_id,comment_content ,created_at, created_by, updated_at, updated_by)
	VALUES (?,?,?,?,?,?,?) `

	_ ,err := r.db.ExecContext( ctx,query, model.UserID,model.PostID,model.CommentContent, model.CreatedAt, model.CreatedBy, model.UpdatedAt,model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}