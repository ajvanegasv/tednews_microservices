package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ajvanegasv/tednews_microservices/services/tednews/models"
)


type ChannelHandler struct{}

var channelModel = new(models.Channel)

func (ch ChannelHandler) GetChannels(c *gin.Context) {
	channels, err := channelModel.GetAllChannels()
	
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, channels)
}
