package centre

import "time"

type CentreModel struct {
	ID        int64
	Name      string
	Location  string
	Status    int
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
