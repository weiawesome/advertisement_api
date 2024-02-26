package advertisement

import "time"

type AddAdvertisementRequest struct {
	Title     *string   `json:"title"`
	StartAt   time.Time `json:"startAt"`
	EndAt     time.Time `json:"endAt"`
	Condition Condition `json:"condition"`
}
type Condition struct {
	AgeStart *int     `json:"ageStart"`
	AgeEnd   *int     `json:"ageEnd"`
	Gender   []string `json:"gender"`
	Country  []string `json:"country"`
	Platform []string `json:"platform"`
}
