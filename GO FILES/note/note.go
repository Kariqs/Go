package note

import "errors"

type Note struct {
	Title   string `json:"title"` //Struct tags are put to set keys for json files.
	Content string `json:"content"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content must not be empty")
	}
	return &Note{Title: title, Content: content}, nil
}

func (note *Note) PrintNote() {
	println("Title: ", note.Title)
	println("Content: ", note.Content)
}
