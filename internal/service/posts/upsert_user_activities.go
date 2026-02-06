package posts

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/kautsarhasby/situs-forum/internal/model/posts"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivity, err := s.postRepo.GetUserActivities(ctx, model)
	if err != nil {
		log.Fatal(err)
	}

	if userActivity == nil {
		//create
		if !request.IsLiked {
			return errors.New("You didnt have liked before")
		}

		if err := s.postRepo.CreateUserActivites(ctx, model); err != nil {
			log.Fatal(err)
			return err
		}

	} else {
		//update
		if err := s.postRepo.UpdateUserActivites(ctx, model); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil

}
