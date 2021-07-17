package main

// Customer entity
type Customer struct {
	CustomerId string `json:"customerId"`
	Enabled    bool   `json:"enabled"`
	Name       string `json:"name"`
}

// Partial customer used for udpating customer entities
type CustomerUpdate struct {
	Enabled *bool   `json:"enabled,omitempty"`
	Name    *string `json:"name,omitempty"`
}
