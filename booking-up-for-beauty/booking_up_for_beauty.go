package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "2006-Jan-02 15:04:05 -0700 MST"
	dateFormatted, err := time.Parse(layout, date)
	if err != nil {
		return dateFormatted
	}
	return time.Now()
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {

	return Schedule(date).After(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
