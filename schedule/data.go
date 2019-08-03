package schedule

import (
	"github.com/flipfit/lib"
	"time"
)

//Multiple Schedules for a week. After week end. All schedule will be refreshed by cron to +7days
var ScheduleData []ScheduleModel = []ScheduleModel{
	{
		ID:        1,
		CentreID:  1,
		StartTime: lib.TimeToUnixSeconds(time.Now()),
		EndTime:   lib.TimeToUnixSeconds(time.Now().Add(time.Minute * time.Duration(90))),
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        2,
		CentreID:  2,
		StartTime: lib.TimeToUnixSeconds(time.Now()),
		EndTime:   lib.TimeToUnixSeconds(time.Now().Add(time.Minute * time.Duration(90))),
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
