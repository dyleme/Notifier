// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: events.sql

package goqueries

import (
	"context"

	domains "github.com/Dyleme/Notifier/internal/domains"
	"github.com/jackc/pgx/v5/pgtype"
)

const addEvent = `-- name: AddEvent :one
INSERT INTO events (
    user_id,
    text,
    task_id,
    task_type,
    next_send_time, 
    notification_params,
    first_send_time,
    last_sended_time
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $5,
    TIMESTAMP '1970-01-01 00:00:00'
) RETURNING id, created_at, user_id, text, description, task_id, task_type, next_send_time, done, notification_params, first_send_time, last_sended_time
`

type AddEventParams struct {
	UserID             int32                      `db:"user_id"`
	Text               string                     `db:"text"`
	TaskID             int32                      `db:"task_id"`
	TaskType           TaskType                   `db:"task_type"`
	NextSendTime       pgtype.Timestamptz         `db:"next_send_time"`
	NotificationParams domains.NotificationParams `db:"notification_params"`
}

func (q *Queries) AddEvent(ctx context.Context, db DBTX, arg AddEventParams) (Event, error) {
	row := db.QueryRow(ctx, addEvent,
		arg.UserID,
		arg.Text,
		arg.TaskID,
		arg.TaskType,
		arg.NextSendTime,
		arg.NotificationParams,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UserID,
		&i.Text,
		&i.Description,
		&i.TaskID,
		&i.TaskType,
		&i.NextSendTime,
		&i.Done,
		&i.NotificationParams,
		&i.FirstSendTime,
		&i.LastSendedTime,
	)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :many
DELETE FROM events
WHERE id = $1
RETURNING id, created_at, user_id, text, description, task_id, task_type, next_send_time, done, notification_params, first_send_time, last_sended_time
`

func (q *Queries) DeleteEvent(ctx context.Context, db DBTX, id int32) ([]Event, error) {
	rows, err := db.Query(ctx, deleteEvent, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.Text,
			&i.Description,
			&i.TaskID,
			&i.TaskType,
			&i.NextSendTime,
			&i.Done,
			&i.NotificationParams,
			&i.FirstSendTime,
			&i.LastSendedTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEvent = `-- name: GetEvent :one
SELECT id, created_at, user_id, text, description, task_id, task_type, next_send_time, done, notification_params, first_send_time, last_sended_time FROM events
WHERE id = $1
`

func (q *Queries) GetEvent(ctx context.Context, db DBTX, id int32) (Event, error) {
	row := db.QueryRow(ctx, getEvent, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UserID,
		&i.Text,
		&i.Description,
		&i.TaskID,
		&i.TaskType,
		&i.NextSendTime,
		&i.Done,
		&i.NotificationParams,
		&i.FirstSendTime,
		&i.LastSendedTime,
	)
	return i, err
}

const getLatestEvent = `-- name: GetLatestEvent :one
SELECT id, created_at, user_id, text, description, task_id, task_type, next_send_time, done, notification_params, first_send_time, last_sended_time FROM events
WHERE task_id = $1
  AND task_type = $2
ORDER BY next_send_time DESC
LIMIT 1
`

type GetLatestEventParams struct {
	TaskID   int32    `db:"task_id"`
	TaskType TaskType `db:"task_type"`
}

func (q *Queries) GetLatestEvent(ctx context.Context, db DBTX, arg GetLatestEventParams) (Event, error) {
	row := db.QueryRow(ctx, getLatestEvent, arg.TaskID, arg.TaskType)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UserID,
		&i.Text,
		&i.Description,
		&i.TaskID,
		&i.TaskType,
		&i.NextSendTime,
		&i.Done,
		&i.NotificationParams,
		&i.FirstSendTime,
		&i.LastSendedTime,
	)
	return i, err
}

const getNearestEvent = `-- name: GetNearestEvent :one
SELECT id, created_at, user_id, text, description, task_id, task_type, next_send_time, done, notification_params, first_send_time, last_sended_time FROM events
WHERE done = false
ORDER BY next_send_time ASC
LIMIT 1
`

func (q *Queries) GetNearestEvent(ctx context.Context, db DBTX) (Event, error) {
	row := db.QueryRow(ctx, getNearestEvent)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UserID,
		&i.Text,
		&i.Description,
		&i.TaskID,
		&i.TaskType,
		&i.NextSendTime,
		&i.Done,
		&i.NotificationParams,
		&i.FirstSendTime,
		&i.LastSendedTime,
	)
	return i, err
}

const listNotSendedEvents = `-- name: ListNotSendedEvents :many
SELECT id, created_at, user_id, text, description, task_id, task_type, next_send_time, done, notification_params, first_send_time, last_sended_time FROM events
WHERE next_send_time <= $1
  AND done = false
`

func (q *Queries) ListNotSendedEvents(ctx context.Context, db DBTX, till pgtype.Timestamptz) ([]Event, error) {
	rows, err := db.Query(ctx, listNotSendedEvents, till)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.Text,
			&i.Description,
			&i.TaskID,
			&i.TaskType,
			&i.NextSendTime,
			&i.Done,
			&i.NotificationParams,
			&i.FirstSendTime,
			&i.LastSendedTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserEvents = `-- name: ListUserEvents :many
SELECT id, created_at, user_id, text, description, task_id, task_type, next_send_time, done, notification_params, first_send_time, last_sended_time FROM events
WHERE user_id = $1
  AND next_send_time BETWEEN $2 AND $3
ORDER BY next_send_time DESC
LIMIT $5 OFFSET $4
`

type ListUserEventsParams struct {
	UserID   int32              `db:"user_id"`
	FromTime pgtype.Timestamptz `db:"from_time"`
	ToTime   pgtype.Timestamptz `db:"to_time"`
	Off      int32              `db:"off"`
	Lim      int32              `db:"lim"`
}

func (q *Queries) ListUserEvents(ctx context.Context, db DBTX, arg ListUserEventsParams) ([]Event, error) {
	rows, err := db.Query(ctx, listUserEvents,
		arg.UserID,
		arg.FromTime,
		arg.ToTime,
		arg.Off,
		arg.Lim,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.Text,
			&i.Description,
			&i.TaskID,
			&i.TaskType,
			&i.NextSendTime,
			&i.Done,
			&i.NotificationParams,
			&i.FirstSendTime,
			&i.LastSendedTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEvent = `-- name: UpdateEvent :one
UPDATE events
SET text = $1,
    next_send_time = $2,
    first_send_time = $3,
    done = $4
WHERE id = $5
RETURNING id, created_at, user_id, text, description, task_id, task_type, next_send_time, done, notification_params, first_send_time, last_sended_time
`

type UpdateEventParams struct {
	Text          string             `db:"text"`
	NextSendTime  pgtype.Timestamptz `db:"next_send_time"`
	FirstSendTime pgtype.Timestamptz `db:"first_send_time"`
	Done          bool               `db:"done"`
	ID            int32              `db:"id"`
}

func (q *Queries) UpdateEvent(ctx context.Context, db DBTX, arg UpdateEventParams) (Event, error) {
	row := db.QueryRow(ctx, updateEvent,
		arg.Text,
		arg.NextSendTime,
		arg.FirstSendTime,
		arg.Done,
		arg.ID,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UserID,
		&i.Text,
		&i.Description,
		&i.TaskID,
		&i.TaskType,
		&i.NextSendTime,
		&i.Done,
		&i.NotificationParams,
		&i.FirstSendTime,
		&i.LastSendedTime,
	)
	return i, err
}
