package server

import (
	"github.com/ajvanegasv/tednews_microservices/services/tednews/config"
	"github.com/ajvanegasv/tednews_microservices/services/tednews/server/cronjobs"
)

func Init() {
	config := config.GetConfig()
	cronjobs.Init()

	r := NewRouter()


	if (config.GetString("server.port") != "") {
		r.Run(config.GetString("server.port"))
	} else {
		r.Run()
	}
}