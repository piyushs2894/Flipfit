package pakage

import (
	"time"
)

//It refers to workout in particular centre at particular schedule
type PackageModel struct {
	ID             int64
	CentreId       int64
	ScheduleId     int64
	WorkoutId      int64
	Name           string
	AvailableSeats int64
	Status         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
