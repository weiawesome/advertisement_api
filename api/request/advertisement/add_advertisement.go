/*
The request body structure for add advertisement is defined here.
*/

package advertisement

import "time"

// AddAdvertisementRequest that the request body to add advertisement
type AddAdvertisementRequest struct {
	// the advertisement's title
	Title *string `json:"title"`
	// the advertisement's start time and end time
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`

	// the advertisement's other condition
	Conditions Condition `json:"conditions"`
}

// Condition that is advertisement's additional constraint
type Condition struct {
	// the advertisement's constraint of gender
	Gender []string `json:"gender"`
	// the advertisement's constraint of country
	Country []string `json:"country"`
	// the advertisement's constraint of platform
	Platform []string `json:"platform"`
	// the advertisement's constraint of age start and end
	AgeStart *int `json:"ageStart"`
	AgeEnd   *int `json:"ageEnd"`
}
