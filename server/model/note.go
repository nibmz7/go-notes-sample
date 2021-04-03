package model

import (
	"encoding/json"
)

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (note *Note) ToString() (string, error) {
	bytes, error := json.Marshal(note)
	return string(bytes), error
}
