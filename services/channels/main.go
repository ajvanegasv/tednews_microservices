package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ajvanegasv/tednews_microservices/services/channels/config"
	"github.com/ajvanegasv/tednews_microservices/services/channels/database"
	"github.com/ajvanegasv/tednews_microservices/services/channels/server"
)

func main() {
	enviroment := flag.String("e", "default", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*enviroment)
	database.Init()
	server.Init()
}
