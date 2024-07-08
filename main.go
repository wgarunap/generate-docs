package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
)

var (
	filepath string
	output   string
)

func init() {
	flag.StringVar(&filepath, "filepath", "./openapi.yaml", "api spec file path")
	flag.StringVar(&output, "output", "output", "output folder name")
	flag.Parse()

}

func main() {

	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		//CDN:     "https://raw.githubusercontent.com/wgarunap/generate-docs/master/static/script/api-reference.js",
		SpecURL: filepath,
		CustomOptions: scalar.CustomOptions{
			PageTitle: "API Documentation",
		},
		DarkMode: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create the directory if it doesn't exist
	err = os.MkdirAll(output, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	filePath := output + "/index.html"
	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the HTML content to the file
	_, err = io.WriteString(file, htmlContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("HTML file saved successfully!")
}
