package timezone

import (
	"errors"
	"time"
)

type Time struct {
	clock string
	date  string
	loc   *time.Location
}

const (
	timeDoublePointsFormat = "15:04"
	timeSpaceFormat        = "15 04"

	dayPointFormat         = "02.01"
	daySpaceFormat         = "02 01"
	dayPointWithYearFormat = "02.01.2006"
	daySpaceWithYearFormat = "02 01 2006"

	dayTimeFormat = "02.01.2006 15:04"
	day           = 24 * time.Hour
)

var (
	ErrCantParseTime = errors.New("can't parse time")
	ErrCantParseDate = errors.New("can't parse date")
)

var (
	timeFormats = []string{timeDoublePointsFormat, timeSpaceFormat}
	dayFormats  = []string{dayPointFormat, daySpaceFormat, dayPointWithYearFormat, daySpaceWithYearFormat}
)

func NewEmpty(loc *time.Location) *Time {
	return &Time{
		loc:   loc,
		date:  "01.01.0001",
		clock: "00:00",
	}
}

func New(date time.Time, clock time.Duration, loc *time.Location) *Time {
	t := date.Add(clock).In(loc)
	return &Time{
		clock: t.Format(timeDoublePointsFormat),
		date:  t.Format(dayPointWithYearFormat),
		loc:   loc,
	}
}

func NewFromTime(t time.Time, loc *time.Location) *Time {
	t = t.In(loc)
	return &Time{
		clock: t.Format(timeDoublePointsFormat),
		date:  t.Format(dayPointWithYearFormat),
		loc:   loc,
	}
}

func (t Time) Clock() time.Duration {
	locT := t.Time()

	return locT.Sub(locT.Truncate(day))
}

func (t *Time) ClockString() string {
	return t.clock
}

func (t Time) Date() time.Time {
	return t.Time().Truncate(day).In(time.UTC)
}

func (t *Time) Time() time.Time {
	locT, err := time.ParseInLocation(dayTimeFormat, t.date+" "+t.clock, t.loc)
	if err != nil {
		panic(err)
	}

	return locT.In(time.UTC)
}

func (t *Time) DateString() string {
	return t.date
}

func (t *Time) String() string {
	return t.date + " " + t.clock
}

func (t *Time) SetClock(clock string) error {
	for _, format := range timeFormats {
		tz, err := time.ParseInLocation(format, clock, t.loc)
		if err == nil {
			t.clock = tz.Format(timeDoublePointsFormat)

			return nil
		}
	}

	return ErrCantParseTime
}

func (t *Time) SetDate(str string) error {
	for _, format := range dayFormats {
		tz, err := time.ParseInLocation(format, str, t.loc)
		if err == nil {
			if year, _, _ := tz.Date(); year == 0 {
				tz = tz.AddDate(time.Now().Year(), 0, 0)
			}

			t.date = tz.Format(dayPointWithYearFormat)

			return nil
		}
	}

	return ErrCantParseDate
}

func TodayDateString(loc *time.Location) string {
	return time.Now().In(loc).Format(dayPointFormat)
}

func TomorrowDateString(loc *time.Location) string {
	return time.Now().AddDate(0, 0, 1).In(loc).Format(dayPointFormat)
}
