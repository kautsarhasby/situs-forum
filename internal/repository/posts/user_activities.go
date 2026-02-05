package posts

import (
	"context"
	"database/sql"

	"github.com/kautsarhasby/situs-forum/internal/model/posts"
)


func (r *repository) GetUserActivities(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error){
	query := `SELECT id, post_id, user_id, is_liked, created_at, created_by, updated_at, updated_by
	FROM user_activities WHERE post_id = ? AND user_id= ?`

	var response posts.UserActivityModel
	row := r.db.QueryRowContext(ctx,query, model.PostID, model.UserID)

	err:= row.Scan(&response.PostID,&response.UserID ,&response.IsLiked ,&response.CreatedAt, &response.CreatedBy, &response.UpdatedAt, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil,err
	}

	return &response,nil
}


func (r *repository) CreateUserActivites(ctx context.Context, model posts.UserActivityModel) error {
	query := `INSERT INTO user_activities (user_id,post_id, is_liked, created_at, created_by, updated_at, updated_by)
	VALUES (?,?,?,?,?,?,?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostID,model.IsLiked, model.CreatedAt, model.CreatedBy, model.UpdatedAt,model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateUserActivites(ctx context.Context, model posts.UserActivityModel) error {
	query := `UPDATE user_activites SET is_liked = ?, updated_at = ?, updated_by = ? WHERE user_id = ? AND post_id = ?`

	_, err := r.db.ExecContext(ctx, query,model.IsLiked, model.UpdatedAt,model.UpdatedBy, model.UserID, model.PostID)
	if err != nil {
		return err
	}
	return nil
}