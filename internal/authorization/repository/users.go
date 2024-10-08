package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Dyleme/Notifier/internal/authorization/repository/queries/goqueries"
	"github.com/Dyleme/Notifier/internal/authorization/service"
	"github.com/Dyleme/Notifier/internal/domains"
	"github.com/Dyleme/Notifier/pkg/serverrors"
	"github.com/Dyleme/Notifier/pkg/sql/pgxconv"
)

func (r *Repository) Create(ctx context.Context, input service.CreateUserInput) (domains.User, error) {
	tx := r.getter.DefaultTrOrDB(ctx, r.db)
	op := "Repository.Create: %w"
	user, err := r.q.AddUser(ctx, tx, goqueries.AddUserParams{
		TgID:       int32(input.TGID),
		TgNickname: input.TGNickname,
	})
	if err != nil {
		if intersection, isUnique := uniqueError(err); isUnique {
			return domains.User{}, fmt.Errorf(op, serverrors.NewUniqueError(intersection, input.TGNickname))
		}

		return domains.User{}, fmt.Errorf(op, serverrors.NewRepositoryError(err))
	}

	return domains.User{
		ID:             int(user.ID),
		TgNickname:     user.TgNickname,
		PasswordHash:   pgxconv.ByteSlice(user.PasswordHash),
		TGID:           int(user.TgID),
		TimeZoneOffset: 0,
		IsTimeZoneDST:  false,
	}, nil
}

func uniqueError(err error) (string, bool) {
	var pgerr *pgconn.PgError
	if errors.As(err, &pgerr) {
		if pgerr.Code == pgerrcode.UniqueViolation {
			if strings.Contains(pgerr.Detail, "email") {
				return "email", true
			}

			return "", true
		}
	}

	return "", false
}

func (r *Repository) Get(ctx context.Context, email string, tgID int) (domains.User, error) {
	op := "Repository.Get: %w"
	tx := r.getter.DefaultTrOrDB(ctx, r.db)
	out, err := r.q.FindUser(ctx, tx, goqueries.FindUserParams{
		TgNickname: email,
		TgID:       int32(tgID),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domains.User{}, fmt.Errorf(op, serverrors.NewNotFoundError(err, "user"))
		}

		return domains.User{}, fmt.Errorf(op, serverrors.NewRepositoryError(err))
	}

	return domains.User{
		ID:             int(out.ID),
		TgNickname:     out.TgNickname,
		PasswordHash:   pgxconv.ByteSlice(out.PasswordHash),
		TGID:           int(out.TgID),
		TimeZoneOffset: int(out.TimezoneOffset),
		IsTimeZoneDST:  out.TimezoneDst,
	}, nil
}

func (r *Repository) Update(ctx context.Context, user domains.User) error {
	tx := r.getter.DefaultTrOrDB(ctx, r.db)
	err := r.q.UpdateUser(ctx, tx, goqueries.UpdateUserParams{
		TgNickname:     user.TgNickname,
		PasswordHash:   pgxconv.Text(string(user.PasswordHash)),
		TimezoneOffset: int32(user.TimeZoneOffset),
		TimezoneDst:    user.IsTimeZoneDST,
		TgID:           int32(user.TGID),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return serverrors.NewNotFoundError(err, "user")
		}

		return fmt.Errorf("find user: %w", serverrors.NewRepositoryError(err))
	}

	return nil
}

func (r *Repository) AddBindingAttempt(ctx context.Context, input service.BindingAttempt) error {
	tx := r.getter.DefaultTrOrDB(ctx, r.db)
	err := r.q.AddBindingAttempt(ctx, tx, goqueries.AddBindingAttemptParams{
		TgID:         int32(input.TGID),
		Code:         input.Code,
		PasswordHash: input.PasswordHash,
	})
	if err != nil {
		return fmt.Errorf("add binding attempt: %w", serverrors.NewRepositoryError(err))
	}

	return nil
}

func (r *Repository) GetLatestBindingAttempt(ctx context.Context, tgID int) (service.BindingAttempt, error) {
	tx := r.getter.DefaultTrOrDB(ctx, r.db)
	ba, err := r.q.GetLatestBindingAttempt(ctx, tx, int32(tgID))
	if err != nil {
		return service.BindingAttempt{}, fmt.Errorf("get latest binding attempt: %w", serverrors.NewRepositoryError(err))
	}

	return service.BindingAttempt{
		ID:           tgID,
		TGID:         int(ba.TgID),
		Code:         ba.Code,
		PasswordHash: ba.PasswordHash,
		Done:         ba.Done,
	}, nil
}

func (r *Repository) UpdateBindingAttemptStatus(ctx context.Context, baID int, done bool) error {
	tx := r.getter.DefaultTrOrDB(ctx, r.db)
	err := r.q.UpdateBindingAttempt(ctx, tx, goqueries.UpdateBindingAttemptParams{
		ID:   int32(baID),
		Done: done,
	})
	if err != nil {
		return fmt.Errorf("update binding attempt: %w", serverrors.NewRepositoryError(err))
	}

	return nil
}
