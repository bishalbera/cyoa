package cyoa

import (
	"encoding/json"
	"io"
)

/*
Story represents a Chose Your Own Adventure story.
Each key is the name of a story chapter (eg. "arc"), and
each value is a Chapter.
*/
type Story map[string]Chapter

/*
Chapter represents a CYOA story chapter (or "arc").
Each chapter includes it's title, the paragrapshs its composed
of, and options available for the reader to take at the end of the 
chapter. If the options are empty it's assumed that you've reached
the end of that particular story path.
*/
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

/*
Options represents a choice offered at the end of a story chapter.
Text is the visible text end users will see,  while Chapter field
will be the key to a chapter stored in the Story object this chapter
was found in.
*/
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}


/*
JsonStory will decode a story using the incoming reader
and encoding/json package. It's assumed that the provided 
reader has the story stored in JSON.
*/
func JsonStory(r io.Reader) (Story, error) {

	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
