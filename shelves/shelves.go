// Package shelves handles everything related to orders
package shelves

import (
	"fmt"
	"time"

	"github.com/projectdiscovery/expirablelru"

	"github.com/markelog/exempla/order"
	"github.com/markelog/exempla/ttl"
)

// Shelves model
type Shelves struct {
	Hot      *expirablelru.Cache `json:"hot"`
	Cold     *expirablelru.Cache `json:"cold"`
	Frozen   *expirablelru.Cache `json:"frozen"`
	Overflow *expirablelru.Cache `json:"overflow"`
}

func fullEviction(key interface{}, value interface{}) {
	fmt.Printf("Order %s was moved to overflow bucket \n", key.(string))
}

// New creates new shelves instance
func New() *Shelves {
	overflow := expirablelru.NewExpirableLRU(15, fullEviction, 0, 0)

	overflowEviction := func(key interface{}, value interface{}) {
		shelfLife := ttl.Get(value.(*order.Order).ShelfLife)

		fmt.Printf("Order %s was moved to overflow bucket \n", key.(string))

		overflow.AddWithTTL(key, value, shelfLife)
	}

	shelves := &Shelves{
		Hot:      expirablelru.NewExpirableLRU(10, overflowEviction, 0, 0),
		Cold:     expirablelru.NewExpirableLRU(10, overflowEviction, 0, 0),
		Frozen:   expirablelru.NewExpirableLRU(10, overflowEviction, 0, 0),
		Overflow: overflow,
	}

	return shelves
}

func (s *Shelves) add(order *order.Order, cache *expirablelru.Cache, cap int, shelfLife time.Duration) {
	if cache.Len() == cap {
		s.Overflow.AddWithTTL(order.ID, order, shelfLife)
	} else {
		s.Hot.AddWithTTL(order.ID, order, shelfLife)
	}
}

// Add adds order to shelves
func (s *Shelves) Add(order *order.Order) {
	shelfLife := ttl.Get(order.ShelfLife)

	switch order.Temperature {
	case "hot":
		s.add(order, s.Hot, 10, shelfLife)
	case "cold":
		s.add(order, s.Cold, 10, shelfLife)
	case "frozen":
		s.add(order, s.Frozen, 10, shelfLife)
	}

	fmt.Printf("Order %s was added to shelves with TTL %fs \n", order.ID, shelfLife.Seconds())
}
