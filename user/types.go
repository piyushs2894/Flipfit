package user

import "time"

type UserModel struct {
	ID        int64
	UserName  string
	Password  string
	Name      string
	Email     string
	Mobile    string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
