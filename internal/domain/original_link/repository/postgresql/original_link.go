package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/Maksim646/ozon_link_shorter/internal/database/postgresql"
	"github.com/Maksim646/ozon_link_shorter/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
)

type OriginalLinkRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IOriginalLinkRepository {
	return &OriginalLinkRepository{sqalxConn: sqalxConn}
}

func (r *OriginalLinkRepository) CreateOriginalLink(ctx context.Context, originalLink string, shorterLinkID int64) error {
	query, params, err := postgresql.Builder.Insert("original_links").
		Columns(
			"link",
			"shorter_link_id",
		).
		Values(
			originalLink,
			shorterLinkID,
		).
		ToSql()
	if err != nil {
		return err
	}
	slog.Debug(postgresql.BuildQuery(query, params))

	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *OriginalLinkRepository) GetOriginalLinkByShorterLinkID(ctx context.Context, shorterLinkID int64) (string, error) {
	var originalLink string
	query, params, err := postgresql.Builder.Select(
		"original_links.original_link",
	).
		From("original_links").
		Where(sq.Eq{"original_links.shorter_link_id": shorterLinkID}).ToSql()
	if err != nil {
		return originalLink, err
	}

	slog.Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &originalLink, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return originalLink, model.ErrOriginalLinktNotFound
		}
	}

	return originalLink, err
}
