package user

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/flipfit/booking"
	"github.com/flipfit/centre"
	"github.com/flipfit/constant"
	"github.com/flipfit/lib"
	"github.com/flipfit/pakage"
	"github.com/flipfit/schedule"
	"io"
	"strconv"
)

//Global Variables to be initialized on app start
var UserModelMap map[string]UserModel
var CurrUser UserModel

func Init() {
	var option int
	var userInput int
	var err error

	fileName := fmt.Sprintf("%s/%s%s", constant.PARENT_DIRECTORY, constant.FILE_PATH, constant.USER_FILE_NAME)

	LoadUserDataFile(fileName)

	fmt.Println("Login or Register. For Login, enter 1 and for SignUp, enter 2. For EXIT, enter 3")
	fmt.Scanln(&option)

	switch option {
	case 1:
		if err = LoginHandler(); err != nil {
			fmt.Println("Failed to Login. Error: ", err)
			return
		}
		goto INIT
	case 2:
		if err = SignupHandler(); err != nil {
			fmt.Println("Failed to Login. Error ", err)
			return
		}
		goto INIT
	case 3:
		fmt.Println("Exiting....... ")
		return
	default:
		fmt.Println("Invalid options entered ", option)
		return
	}

INIT:
	fmt.Println("Do You want to see all relevant centres? Press 1 for yes. 2 to see all bookings. 3 to exit \n")
	fmt.Scanln(&option)
	switch option {
	case 1:
		centres, err := centre.GetAllCentres()
		if err != nil {
			fmt.Println("Failed to GetAllCentres. Error ", err)
			return
		}

		//Fetch all centres
		fmt.Printf("Select Centre: %+v\n. Press centre ID as input\n", centres)
		fmt.Scanln(&userInput)

		////Fetch all schedules
		schedules, err := schedule.GetAllSchedules(int64(userInput))
		if err != nil {
			fmt.Println("Invalid ScheduleID id. Error ", err, userInput)
			goto INIT
		}
		fmt.Printf("Select Schedule: %+v\n. Press schedule ID as input\n", schedules)
		fmt.Scanln(&userInput)

		//Fetch all packages
		pakages, err := pakage.GetPackageByScheduleID(int64(userInput))
		if err != nil {
			fmt.Println("Invalid Schedule ID id. Error ", err)
			goto INIT
		}
		fmt.Printf("Select Package You want to book: %+v\n. Press package ID as input\n", pakages)

		fmt.Scanln(&userInput)
		bookingInfo, err := booking.BookWorkoutActivity(CurrUser.ID, pakage.PackageModelMap[int64(userInput)])
		if err != nil {
			fmt.Println("Invalid Schedule ID id. Error ", err)
			goto INIT
		}

		fmt.Printf("Booking Done: %+v\n", bookingInfo)
		goto INIT
	case 2:
		bookings := booking.ViewAllBookings(CurrUser.ID)
		fmt.Printf("BOOKING INFO :   %+v\n", bookings)
	case 3:
		fmt.Println("Exiting....... ")
		return
	}
}

func LoadUserDataFile(fileName string) {
	//Initialize global UserDataMap
	UserModelMap = make(map[string]UserModel)

	csvFile, err := lib.OpenFile(fileName)
	if err != nil {
		fmt.Printf("[LoadUserDataFile] Error: %+v\n", err)
	}

	defer csvFile.Close()
	//Checking if userName exists or not
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error in reading csv records: ", err)
			return
		}

		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid User ID ", err)
		}

		var row UserModel
		row = UserModel{
			ID:       id,
			UserName: record[1],
			Password: record[2],
			Name:     record[3],
			Email:    record[4],
			Mobile:   record[5],
		}

		UserModelMap[row.UserName] = row
	}

	defer csvFile.Close()

	return
}
