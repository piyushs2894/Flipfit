package booking

//Global Variables to be initialized on app start

//
var BookingMap map[int64]BookingModel

//Initialize past bookings by loading data from database. Currently taking it as empty array
func Init() {
	BookingMap = make(map[int64]BookingModel)

	for _, booking := range Bookings {
		BookingMap[booking.ID] = booking
	}
}
