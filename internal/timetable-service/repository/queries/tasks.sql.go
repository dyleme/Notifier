// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: tasks.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addTask = `-- name: AddTask :one
INSERT INTO tasks (
    user_id,
    message,
    required_time
) VALUES (
             $1,
             $2,
             $3
         )
RETURNING id, created_at, message, user_id, required_time, periodic, done, archived
`

type AddTaskParams struct {
	UserID       int32
	Message      string
	RequiredTime pgtype.Interval
}

func (q *Queries) AddTask(ctx context.Context, arg AddTaskParams) (Task, error) {
	row := q.db.QueryRow(ctx, addTask, arg.UserID, arg.Message, arg.RequiredTime)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Message,
		&i.UserID,
		&i.RequiredTime,
		&i.Periodic,
		&i.Done,
		&i.Archived,
	)
	return i, err
}

const deleteTask = `-- name: DeleteTask :one
DELETE FROM tasks
WHERE id = $1
  AND user_id = $2
RETURNING count(*) as deleted_amount
`

type DeleteTaskParams struct {
	ID     int32
	UserID int32
}

func (q *Queries) DeleteTask(ctx context.Context, arg DeleteTaskParams) (int64, error) {
	row := q.db.QueryRow(ctx, deleteTask, arg.ID, arg.UserID)
	var deleted_amount int64
	err := row.Scan(&deleted_amount)
	return deleted_amount, err
}

const getTask = `-- name: GetTask :one
SELECT id, created_at, message, user_id, required_time, periodic, done, archived
  FROM tasks
 WHERE id = $1
   AND user_id = $2
`

type GetTaskParams struct {
	ID     int32
	UserID int32
}

func (q *Queries) GetTask(ctx context.Context, arg GetTaskParams) (Task, error) {
	row := q.db.QueryRow(ctx, getTask, arg.ID, arg.UserID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Message,
		&i.UserID,
		&i.RequiredTime,
		&i.Periodic,
		&i.Done,
		&i.Archived,
	)
	return i, err
}

const listUserTasks = `-- name: ListUserTasks :many
SELECT id, created_at, message, user_id, required_time, periodic, done, archived
  FROM tasks
 WHERE tasks.user_id = $1
`

func (q *Queries) ListUserTasks(ctx context.Context, userID int32) ([]Task, error) {
	rows, err := q.db.Query(ctx, listUserTasks, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Message,
			&i.UserID,
			&i.RequiredTime,
			&i.Periodic,
			&i.Done,
			&i.Archived,
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

const updateTask = `-- name: UpdateTask :exec
UPDATE tasks
   SET
       required_time = $1,
       message = $2,
       periodic = $3,
       done = $4,
       archived = $5
 WHERE id = $6
   AND user_id = $7
`

type UpdateTaskParams struct {
	RequiredTime pgtype.Interval
	Message      string
	Periodic     bool
	Done         bool
	Archived     bool
	ID           int32
	UserID       int32
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.db.Exec(ctx, updateTask,
		arg.RequiredTime,
		arg.Message,
		arg.Periodic,
		arg.Done,
		arg.Archived,
		arg.ID,
		arg.UserID,
	)
	return err
}
