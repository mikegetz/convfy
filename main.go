package main

import (
	"fmt"
	"os"

	"github.com/MGuitar24/confconv"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}
	flag := os.Args[1]
	fileName := os.Args[2]
	validateArgs(flag, fileName)
	convert(flag, fileName)
}

func validateArgs(flag string, fileName string) {
	if flag != "-j" && flag != "-y" {
		printHelp()
		os.Exit(1)
	}
	if fileName == "" {
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Usage: confconv [flag] [file]")
	fmt.Println("Flags:")
	fmt.Println("  -j: Convert JSON to YAML")
	fmt.Println("  -y: Convert YAML to JSON")
}

func convert(flag string, fileName string) {
	fmt.Printf("Converting %s...\n", fileName)
	switch flag {
	case "-j":
		{
			err := confconv.RewriteJSONToYAML(fileName)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "-y":
		{
			err := confconv.RewriteYAMLToJSON(fileName)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
