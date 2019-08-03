package workout

import (
	"time"
)

type WorkoutModel struct {
	ID        int64
	Name      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
