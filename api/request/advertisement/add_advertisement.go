/*
The request body structure for add advertisement is defined here.
*/

package advertisement

import "time"

// AddAdvertisementRequest that the request body to add advertisement
type AddAdvertisementRequest struct {
	Title      *string   `json:"title"`      //the advertisement's title
	StartAt    time.Time `json:"startAt"`    // the advertisement's start time
	EndAt      time.Time `json:"endAt"`      // the advertisement's end time
	Conditions Condition `json:"conditions"` // the advertisement's other condition
}

// Condition that is advertisement's additional constraint
type Condition struct {
	AgeStart *int     `json:"ageStart"` // the advertisement's constraint of age start
	AgeEnd   *int     `json:"ageEnd"`   // the advertisement's constraint of age end
	Gender   []string `json:"gender"`   // the advertisement's constraint of gender
	Country  []string `json:"country"`  // the advertisement's constraint of country
	Platform []string `json:"platform"` // the advertisement's constraint of platform

}
