package cyoa

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