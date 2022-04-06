package api

import (
	"time"

	"github.com/go-co-op/gocron"
)

func RunBackgroundFunc() {
	cron := gocron.NewScheduler(time.UTC)

	// Everyday run func routine
	cron.Every(1).Day().Do(Routine)

	cron.StartAsync()
	cron.StartBlocking()
}
