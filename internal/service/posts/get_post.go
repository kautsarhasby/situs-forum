package posts

import (
	"context"

	"github.com/kautsarhasby/situs-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, id int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id to database")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostID(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("error get count like by post id to database")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentsByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("error get comments by post id to database")
		return nil, err
	}

	return &posts.GetPostResponse{

		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}
