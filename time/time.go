package booking

import (
	"fmt"
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
	//"7/25/2019 13:45:00"
	d, _ := time.Parse("1/2/2006 15:04:05", date)
	t:= d.Format(time.UnixDate)
	// fmt.Print("date after parsing",d.Month())
	fmt.Println("Same, in UTC:", t)
	// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
