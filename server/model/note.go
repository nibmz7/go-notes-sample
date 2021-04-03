package model

import (
	"encoding/json"
)

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (note *Note) ToString() (string, error) {
	bytes, error := json.Marshal(note)
	return string(bytes), error
}
