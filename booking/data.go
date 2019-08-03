package booking

import (
	"time"
)

var Bookings []BookingModel = []BookingModel{
	{
		ID:         1,
		UserID:     1,
		CentreID:   1,
		ScheduleID: 1,
		PackageID:  1,
		Status:     1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
	{
		ID:         1,
		UserID:     1,
		CentreID:   1,
		ScheduleID: 2,
		PackageID:  4,
		Status:     1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
}
