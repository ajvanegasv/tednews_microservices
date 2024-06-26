package server

import (
	"github.com/ajvanegasv/tednews_microservices/services/channels/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	channelGroup := router.Group("channels")
	{
		channelHandler := new(handlers.ChannelHandler)
		channelGroup.GET("/", channelHandler.GetChannels)
		channelGroup.GET("/:id", channelHandler.GetChannel)
	}

	playlistGroup := router.Group("playlists")
	{
		playlistHandler := new(handlers.PlaylistHandler)
		playlistGroup.GET("/", playlistHandler.GetPlaylists)
		playlistGroup.GET("/:id", playlistHandler.GetPlaylistById)
		
	}

	return router
}