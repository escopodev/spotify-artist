package models

import "fmt"

type Artist struct {
	ID         int64    `json:"id"`
	UID        string   `json:"uid"`
	Name       string   `json:"name"`
	Popularity int64    `json:"popularity"`
	Genres     []string `json:"genres"`

	URI          string   `json:"uri,omitempty"`
	Href         string   `json:"href,omitempty"`
	ExternalURLs []string `json:"external_urls,omitempty"`
	Followers    int      `json:"followers,omitempty"`

	Images []Image `json:"images"`
	Users  []User  `json:"users"`
}

func (a *Artist) Compute() {
	a.URI = fmt.Sprintf("spotify:artist:%s", a.UID)
	a.Href = fmt.Sprintf("https://api.spotify.com/v1/artists/%s", a.UID)
	a.ExternalURLs = []string{fmt.Sprintf("https://open.spotify.com/artist/%s", a.UID)}
	a.Followers = len(a.Users)
}
