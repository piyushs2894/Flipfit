package schedule

//Global Variables to be initialized on app start

//Refers to different timeslot in particular centres
var ScheduleMap map[int64]ScheduleModel

func Init() {
	ScheduleMap = make(map[int64]ScheduleModel)

	LoadScheduleDataFile()
}

func LoadScheduleDataFile() {
	for _, sch := range ScheduleData {
		ScheduleMap[sch.ID] = sch
	}
}
