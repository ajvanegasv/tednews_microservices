package handlers

import (
	"github.com/ajvanegasv/tednews_microservices/services/channels/models"
	"github.com/gin-gonic/gin"
)

type PlaylistHandler struct{}
var playlistModel = new(models.Playlist)

func (ph PlaylistHandler) GetPlaylists(c *gin.Context) {
	playlists, err := playlistModel.GetAllPlaylists()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, playlists)
}
