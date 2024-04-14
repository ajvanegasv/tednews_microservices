package cronjobs

import (
	"time"
	"github.com/go-co-op/gocron"
)

func Init() {
	scheduler := gocron.NewScheduler(time.UTC) 

	scheduler.Every(1).Day().At("12:00").Do(UpdateInfoTednewsChannel)

	scheduler.StartAsync()
}