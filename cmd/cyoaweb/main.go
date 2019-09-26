package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main()  {
	fileName := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *fileName)

	// open file with filehandle
	fileHandle, err := os.Open(*fileName)
	// make sure there is no errors
	// usually panicking is not a great idea...
	if err  != nil {
		panic(err)
	}

}