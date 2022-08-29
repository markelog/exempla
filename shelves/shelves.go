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

	shelves := &Shelves{}

	overflowEviction := func(key interface{}, value interface{}) {
		ord := value.(*order.Order)
		shelfLife := ttl.Get(ord.ShelfLife)
		added := false

		switch ord.Temperature {
		case "hot":
			added = shelves.add(ord, shelves.Hot, 10, shelfLife)
		case "cold":
			added = shelves.add(ord, shelves.Cold, 10, shelfLife)
		case "frozen":
			added = shelves.add(ord, shelves.Frozen, 10, shelfLife)
		}

		if added {
			fmt.Printf(`Order %s was readded to shelves with TTL %fs \n`, ord.ID, shelfLife.Seconds())
		} else {
			fmt.Printf(`Order %s was removed from shelves\n`, ord.ID)
		}
	}

	hot := expirablelru.NewExpirableLRU(10, overflowEviction, 0, 0)
	cold := expirablelru.NewExpirableLRU(10, overflowEviction, 0, 0)
	frozen := expirablelru.NewExpirableLRU(10, overflowEviction, 0, 0)

	shelves.Hot = hot
	shelves.Cold = cold
	shelves.Frozen = frozen
	shelves.Overflow = overflow

	return shelves
}

func (s *Shelves) add(order *order.Order, cache *expirablelru.Cache, cap int, shelfLife time.Duration) bool {
	if cache.Len() <= cap {
		cache.AddWithTTL(order.ID, order, shelfLife)
		return true
	}

	return false
}

// Add adds order to shelves
func (s *Shelves) Add(order *order.Order) {
	shelfLife := ttl.Get(order.ShelfLife)
	added := false

	switch order.Temperature {
	case "hot":
		added = s.add(order, s.Hot, 10, shelfLife)
	case "cold":
		added = s.add(order, s.Cold, 10, shelfLife)
	case "frozen":
		added = s.add(order, s.Frozen, 10, shelfLife)
	}

	if !added {
		s.Overflow.AddWithTTL(order.ID, order, shelfLife)
	}

	fmt.Printf("Order %s was added to shelves with TTL %fs \n", order.ID, shelfLife.Seconds())
}

// Take checks if order is exist and removes it
func (s *Shelves) Take(order *order.Order) bool {
	exist := false

	switch order.Temperature {
	case "hot":
		_, exist = s.Hot.Peek(order.ID)
		if _, exist = s.Hot.Peek(order.ID); exist {
			s.Hot.Remove(order.ID)
		}
	case "cold":
		if _, exist = s.Cold.Peek(order.ID); exist {
			s.Cold.Remove(order.ID)
		}
	case "frozen":
		if _, exist = s.Frozen.Peek(order.ID); exist {
			s.Frozen.Remove(order.ID)
		}
	}

	if !exist {
		if _, exist = s.Overflow.Peek(order.ID); exist {
			s.Overflow.Remove(order.ID)
		}
	}

	return exist
}
