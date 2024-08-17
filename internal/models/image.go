package models

import "fmt"

type Image struct {
	ID     int64  `json:"-"`
	UID    string `json:"-"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

func (i *Image) Compute() {
	i.URL = fmt.Sprintf("https://i.scdn.co/image/%s", i.UID)
}
