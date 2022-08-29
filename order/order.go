package order

// Order model
type Order struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	// TODO: should be enum
	Temperature string  `json:"temp"`
	ShelfLife   int     `json:"shelfLife"`
	DecayRate   float64 `json:"decayRate"`
}

// Prepare cooks an order
func (o *Order) Prepare() {}
