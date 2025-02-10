package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrOriginalLinktNotFound = errors.New("original link not found")
)

type OriginalLink struct {
	ID int64 `db:"id"`

	OriginalLink  string        `db:"link"`
	ShorterLinkID sql.NullInt64 `db:"shorter_link_id"`
	CreatedAt     time.Time     `db:"created_at"`
}

type IOriginalLinkRepository interface {
	CreateOriginalLink(ctx context.Context, originalLink string, shorterLinkID int64) error

	GetOriginalLinkByShorterLinkID(ctx context.Context, shorterLinkID int64) (string, error)
}

type IOriginalLinkUsecase interface {
	CreateOriginalLink(ctx context.Context, originalLink string, shorterLinkID int64) error

	GetOriginalLinkByShorterLinkID(ctx context.Context, shorterLinkID int64) (string, error)
}
