package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/dyleme/Notifier/internal/domain"
	"github.com/dyleme/Notifier/internal/domain/apperr"
	"github.com/dyleme/Notifier/internal/repository/queries/goqueries"
	"github.com/dyleme/Notifier/pkg/database/txmanager"
)

type TgImagesRepository struct {
	q      *goqueries.Queries
	getter *txmanager.Getter
	cache  TgImageCache
}

type TgImageCache GenericCache[domain.TgImage]

func NewTGImagesRepository(getter *txmanager.Getter, cache TgImageCache) *TgImagesRepository {
	return &TgImagesRepository{
		q:      goqueries.New(),
		cache:  cache,
		getter: getter,
	}
}

func (t TgImagesRepository) Add(ctx context.Context, filename, tgFileID string) error {
	op := "TgImagesRepository.Add: %w"

	tx := t.getter.GetTx(ctx)
	_, err := t.q.AddTgImage(ctx, tx, goqueries.AddTgImageParams{
		Filename: filename,
		TgFileID: tgFileID,
	})
	if err != nil {
		if intersection, isUnique := uniqueError(err, []string{"filename"}); isUnique {
			return fmt.Errorf(op, apperr.UniqueError{Name: intersection, Value: filename})
		}

		return fmt.Errorf(op, err)
	}

	return nil
}

func uniqueError(err error, columnNames []string) (string, bool) {
	var pgerr *pgconn.PgError
	if errors.As(err, &pgerr) {
		if pgerr.Code == pgerrcode.UniqueViolation {
			for _, columnName := range columnNames {
				if strings.Contains(pgerr.Detail, columnName) {
					return columnName, true
				}
			}

			return "", true
		}
	}

	return "", false
}

func (t TgImagesRepository) Get(ctx context.Context, filename string) (domain.TgImage, error) {
	image, err := t.cache.Wrap(ctx, filename, func() (domain.TgImage, error) {
		tx := t.getter.GetTx(ctx)
		tgImage, err := t.q.GetTgImage(ctx, tx, filename)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return domain.TgImage{}, apperr.ErrNotFound
			}

			return domain.TgImage{}, fmt.Errorf("get tg image: %w", err)
		}
		return domain.TgImage{
			Filename: tgImage.Filename,
			TgFileID: tgImage.TgFileID,
		}, nil
	})
	if err != nil {
		return domain.TgImage{}, err
	}

	return image, nil
}
