package main

// https://gobyexample.com/command-line-flags
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"marc-util-go/marcparser"
)

func main() {

	// Get file from command line
	// ---------------------------------------------
	filePtr := flag.String(
		"f", "", "file containing a marc header to read")
	flag.Parse()

	if *filePtr == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Read file
	// ---------------------------------------------
	// fmt.Println("file: ", *filePtr)

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mh := marcparser.Init()

	for scanner.Scan() {
		fmt.Println(len(scanner.Text()))
		headerstring := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal("Shouldn't see an error scanning a string")
		}

		fmt.Println(headerstring)

		mh.Parser(headerstring)
	}
	mh.PrintHeader()
}
