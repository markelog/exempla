package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/markelog/exempla/filejson"
)

/**
 * Main entry-point for the CSS coding challenge.
 *
 * This scaffolding is meant as an example on how to get started on the coding challenge. You can
 * choose to use however much of this scaffolding as you want and are free to modify all of it or
 * create their own. The libraries used by the scaffolding are only for illustrative purposes and
 * you can use whatever libraries you find useful.
 */
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
