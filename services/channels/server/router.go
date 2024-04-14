package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ajvanegasv/tednews_microservices/services/tednews/handlers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	
	channelHandler := new(handlers.ChannelHandler)
	router.GET("/channels", channelHandler.GetChannels)

	return router
}