package cyoa

import (
	"encoding/json"
	cyoa "gophercize/cyoa/students/manan"
	"io"
)

// since all the JSON data is delimited the same way
// can store it all in an Options struct
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Paragraphs   []string `json:"story"`
	Options []Option `json:"options"`
}

// Will read in all the individual story arcs into a hash.
type Story map[string]Chapter

// Takes in an io.Reader returns either story or error
func JsonStory(r io.Reader) (cyoa.Story, error) {
	decoder := json.NewDecoder(r)
	var stori cyoa.Story
	if err := decoder.Decode(&stori); err != nil {
		// if there is an error return nil for story and the error
		return nil, err
	}
	return stori, nil
}