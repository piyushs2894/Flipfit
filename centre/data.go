package centre

import (
	"time"
)

var CentreData []CentreModel = []CentreModel{
	{
		ID:        1,
		Name:      "Aeroctiy Gym",
		Location:  "Near IGI- Terminal 1",
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(1, 0, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        2,
		Name:      "Shankar Vihar Gym",
		Location:  "Near Shankar Vihar Metro",
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(1, 0, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
