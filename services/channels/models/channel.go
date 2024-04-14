package models

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"github.com/ajvanegasv/tednews_microservices/services/tednews/config"
	"github.com/ajvanegasv/tednews_microservices/services/tednews/database"
	"github.com/ajvanegasv/tednews_microservices/services/tednews/structs/youtube"
)

type Channel youtube.Channel

func (c Channel) GetChannelByIdFromYoutubeAPI(id []string) ([]Channel, error) {
	conf := config.GetConfig()
	youtubeApi := conf.GetString("server.youtube_api_url")
	youtubeApiKey := conf.GetString("server.youtube_api_key")

	res, getErr := http.Get(youtubeApi + "/channels?part=snippet&part=contentDetails&id=" + strings.Join(id, ",") + "&key=" + youtubeApiKey)

	if getErr != nil {
		return []Channel{}, getErr
	}

	defer res.Body.Close()
	body, ioErr := io.ReadAll(res.Body)

	if ioErr != nil {
		return []Channel{}, ioErr
	}

	apiResponse, err := youtube.UnmarshalApiResponse(body)
	
	if err != nil {
		return []Channel{}, err
	}

	var channels []Channel	
	for _, item := range apiResponse.Items {
		channels = append(channels, Channel(item))
	}

	return channels, nil
}

func (c *Channel) Save() error {
	redisDb := database.GetRedisDb()
	ctx := context.Background()
	_, err := redisDb.HSet(ctx, "channels", c.ID, c).Result()

	if err != nil {
		return err
	}

	return nil
}

func (c Channel) GetAllChannels() ([]Channel, error) {
	redisDb := database.GetRedisDb()
	ctx := context.Background()
	channels, err := redisDb.HGetAll(ctx, "channels").Result()

	if err != nil {
		return []Channel{}, err
	}

	var channelList []Channel
	for _, channel := range channels {
		var c Channel
		json.Unmarshal([]byte(channel), &c)
		channelList = append(channelList, c)
	}

	return channelList, nil
}