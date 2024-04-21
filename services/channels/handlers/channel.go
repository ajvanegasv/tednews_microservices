package handlers

import (
	"github.com/ajvanegasv/tednews_microservices/services/channels/models"
	"github.com/gin-gonic/gin"
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

func (ch ChannelHandler) GetChannel(c *gin.Context) {
	channel, err := channelModel.GetChannelById(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, channel)
}
