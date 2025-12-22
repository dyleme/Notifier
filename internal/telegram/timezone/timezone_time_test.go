package timezone

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimezoneTime_SetClockSetDate(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name          string
		settedClock   string
		settedDate    string
		tz            *time.Location
		expectedStr   string
		expectedClock time.Duration
		expectedDate  time.Time
	}{
		{
			name:          "utc",
			settedClock:   "10:00",
			settedDate:    "15.11.2025",
			tz:            time.UTC,
			expectedStr:   "15.11.2025 10:00",
			expectedClock: 10 * time.Hour,
			expectedDate:  time.Date(2025, 11, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:          "time in utc+3",
			settedClock:   "10:00",
			settedDate:    "15.11.2025",
			tz:            time.FixedZone("UTC+3", 3*60*60),
			expectedStr:   "15.11.2025 10:00",
			expectedClock: 7 * time.Hour,
			expectedDate:  time.Date(2025, 11, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:          "occuring in previous day",
			settedClock:   "01:00",
			settedDate:    "15.11.2025",
			tz:            time.FixedZone("UTC+3", 3*60*60),
			expectedStr:   "15.11.2025 01:00",
			expectedClock: 22 * time.Hour,
			expectedDate:  time.Date(2025, 11, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			name:          "occuring in next day",
			settedClock:   "22:00",
			settedDate:    "15.11.2025",
			tz:            time.FixedZone("UTC-3", -3*60*60),
			expectedStr:   "15.11.2025 22:00",
			expectedClock: 1 * time.Hour,
			expectedDate:  time.Date(2025, 11, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			name:          "00:00",
			settedClock:   "00:00",
			settedDate:    "15.11.2025",
			tz:            time.UTC,
			expectedStr:   "15.11.2025 00:00",
			expectedClock: 0 * time.Hour,
			expectedDate:  time.Date(2025, 11, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:          "00:00 in utc+3",
			settedClock:   "00:00",
			settedDate:    "15.11.2025",
			tz:            time.FixedZone("UTC+3", 3*60*60),
			expectedStr:   "15.11.2025 00:00",
			expectedClock: 21 * time.Hour,
			expectedDate:  time.Date(2025, 11, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			name:          "00:00 in utc-3",
			settedClock:   "00:00",
			settedDate:    "15.11.2025",
			tz:            time.FixedZone("UTC-3", -3*60*60),
			expectedStr:   "15.11.2025 00:00",
			expectedClock: 3 * time.Hour,
			expectedDate:  time.Date(2025, 11, 15, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ttz := NewEmpty(tc.tz)
			ttz.SetClock(tc.settedClock)
			ttz.SetDate(tc.settedDate)

			assert.Equal(t, tc.expectedStr, ttz.String())
			assert.Equal(t, tc.expectedClock, ttz.Clock())
			assert.Equal(t, tc.expectedDate, ttz.Date())
		})
	}
}

func TestTimezoneTime_NewEmpty(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		ttz := NewEmpty(time.UTC)

		assert.Equal(t, "01.01.0001", ttz.date)
		assert.Equal(t, "00:00", ttz.clock)
	})

	t.Run("fill", func(t *testing.T) {
		t.Parallel()
		ttz := NewEmpty(time.UTC)
		ttz.SetClock("10:00")
		ttz.SetDate("15.11.2025")

		assert.Equal(t, "15.11.2025", ttz.date)
		assert.Equal(t, "10:00", ttz.clock)
	})
}

func TestTime_New(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		clock        time.Duration
		date         time.Time
		tz           *time.Location
		expectedTime time.Time
		expectedStr  string
	}{
		{
			name:         "utc",
			date:         time.Date(2025, 11, 15, 0, 0, 0, 0, time.UTC),
			clock:        12 * time.Hour,
			tz:           time.UTC,
			expectedTime: time.Date(2025, 11, 15, 12, 0, 0, 0, time.UTC),
			expectedStr:  "15.11.2025 12:00",
		},
		{
			name:         "utc+3",
			date:         time.Date(2025, 11, 15, 0, 0, 0, 0, time.UTC),
			clock:        12 * time.Hour,
			tz:           time.FixedZone("UTC+3", 3*60*60),
			expectedTime: time.Date(2025, 11, 15, 12, 0, 0, 0, time.UTC),
			expectedStr:  "15.11.2025 15:00",
		},
		{
			name:         "through days border",
			clock:        23 * time.Hour,
			date:         time.Date(2025, 11, 14, 0, 0, 0, 0, time.UTC),
			tz:           time.FixedZone("UTC+3", 3*60*60),
			expectedTime: time.Date(2025, 11, 14, 23, 0, 0, 0, time.UTC),
			expectedStr:  "15.11.2025 02:00",
		},
		{
			name:         "only clock",
			date:         time.Time{},
			clock:        12 * time.Hour,
			tz:           time.FixedZone("UTC+3", 3*60*60),
			expectedTime: time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC),
			expectedStr:  "01.01.0001 15:00",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ttz := New(tc.date, tc.clock, tc.tz)

			assert.Equal(t, tc.expectedTime, ttz.Time().In(time.UTC))
			assert.Equal(t, tc.expectedStr, ttz.String())
		})
	}
}
