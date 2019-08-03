package workout

//Global Variables to be initialized on app start
var WorkoutTypeMap map[int64]WorkoutModel

func Init() {
	WorkoutTypeMap = make(map[int64]WorkoutModel)

	LoadWorkoutType()
}

func LoadWorkoutType() {
	for _, work := range WorkoutData {
		WorkoutTypeMap[work.ID] = work
	}
}
