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

type Playlist structs.Playlist

func (p Playlist) GetPlaylistByIdFromYoutubeAPI(id []string) ([]Playlist, error) {
	conf := config.GetConfig()
	youtubeApi := conf.GetString("server.youtube_api_url")
	youtubeApiKey := conf.GetString("server.youtube_api_key")
	endpoint := youtubeApi + "/playlistItems?part=snippet&part=contentDetails&id=" + strings.Join(id, ",") + "&key=" + youtubeApiKey

	res, getErr := http.Get(endpoint)

	if getErr != nil {
		return []Playlist{}, getErr
	}

	defer res.Body.Close()
	body, ioErr := io.ReadAll(res.Body)

	if ioErr != nil {
		return []Playlist{}, ioErr
	}

	apiResponse, err := structs.UnmarshalPlaylistResponse(body)

	if err != nil {
		return []Playlist{}, err
	}

	var playlists []Playlist
	for _, item := range apiResponse.Items {
		playlists = append(playlists, Playlist(item))
	}

	return playlists, nil
}

func (p *Playlist) Save() error {
	redisDb := database.GetRedisDb()
	ctx := context.Background()

	data, err := p.MarshalBinary()
	if err != nil {
		return err
	}

	_, err = redisDb.HSet(ctx, "playlists", p.ID, data).Result()
	if err != nil {
		return err
	}

	return nil
}

func (p *Playlist) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Playlist) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}

func (p Playlist) GetAllPlaylists() ([]Playlist, error) {
	redisDb := database.GetRedisDb()
	ctx := context.Background()
	playlists, err := redisDb.HGetAll(ctx, "playlists").Result()

	if err != nil {
		return []Playlist{}, err
	}

	var result []Playlist
	for _, playlist := range playlists {
		var p Playlist
		err := p.UnmarshalBinary([]byte(playlist))
		if err != nil {
			return []Playlist{}, err
		}
		result = append(result, p)
	}

	return result, nil
}