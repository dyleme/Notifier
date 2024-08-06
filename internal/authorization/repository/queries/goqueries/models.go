// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package goqueries

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type TaskType string

const (
	TaskTypePeriodicTask TaskType = "periodic_task"
	TaskTypeBasicTask    TaskType = "basic_task"
)

func (e *TaskType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TaskType(s)
	case string:
		*e = TaskType(s)
	default:
		return fmt.Errorf("unsupported scan type for TaskType: %T", src)
	}
	return nil
}

type NullTaskType struct {
	TaskType TaskType
	Valid    bool // Valid is true if TaskType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTaskType) Scan(value interface{}) error {
	if value == nil {
		ns.TaskType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TaskType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTaskType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TaskType), nil
}

type BasicTask struct {
	ID                 int32              `db:"id"`
	CreatedAt          pgtype.Timestamp   `db:"created_at"`
	Text               string             `db:"text"`
	Description        pgtype.Text        `db:"description"`
	UserID             int32              `db:"user_id"`
	Start              pgtype.Timestamptz `db:"start"`
	NotificationParams []byte             `db:"notification_params"`
	Notify             bool               `db:"notify"`
}

type BindingAttempt struct {
	ID             int32            `db:"id"`
	TgID           int32            `db:"tg_id"`
	LoginTimestamp pgtype.Timestamp `db:"login_timestamp"`
	Code           string           `db:"code"`
	Done           bool             `db:"done"`
	PasswordHash   string           `db:"password_hash"`
}

type DefaultUserNotificationParam struct {
	UserID    int32            `db:"user_id"`
	CreatedAt pgtype.Timestamp `db:"created_at"`
	Params    []byte           `db:"params"`
}

type Event struct {
	ID                 int32              `db:"id"`
	CreatedAt          pgtype.Timestamp   `db:"created_at"`
	UserID             int32              `db:"user_id"`
	Text               string             `db:"text"`
	Description        pgtype.Text        `db:"description"`
	TaskID             int32              `db:"task_id"`
	TaskType           TaskType           `db:"task_type"`
	NextSendTime       pgtype.Timestamptz `db:"next_send_time"`
	Done               bool               `db:"done"`
	NotificationParams []byte             `db:"notification_params"`
	FirstSendTime      pgtype.Timestamptz `db:"first_send_time"`
	LastSendedTime     pgtype.Timestamptz `db:"last_sended_time"`
	Notify             bool               `db:"notify"`
}

type KeyValue struct {
	Key   string `db:"key"`
	Value []byte `db:"value"`
}

type PeriodicTask struct {
	ID                 int32              `db:"id"`
	CreatedAt          pgtype.Timestamp   `db:"created_at"`
	Text               string             `db:"text"`
	Description        pgtype.Text        `db:"description"`
	UserID             int32              `db:"user_id"`
	Start              pgtype.Timestamptz `db:"start"`
	SmallestPeriod     int32              `db:"smallest_period"`
	BiggestPeriod      int32              `db:"biggest_period"`
	NotificationParams []byte             `db:"notification_params"`
	Notify             bool               `db:"notify"`
}

type SmthToTag struct {
	SmthID int32 `db:"smth_id"`
	TagID  int32 `db:"tag_id"`
	UserID int32 `db:"user_id"`
}

type Tag struct {
	ID        int32              `db:"id"`
	CreatedAt pgtype.Timestamptz `db:"created_at"`
	Name      string             `db:"name"`
	UserID    int32              `db:"user_id"`
}

type TgImage struct {
	ID       int32  `db:"id"`
	Filename string `db:"filename"`
	TgFileID string `db:"tg_file_id"`
}

type User struct {
	ID             int32       `db:"id"`
	PasswordHash   pgtype.Text `db:"password_hash"`
	TgID           int32       `db:"tg_id"`
	TimezoneOffset int32       `db:"timezone_offset"`
	TimezoneDst    bool        `db:"timezone_dst"`
	TgNickname     string      `db:"tg_nickname"`
}
