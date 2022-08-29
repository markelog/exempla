package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/markelog/exempla/filejson"
)

func main() {
	f := flag.String("f", "css_example.json", "The path to the input json file")
	flag.Parse()
	fileName := *f
	results, err := filejson.ReadFromJsonFile(fileName)
	if err != nil {
		fmt.Printf("failed to parse inputs with error: %v", err)
		os.Exit(1)
	}

	// PLACEHOLDER - real logic to use parsed data will go here
	fmt.Printf("%d items read from file %s\n", len(results), fileName)
	for _, entry := range results {
		str, err := entry.JsonString()
		if err != nil {
			fmt.Printf("failed to convert %v into json string with error: %v\n", entry, err)
			os.Exit(1)
		} else {
			fmt.Println(str)
		}
	}
}
