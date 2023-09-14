// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package queries

import (
	domains "github.com/Dyleme/Notifier/internal/timetable-service/domains"
	"github.com/jackc/pgx/v5/pgtype"
)

type DefaultUserNotificationParam struct {
	UserID    int32
	CreatedAt pgtype.Timestamp
	Params    domains.NotificationParams
}

type Event struct {
	ID           int32
	CreatedAt    pgtype.Timestamp
	Text         string
	Description  pgtype.Text
	UserID       int32
	Start        pgtype.Timestamp
	Done         bool
	Notification domains.Notification
}

type Task struct {
	ID           int32
	CreatedAt    pgtype.Timestamp
	Message      string
	UserID       int32
	RequiredTime pgtype.Interval
	Periodic     bool
	Done         bool
	Archived     bool
}

type User struct {
	ID           int32
	Email        pgtype.Text
	PasswordHash pgtype.Text
	TgID         pgtype.Int4
}
