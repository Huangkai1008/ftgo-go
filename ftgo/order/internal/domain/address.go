package domain

// Address represents an address resource.
type Address struct {
	// Required: true
	Street string `json:"street"`

	// Required: true
	City string `json:"city"`

	// Required: true
	State string `json:"state"`

	// Required: true
	Zip string `json:"zip"`
}
