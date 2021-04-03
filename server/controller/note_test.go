package controller

import (
	"fmt"
	"nibmz7/go-notes-sample/server/model"
	"nibmz7/go-notes-sample/server/service"
	"testing"
)

var (
	postNoteService func(note model.Note)
)

type noteServiceMock struct{}

func (notesService *noteServiceMock) AddNote(note model.Note) {
	postNoteService(note)
}

func TestPostNote(t *testing.T) {
	service.NoteService = &noteServiceMock{}
	postNoteService = func(note model.Note) {
		fmt.Println("note added")
	}
}
