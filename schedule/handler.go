package schedule

import "fmt"

func GetAllSchedules(centreID int64) ([]ScheduleModel, error) {
	var result []ScheduleModel

	//As all package Models are already loaded into global variable PakageData. So
	//directly querying on that, otherwise mysql of es query has to called to fetch packages
	for _, sch := range ScheduleMap {
		if sch.CentreID == centreID {
			result = append(result, sch)
		}
	}
	if len(result) == 0 {
		return result, fmt.Errorf("No Schedules available for this centre")
	}

	return result, nil
}
