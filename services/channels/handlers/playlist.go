package handlers

import (
	"github.com/ajvanegasv/tednews_microservices/services/channels/models"
	"github.com/gin-gonic/gin"
)

type PlaylistHandler struct{}
var playlistModel = new(models.Playlist)

func (ph PlaylistHandler) GetPlaylists(c *gin.Context) {
	filters := []map[string]string{
		{"channelId": c.Query("channelId")},
	}

	playlists, err := playlistModel.GetAllPlaylists(filters)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, playlists)
}

func (ph PlaylistHandler) GetPlaylistById(c *gin.Context) {
	id := c.Param("id")
	playlist, err := playlistModel.GetPlaylistById(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, playlist)
}
