package main

import (
	"github.com/flipfit/booking"
	"github.com/flipfit/centre"
	"github.com/flipfit/constant"
	"github.com/flipfit/lib"
	"github.com/flipfit/pakage"
	"github.com/flipfit/schedule"
	"github.com/flipfit/user"
	"github.com/flipfit/workout"
)

func main() {
	constant.PARENT_DIRECTORY = lib.GetParentDirectory()

	centre.Init()
	schedule.Init()
	workout.Init()
	pakage.Init()
	booking.Init()
	user.Init()
}
