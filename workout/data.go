package workout

import (
	"time"
)

var WorkoutData []WorkoutModel = []WorkoutModel{
	{
		ID:        1,
		Name:      "Cardio",
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        2,
		Name:      "Swimming",
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        3,
		Name:      "Weight Training",
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        4,
		Name:      "Zumba",
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        5,
		Name:      "Yoga",
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
