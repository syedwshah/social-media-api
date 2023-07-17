package twitter

import (
	"context"
	"fmt"
	"strings"
	"time"
)

var (
	TweetMinLenght = 2
	TweetMaxLenght = 250
)

type CreateTweetInput struct {
	Body string
}

func (in *CreateTweetInput) Sanitize() {
	in.Body = strings.TrimSpace(in.Body)
}

func (in CreateTweetInput) Validate() error {
	if len(in.Body) < TweetMinLenght {
		return fmt.Errorf("%w: body too short, should be at least %d characters", ErrValidation, TweetMinLenght)
	}
	if len(in.Body) > TweetMaxLenght {
		return fmt.Errorf("%w: body too long, should be at most %d characters", ErrValidation, TweetMaxLenght)
	}
	return nil
}

type Tweet struct {
	ID        string
	Body      string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TweetService interface {
	All(ctx context.Context) ([]Tweet, error)
	Create(ctx context.Context, input CreateTweetInput) (Tweet, error)
	GetByID(ctx context.Context, id string) (Tweet, error)
}

type TweetRepo interface {
	All(ctx context.Context) ([]Tweet, error)
	Create(ctx context.Context, tweet Tweet) (Tweet, error)
	GetByID(ctx context.Context, id string) (Tweet, error)
}
