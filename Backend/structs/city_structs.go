package structs

// CitySearchRequest represents the expected JSON body for city search requests.
type CitySearchRequest struct {
	City string `json:"city"`
}

// CityResult represents the structure of the city search results.
type CityResult struct {
	City    string `bson:"city" json:"city"`
	Country string `bson:"country" json:"country"`
	Airport string `bson:"name" json:"airport"`
	IATA    string `bson:"iata_code" json:"iata"`
}
