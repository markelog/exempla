// Exempla package
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/markelog/exempla/orders"
	"github.com/markelog/exempla/shelves"
)

func main() {
	inputLink := flag.String("f", "orders.json", "The path to the input json file")
	ingestRateLink := flag.Int("i", 2, "Ingest rate per second")

	flag.Parse()

	input := *inputLink
	ingestRate := *ingestRateLink

	data, err := orders.MakeOrdersFromFile(input)
	if err != nil {
		fmt.Printf("failed to parse inputs with error: %v", err)
		os.Exit(1)
	}

	readyShleves := shelves.New()
	requests := orders.New(readyShleves)

	for i := 0; i < len(data); i += ingestRate {
		end := min(i+ingestRate, len(data))

		requests.Ingest(data[i:end])
		time.Sleep(1 * time.Second)

		fmt.Printf("Ingested from %d to %d orders\n", i, end)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
