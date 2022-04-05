package main

import (
	"fmt"
	"log"
	"os"

	xj "github.com/basgys/goxml2json"
)

func main() {
	lg := log.New(os.Stdout, "xml2json ", log.LstdFlags)
	fmt.Println("Starting the main application!")

	file, err := os.Open("src.xml")
	if err != nil {
		lg.Fatal("failed to read the source file")
		os.Exit(1)
	}

	json, err := xj.Convert(file)
	if err != nil {
		log.Fatal("failed to convert xml file!")
		os.Exit(1)
	}
	fmt.Println(json.String())

}
