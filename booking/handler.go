package booking

import (
	"fmt"
	"github.com/flipfit/pakage"
	"time"
)

type BookingModel struct {
	ID         int64
	UserID     int64
	CentreID   int64
	ScheduleID int64
	PackageID  int64
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func BookWorkoutActivity(userID int64, pkg pakage.PackageModel) (BookingModel, error) {
	var booking BookingModel

	//Do user authentication if valid or not

	//Check availability
	if pkg.AvailableSeats < 1 {
		return booking, fmt.Errorf("No seats available")
	}

	booking = BookingModel{
		ID:         getLastBookingID() + 1,
		UserID:     userID,
		CentreID:   pkg.CentreId,
		ScheduleID: pkg.ScheduleId,
		PackageID:  pkg.ID,
		Status:     1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	//Register new booking in global variable
	registerBooking(booking)

	return booking, nil

}

//Get last booking id. Just used for local datastore. This not needed in case of any database
func getLastBookingID() int64 {
	var lastID int64

	for _, booking := range Bookings {
		lastID = booking.ID
	}

	return lastID
}

func registerBooking(booking BookingModel) {
	//Register new booking in global variable
	Bookings = append(Bookings, booking)
	pkg := pakage.PackageModelMap[booking.PackageID]
	pkg.AvailableSeats = pkg.AvailableSeats - 1
	pakage.PackageModelMap[booking.PackageID] = pkg

}

// To cancel just marking status = 0. Not allowing deletion
func CancelBooking(booking BookingModel) {
	//Register new booking in global variable
	for id, _ := range Bookings {
		if Bookings[id] == booking {
			Bookings[id].Status = 0
		}
	}
}

func ViewAllBookings(userID int64) []BookingModel {
	var bookings []BookingModel
	for _, booking := range Bookings {
		if userID == booking.UserID {
			bookings = append(bookings, booking)
		}
	}
	return bookings
}
