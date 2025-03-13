package file

import (
	"encoding/json"
	"os"

	"example.com/file/note"
)

func WriteToFile(newNote note.Note, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	json, err := json.Marshal(newNote)
	if err != nil {
		panic(err)
	}
	os.WriteFile(fileName, json, 0644)
}
