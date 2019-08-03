package pakage

import "fmt"

func GetPackageByScheduleID(scheduleID int64) ([]PackageModel, error) {
	var result []PackageModel

	//As all package Models are already loaded into global variable PakageData. So
	//directly querying on that, otherwise mysql of es query has to called to fetch packages
	for _, pkg := range PackageModelMap {
		if pkg.ScheduleId == scheduleID {
			result = append(result, pkg)
		}
	}
	if len(result) == 0 {
		return result, fmt.Errorf("No Workout available in this schedule")
	}

	return result, nil
}
