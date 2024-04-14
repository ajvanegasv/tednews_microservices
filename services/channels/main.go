package main

import (
	"github.com/ajvanegasv/tednews_microservices/services/tednews/database"
	"github.com/ajvanegasv/tednews_microservices/services/tednews/server"
)

func main() {
	database.Init()
	server.Init()
}