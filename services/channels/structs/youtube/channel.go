package youtube

import "encoding/json"

func UnmarshalApiResponse(data []byte) (ApiResponse, error) {
	var r ApiResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ApiResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ApiResponse struct {
	Kind     string   `json:"kind"`
	Etag     string   `json:"etag"`
	PageInfo PageInfo `json:"pageInfo"`
	Items    []Channel   `json:"items"`
}

type Channel struct {
	Kind           string         `json:"kind"`
	Etag           string         `json:"etag"`
	ID             string         `json:"id"`
	Snippet        Snippet        `json:"snippet"`
	ContentDetails ContentDetails `json:"contentDetails"`
}

type ContentDetails struct {
	RelatedPlaylists RelatedPlaylists `json:"relatedPlaylists"`
}

type RelatedPlaylists struct {
	Likes   string `json:"likes"`
	Uploads string `json:"uploads"`
}

type Snippet struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CustomURL   string     `json:"customUrl"`
	PublishedAt string     `json:"publishedAt"`
	Thumbnails  Thumbnails `json:"thumbnails"`
	Localized   Localized  `json:"localized"`
	Country     string     `json:"country"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Thumbnails struct {
	Default Default `json:"default"`
	Medium  Default `json:"medium"`
	High    Default `json:"high"`
}

type Default struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}
