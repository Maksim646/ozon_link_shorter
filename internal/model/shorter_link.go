package model

import (
	"context"
	// "database/sql"
	"time"
)

type ShorterLink struct {
	ID int64 `db:"id"`

	ShorterLink string    `db:"shorter_link"`
	CreatedAt   time.Time `db:"created_at"`
}

type IShorterLinkRepository interface {
	CreateShorterLink(ctx context.Context, originalLink string) error

	GetShorterLinkByOriginalLink(ctx context.Context, originalLink string) (string, error)
}

type IShorterLinkUsecase interface {
	CreateShorterLink(ctx context.Context, originalLink string) error

	GetShorterLinkByOriginalLink(ctx context.Context, originalLink string) (string, error)
}
