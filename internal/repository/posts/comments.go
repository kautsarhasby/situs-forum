package posts

import (
	"context"
	"database/sql"

	"github.com/kautsarhasby/situs-forum/internal/model/posts"
)

func (r *repository) GetComments(ctx context.Context) (*posts.CommentModel, error){
	query := `SELECT id, post_id, user_id, comment_content, created_at, created_by, updated_at, updated_by
	FROM comments`

	row:=r.db.QueryRowContext(ctx,query)

	var response posts.CommentModel
	err := row.Scan(&response.ID,&response.PostID, &response.UserID,&response.CommentContent, &response.CreatedAt, &response.CreatedBy, &response.UpdatedAt, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil,nil
		}
		return nil,err
	}
	return &response, nil
}


func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments (user_id, post_id, comment_content ,created_at, created_by, updated_at, updated_by)
	VALUES (?,?,?,?,?,?,?) `

	_ ,err := r.db.ExecContext( ctx,query, model.UserID,model.PostID,model.CommentContent, model.CreatedAt, model.CreatedBy, model.UpdatedAt,model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}