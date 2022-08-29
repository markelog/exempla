package shelves_test

import (
	"testing"

	"github.com/markelog/exempla/order"
	"github.com/markelog/exempla/shelves"
)

func TestShelves_ShortTTL(t *testing.T) {

	ord := &order.Order{
		ID:          "a8cfcb76-7f24-4420-a5ba-d46dd77bdffd",
		Name:        "Banana Split",
		Temperature: "frozen",
		ShelfLife:   0,
		DecayRate:   0.63,
	}

	s := shelves.New()

	s.Add(ord)

	if s.Take(ord) {
		t.Fatalf("Order still in the shelves")
	}
}

func TestShelves_LongTTL(t *testing.T) {
	ord := &order.Order{
		ID:          "a8cfcb76-7f24-4420-a5ba-d46dd77bdffd",
		Name:        "Banana Split",
		Temperature: "frozen",
		ShelfLife:   10000,
		DecayRate:   0.63,
	}

	s := shelves.New()

	s.Add(ord)

	if !s.Take(ord) {
		t.Fatalf("Order is not in the shelves")
	}
}
