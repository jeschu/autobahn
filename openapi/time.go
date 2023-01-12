package openapi

import (
	"strings"
	"time"
)

const timeFormat = "2006-01-02T15:04:05Z0700"

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	parsed, err := time.Parse(timeFormat, value)
	if err != nil {
		return err
	}
	*t = Time(parsed)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(timeFormat) + `"`), nil
}

// parsing time "2022-11-07T19:00:00.000+0100" as "2006-01-02T15:04:05Z07:00": cannot parse "+0100" as "Z07:00"
