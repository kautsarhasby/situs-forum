package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/kautsarhasby/situs-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)


func (s *service) CreatePost(ctx context.Context,userID int64, request posts.CreatePostRequest ) error {
	postHastags := strings.Join(request.PostHashtags, ",")

	now:= time.Now()
	model := posts.PostModel{
		UserID: userID,
		PostTitle: request.PostTitle,
		PostContent: request.PostContent,
		PostHashtags: postHastags,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	err:= s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Msg("Cannot Create Post")
		return err
	}

	return nil
}