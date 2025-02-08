package structs

// FlightSearchRequest represents the request structure for flight user search.
type FlightSearchRequest struct {
	SourceIATA      string `json:"source_iata"`
	DestinationIATA string `json:"destination_iata"`
	Airline         string `json:"airline"`
	Date            string `json:"date"`
	DepartureTime   string `json:"departure_time"`
}

// Flight represents a flight with user details.
type Flight struct {
	SourceIATA      string   `bson:"source_iata" json:"source_iata"`
	DestinationIATA string   `bson:"destination_iata" json:"destination_iata"`
	Airline         string   `bson:"airline" json:"airline"`
	Date            string   `bson:"date" json:"date"`
	DepartureTime   string   `bson:"departure_time" json:"departure_time"`
	UserEmails      []string `bson:"users" json:"users"`
}

// FlightAddUserRequest represents the request structure for adding a user to a flight.
type FlightAddUserRequest struct {
	SourceIATA      string `json:"source_iata"`
	DestinationIATA string `json:"destination_iata"`
	Airline         string `json:"airline"`
	Date            string `json:"date"`
	DepartureTime   string `json:"departure_time"`
}
