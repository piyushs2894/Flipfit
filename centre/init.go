package centre

//Global Variables to be initialized on app start

//Refers to different timeslot in particular centres
var CentresMap map[int64]CentreModel

func Init() {
	CentresMap = make(map[int64]CentreModel)

	for _, centre := range CentreData {
		CentresMap[centre.ID] = centre
	}
}
