package cronjobs

import (
	"log"

	"github.com/ajvanegasv/tednews_microservices/services/channels/config"
	"github.com/ajvanegasv/tednews_microservices/services/channels/models"
)

var channelModel = new(models.Channel)
var playlistModel = new(models.Playlist)

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

func UpdateInfoPlaylist() {
	channels, error := channelModel.GetAllChannels()

	if (error != nil) {
		log.Println(error)
		return
	}

	for _, channel := range channels {
		contentDetail := channel.GetChannelContentDetails()

		playlists, err  := playlistModel.GetPlaylistByIdFromYoutubeAPI([]string{contentDetail.RelatedPlaylists.Uploads})

		if err != nil {
			log.Println(err)
			continue
		}

		for _, playlist := range playlists {
			err := playlist.Save()

			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}


