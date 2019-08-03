package schedule

import (
	"time"
)

type ScheduleModel struct {
	ID       int64
	CentreID int64

	StartTime int64
	EndTime   int64

	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
