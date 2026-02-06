package posts

import (
	"context"

	"github.com/kautsarhasby/situs-forum/internal/configs"
	"github.com/kautsarhasby/situs-forum/internal/model/posts"
)

type postsRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
	GetUserActivities(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivites(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivites(ctx context.Context, model posts.UserActivityModel) error
	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error)
	CountLikeByPostID(ctx context.Context, postID int64) (int, error)
	GetCommentsByID(ctx context.Context, postID int64) ([]posts.Comment, error)
	GetPostByID(ctx context.Context, id int64) (*posts.Post, error)
}

type service struct {
	cfg      *configs.Config
	postRepo postsRepository
}

func NewService(postRepo postsRepository, cfg *configs.Config) *service {
	return &service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
