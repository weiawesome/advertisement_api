/*
The response structure for get advertisement is defined here.
*/

package advertisement

import "time"

// GetAdvertisementResponse is response for the get advertisement
type GetAdvertisementResponse struct {
	Items []Item // list of advertisements
}

// Item is an advertisement in response should have content
type Item struct {
	Title string    `json:"title"` // advertisement's title
	EndAt time.Time `json:"endAt"` // advertisement's end time
}
