package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gophercize/cyoa"
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


	// pass in an IO.reader (any file would be of that type)
	decoder := json.NewDecoder(fileHandle)

	var story cyoa.Story

	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}

	// %+v will print out the entire struct, the + will print the fields too
	// good for printing a struct
	fmt.Printf("%+v\n", story)
}