// Package orders handles everything related to orders
package orders

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/markelog/exempla/courier"
	"github.com/markelog/exempla/order"
	"github.com/markelog/exempla/shelves"
)

// Orders is a collection of orders
type Orders struct {
	Orders  []*order.Order
	Shelves *shelves.Shelves
}

// New creates new orders instance
func New(shelves *shelves.Shelves) *Orders {
	return &Orders{
		Shelves: shelves,
	}
}

// Ingest ingests data to orders instance
func (o *Orders) Ingest(orders []*order.Order) {
	o.Orders = orders
	o.Prepare(orders)
}

// Prepare cooks collection of orders
func (o *Orders) Prepare(orders []*order.Order) {
	var wg sync.WaitGroup

	for _, order := range orders {
		order.Prepare()
		o.Shelves.Add(order)

		delivery := courier.New(order, o.Shelves)
		wg.Add(1)

		go func(id string) {
			defer wg.Done()

			if delivery.Pickup() {
				fmt.Printf("Successfully picked up order %s\n", id)
			} else {
				fmt.Printf("Failed to pick up order %s\n", id)
			}
		}(order.ID)
	}

	wg.Wait()
}

// MakeOrdersFromFile Makes orders from JSON file
func MakeOrdersFromFile(name string) ([]*order.Order, error) {
	jsonFile, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s with error: %v", name, err)
	}

	defer func() {
		if err := jsonFile.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file at %s with error %v", name, err)
	}

	return readFromJSONContent(bytes)
}

func readFromJSONContent(bytes []byte) ([]*order.Order, error) {
	var result []*order.Order
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, fmt.Errorf("failed to deserialize with error %v", err)
	}
	return result, nil
}
