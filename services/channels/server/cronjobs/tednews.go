package cronjobs

import (
	"log"

	"github.com/ajvanegasv/tednews_microservices/services/channels/config"
	"github.com/ajvanegasv/tednews_microservices/services/channels/models"
)

var channelModel = new(models.Channel)

func UpdateInfoTednewsChannel() {
	conf := config.GetConfig()

	channelsApiKeys := []string{conf.GetString("server.youtube_channel_id")}
	channels, err := channelModel.GetChannelByIdFromYoutubeAPI(channelsApiKeys)

	if err != nil {
		log.Println(err)
		return
	}

	if len(channels) == 0 {
		log.Println("No channels found")
		return
	}

	for _, channel := range channels {
		err := channel.Save()

		if err != nil {
			log.Println(err)
			continue
		}
	}
}
