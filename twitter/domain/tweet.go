package domain

import (
	"context"

	"github.com/syedwshah/twitter"
)

type TweetService struct {
	TweetRepo twitter.TweetRepo
}

func NewTweetService(tr twitter.TweetRepo) *TweetService {
	return &TweetService{
		TweetRepo: tr,
	}
}

func (ts *TweetService) All(ctx context.Context) ([]twitter.Tweet, error) {
	panic("not implemented") // TODO: Implement
}

func (ts *TweetService) Create(ctx context.Context, input twitter.CreateTweetInput) (twitter.Tweet, error) {
	currentUserID, err := twitter.GetUserIDFromContext(ctx)
	if err != nil {
		return twitter.Tweet{}, twitter.ErrUnauthenticated
	}

	input.Sanitize()

	if err := input.Validate(); err != nil {
		return twitter.Tweet{}, err
	}

	tweet, err := ts.TweetRepo.Create(ctx, twitter.Tweet{
		Body:   input.Body,
		UserID: currentUserID,
	})
	if err != nil {
		return twitter.Tweet{}, err
	}

	return tweet, nil
}

func (ts *TweetService) GetByID(ctx context.Context, id string) (twitter.Tweet, error) {
	panic("not implemented") // TODO: Implement
}