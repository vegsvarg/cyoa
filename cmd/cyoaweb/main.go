package main

import (
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


	// pass in an IO.reader (any file would be of that type) to cyoa.JsonStory converter method
	// returns cyoa.Story object and nil for error if successful
	// returns nil for cyoa.Story object and the error if fails
	decoded, err := cyoa.JsonStory(fileHandle)

	fmt.Printf("%+v\n", decoded)
}