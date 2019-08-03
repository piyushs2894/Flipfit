package pakage

import (
	"time"
)

var PakageData []PackageModel = []PackageModel{
	{
		ID:             1,
		CentreId:       1,
		ScheduleId:     1,
		WorkoutId:      1,
		Name:           "Aerocity- Cardio",
		AvailableSeats: 30,
		Status:         1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
	{
		ID:             2,
		CentreId:       1,
		ScheduleId:     1,
		WorkoutId:      2,
		Name:           "Aerocity-Swimming",
		AvailableSeats: 30,
		Status:         1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
	{
		ID:             3,
		CentreId:       1,
		ScheduleId:     1,
		WorkoutId:      3,
		Name:           "Aerocity-Weight Training",
		AvailableSeats: 30,
		Status:         1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
	{
		ID:             4,
		CentreId:       1,
		ScheduleId:     2,
		WorkoutId:      4,
		Name:           "Aerocity- Zumba",
		AvailableSeats: 30,
		Status:         1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
	{
		ID:             2,
		CentreId:       1,
		ScheduleId:     2,
		WorkoutId:      1,
		Name:           "Aerocity-Cardio",
		AvailableSeats: 30,
		Status:         1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
	{
		ID:             3,
		CentreId:       1,
		ScheduleId:     2,
		WorkoutId:      5,
		Name:           "Aerocity-Yoga",
		AvailableSeats: 30,
		Status:         1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
}
