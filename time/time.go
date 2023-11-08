package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {

	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t

}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {

	d, _ := time.Parse("January 2, 2006 15:04:05", date)
	now := time.Now()
	isAfter := now.After(d)
	return isAfter

}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {

	d, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	if d.Hour() >= 12 && d.Hour() < 18 {

		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	d, _ := time.Parse("1/2/2006 15:04:05", date)

	return d.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")

	 
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
