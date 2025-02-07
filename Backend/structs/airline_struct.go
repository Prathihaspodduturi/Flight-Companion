package structs

// AirlineSearchRequest represents the expected JSON body for airline search requests.
type AirlineSearchRequest struct {
	Airline string `json:"airline"`
}

// AirlineResult represents the structure of the airline search results.
type AirlineResult struct {
	Airline string `bson:"airline" json:"airline"`
}
