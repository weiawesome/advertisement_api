package advertisement

import "time"

type GetAdvertisementResponse struct {
	Items []Item
}

type Item struct {
	Title string    `json:"title"`
	EndAt time.Time `json:"endAt"`
}
