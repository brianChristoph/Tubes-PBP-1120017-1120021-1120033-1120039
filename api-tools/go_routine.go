package api

import (
	c "github.com/Tubes-PBP/controllers"
)

// Daily Func
func Routine() {
	go c.DeleteUserPeriodically()
	// go c.DeleteMovieSchedulePeriodically()
}
