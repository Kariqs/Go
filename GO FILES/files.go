package main

import (
	"fmt"

	"example.com/file/file"
	"example.com/file/note"
)

func main() {
	title := getUserData("Note title:")
	content := getUserData("Note content:")

	newNote, err := note.New(title, content)
	if err != nil {
		panic(err)
	}
	newNote.PrintNote()
	file.WriteToFile(*newNote, "note.json")
}

func getUserData(text string) string {
	var value string
	fmt.Print(text)
	fmt.Scanln(&value)
	return value
}
