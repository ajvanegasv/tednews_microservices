package structs

import "encoding/json"

func UnmarshalChannelApiResponse(data []byte) (ChannelApiResponse, error) {
	var r ChannelApiResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelApiResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelApiResponse struct {
	Kind     string          `json:"kind"`
	Etag     string          `json:"etag"`
	PageInfo ChannelPageInfo `json:"pageInfo"`
	Items    []Channel       `json:"items"`
}

type Channel struct {
	Kind           string                `json:"kind"`
	Etag           string                `json:"etag"`
	ID             string                `json:"id"`
	Snippet        ChannelSnippet        `json:"snippet"`
	ContentDetails ChannelContentDetails `json:"contentDetails"`
}

type ChannelContentDetails struct {
	RelatedPlaylists ChannelRelatedPlaylists `json:"relatedPlaylists"`
}

type ChannelRelatedPlaylists struct {
	Likes   string `json:"likes"`
	Uploads string `json:"uploads"`
}

type ChannelSnippet struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	CustomURL   string            `json:"customUrl"`
	PublishedAt string            `json:"publishedAt"`
	Thumbnails  ChannelThumbnails `json:"thumbnails"`
	Localized   ChannelLocalized  `json:"localized"`
	Country     string            `json:"country"`
}

type ChannelLocalized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ChannelThumbnails struct {
	Default ChannelDefault `json:"default"`
	Medium  ChannelDefault `json:"medium"`
	High    ChannelDefault `json:"high"`
}

type ChannelDefault struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type ChannelPageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}
