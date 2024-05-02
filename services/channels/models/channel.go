package models

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/ajvanegasv/tednews_microservices/services/channels/config"
	"github.com/ajvanegasv/tednews_microservices/services/channels/database"
	"github.com/ajvanegasv/tednews_microservices/services/channels/structs"
)

type Channel structs.Channel
type ChannelContentDetails structs.ChannelContentDetails

func (c Channel) GetChannelByIdFromYoutubeAPI(id []string) ([]Channel, error) {
	conf := config.GetConfig()
	youtubeApi := conf.GetString("server.youtube_api_url")
	youtubeApiKey := conf.GetString("server.youtube_api_key")

	endpoint := youtubeApi + "/channels?part=snippet&part=contentDetails&id=" + strings.Join(id, ",") + "&key=" + youtubeApiKey
	res, getErr := http.Get(endpoint)

	if getErr != nil {
		return []Channel{}, getErr
	}

	defer res.Body.Close()
	body, ioErr := io.ReadAll(res.Body)

	if ioErr != nil {
		return []Channel{}, ioErr
	}

	apiResponse, err := structs.UnmarshalChannelApiResponse(body)

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

	data, err := c.MarshalBinary()
	if err != nil {
		return err
	}

	// Guardar la secuencia de bytes en Redis
	_, err = redisDb.HSet(ctx, "channels", c.ID, data).Result()
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

func (c Channel) GetChannelById(id string) (Channel, error) {
	redisDb := database.GetRedisDb()
	ctx := context.Background()
	channel, err := redisDb.HGet(ctx, "channels", id).Result()

	if err != nil {
		return Channel{}, err
	}

	var ch Channel
	json.Unmarshal([]byte(channel), &ch)

	return ch, nil

}

func (c *Channel) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Channel) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *Channel) GetChannelContentDetails() ChannelContentDetails {
	return ChannelContentDetails(c.ContentDetails)
}
