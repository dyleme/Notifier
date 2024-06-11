package domains_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Dyleme/Notifier/internal/domains"
)

func TestNewSendingNotification(t *testing.T) {
	t.Parallel()
	t.Run("check mapping", func(t *testing.T) {
		t.Parallel()
		notif := domains.Notification{
			ID:          1,
			UserID:      2,
			Text:        "text",
			Description: "description",
			TaskType:    domains.BasicTaskType,
			TaskID:      3,
			Params: &domains.NotificationParams{
				Period: time.Hour,
				Params: domains.Params{
					Telegram: 4,
				},
			},
			SendTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Sended:   true,
			Done:     false,
		}
		params := domains.NotificationParams{
			Period: time.Minute,
			Params: domains.Params{
				Telegram: 5,
			},
		}
		actual := domains.NewSendingNotification(notif, params)

		expected := domains.SendingNotification{
			NotificationID: 1,
			UserID:         2,
			Message:        "text",
			Description:    "description",
			Params:         params,
			SendTime:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		require.Equal(t, expected, actual)
	})
}
