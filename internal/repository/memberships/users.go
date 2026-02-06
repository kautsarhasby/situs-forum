package memberships

import (
	"context"
	"database/sql"

	"github.com/kautsarhasby/situs-forum/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string, id int64) (*memberships.UserModel, error) {
	query := `SELECT id, email, password, username, created_at, created_by, updated_at, updated_by
	FROM users WHERE email = ? OR username = ? OR id = ?`
	row := r.db.QueryRowContext(ctx, query, email, username, id)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Password, &response.Username, &response.CreatedAt, &response.CreatedBy, &response.UpdatedAt, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := `INSERT INTO users (email, password, username, created_at, created_by, updated_at, updated_by)
	VALUES (?,?,?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.Username, model.CreatedAt, model.CreatedBy, model.UpdatedAt, model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
