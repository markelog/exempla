package courier

import (
	"math/rand"
	"time"

	"github.com/markelog/exempla/order"
	"github.com/markelog/exempla/shelves"
)

// Delivery is a courier delivery model
type Delivery struct {
	Order   *order.Order
	Shelves *shelves.Shelves
}

// New creates new courier instance
func New(order *order.Order, shelves *shelves.Shelves) *Delivery {
	return &Delivery{
		Order:   order,
		Shelves: shelves,
	}
}

// Pickup picks up order
func (d *Delivery) Pickup() bool {
	time.Sleep(time.Duration(rand.Intn(6)) * time.Second)

	return d.Shelves.Take(d.Order)
}
