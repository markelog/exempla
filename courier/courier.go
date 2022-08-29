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

	order := d.Order
	exist := false

	switch d.Order.Temperature {
	case "hot":
		_, exist = d.Shelves.Hot.Peek(order.ID)
	case "cold":
		_, exist = d.Shelves.Cold.Peek(order.ID)
	case "frozen":
		_, exist = d.Shelves.Frozen.Peek(order.ID)
	}

	if !exist {
		_, exist = d.Shelves.Overflow.Peek(order.ID)
	}

	return exist
}
