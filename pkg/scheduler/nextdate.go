package scheduler

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat is empty")
	}

	start, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", errors.New("invalid start date")
	}

	repeatParts := strings.Fields(repeat)
	if len(repeatParts) == 0 {
		return "", errors.New("invalid repeat format")
	}

	repeatType := repeatParts[0]

	switch repeatType {
	case "d":
		if len(repeatParts) != 2 {
			return "", errors.New("invalid repeat format for daily repeat")
		}

		interval, err := strconv.Atoi(repeatParts[1])
		if err != nil || interval <= 0 || interval > 400 {
			return "", errors.New("invalid repeat interval for daily repeat")
		}

		for {
			start = start.AddDate(0, 0, interval)
			if start.After(now) {
				return start.Format(DateFormat), nil
			}
		}
	case "y":
		if len(repeatParts) == 1 {
			for {
				start = start.AddDate(1, 0, 0)
				if start.After(now) {
					return start.Format(DateFormat), nil
				}
			}
		}

		return "", errors.New("invalid repeat format for yearly repeat")
	case "w":
		if len(repeatParts) != 2 {
			return "", errors.New("invalid repeat format for weekly repeat")
		}

		return nextByWeek(start, now, repeatParts[1])
	case "m":
		return nextByMonth(start, now, repeatParts)
	default:
		return "", errors.New("unsupported repeat format")
	}
}

func nextByWeek(start, now time.Time, value string) (string, error) {
	var weekdays [8]bool

	for _, raw := range strings.Split(value, ",") {
		day, err := strconv.Atoi(raw)
		if err != nil || day < 1 || day > 7 {
			return "", errors.New("invalid week day")
		}

		weekdays[day] = true
	}

	for date := nextSearchDate(start, now); ; date = date.AddDate(0, 0, 1) {
		weekday := int(date.Weekday())
		if weekday == 0 {
			weekday = 7
		}

		if weekdays[weekday] {
			return date.Format(DateFormat), nil
		}
	}
}

func nextByMonth(start, now time.Time, parts []string) (string, error) {
	if len(parts) != 2 && len(parts) != 3 {
		return "", errors.New("invalid repeat format for monthly repeat")
	}

	var days [32]bool
	lastDay := false
	beforeLastDay := false

	for _, raw := range strings.Split(parts[1], ",") {
		day, err := strconv.Atoi(raw)
		if err != nil {
			return "", errors.New("invalid month day")
		}

		switch {
		case day >= 1 && day <= 31:
			days[day] = true
		case day == -1:
			lastDay = true
		case day == -2:
			beforeLastDay = true
		default:
			return "", errors.New("invalid month day")
		}
	}

	var months [13]bool
	if len(parts) == 3 {
		for _, raw := range strings.Split(parts[2], ",") {
			month, err := strconv.Atoi(raw)
			if err != nil || month < 1 || month > 12 {
				return "", errors.New("invalid month")
			}

			months[month] = true
		}
	} else {
		for month := 1; month <= 12; month++ {
			months[month] = true
		}
	}

	for date := nextSearchDate(start, now); ; date = date.AddDate(0, 0, 1) {
		month := int(date.Month())
		day := date.Day()
		last := lastMonthDay(date)

		if months[month] && (days[day] || lastDay && day == last || beforeLastDay && day == last-1) {
			return date.Format(DateFormat), nil
		}
	}
}

func nextSearchDate(start, now time.Time) time.Time {
	if now.After(start) {
		start = now
	}

	return start.AddDate(0, 0, 1)
}

func lastMonthDay(date time.Time) int {
	firstNextMonth := time.Date(date.Year(), date.Month()+1, 1, 0, 0, 0, 0, date.Location())
	return firstNextMonth.AddDate(0, 0, -1).Day()
}
