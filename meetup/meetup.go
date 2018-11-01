package meetup

import "time"

type WeekSchedule string

var (
	First  WeekSchedule = "First"
	Second WeekSchedule = "Second"
	Third  WeekSchedule = "Third"
	Fourth WeekSchedule = "Fourth"
	Teenth WeekSchedule = "Teenth"
	Last   WeekSchedule = "Last"
)

func Day(week WeekSchedule, day time.Weekday, month time.Month, year int) int {

	return 0
}
