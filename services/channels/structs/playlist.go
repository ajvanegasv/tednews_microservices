package structs

import (
	"encoding/json"
	"time"
)

func UnmarshalPlaylistResponse(data []byte) (PlaylistResponse, error) {
	var r PlaylistResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlaylistResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PlaylistResponse struct {
	Kind          string           `json:"kind"`
	Etag          string           `json:"etag"`
	NextPageToken string           `json:"nextPageToken"`
	Items         []Playlist   `json:"items"`
	PageInfo      PlaylistPageInfo `json:"pageInfo"`
}

type Playlist struct {
	Kind           string                 `json:"kind"`
	Etag           string                 `json:"etag"`
	ID             string                 `json:"id"`
	Snippet        PlaylistSnippet        `json:"snippet"`
	ContentDetails PlaylistContentDetails `json:"contentDetails"`
}

type PlaylistContentDetails struct {
	VideoID          string    `json:"videoId"`
	VideoPublishedAt time.Time `json:"videoPublishedAt"`
}

type PlaylistSnippet struct {
	PublishedAt            time.Time          `json:"publishedAt"`
	ChannelID              string             `json:"channelId"`
	Title                  string             `json:"title"`
	Description            string             `json:"description"`
	Thumbnails             PlaylistThumbnails `json:"thumbnails"`
	ChannelTitle           string             `json:"channelTitle"`
	PlaylistID             string             `json:"playlistId"`
	Position               int64              `json:"position"`
	ResourceID             PlaylistResourceID `json:"resourceId"`
	VideoOwnerChannelTitle string             `json:"videoOwnerChannelTitle"`
	VideoOwnerChannelID    string             `json:"videoOwnerChannelId"`
}

type PlaylistResourceID struct {
	Kind    string `json:"kind"`
	VideoID string `json:"videoId"`
}

type PlaylistThumbnails struct {
	Default  ChannelDefault `json:"default"`
	Medium   ChannelDefault `json:"medium"`
	High     ChannelDefault `json:"high"`
	Standard ChannelDefault `json:"standard"`
	Maxres   ChannelDefault `json:"maxres"`
}

type PlaylistDefault struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type PlaylistPageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}
