package service

import (
	"github.com/google/uuid"
	"nibmz7/go-notes-sample/server/model"
)

var notes = make(map[string]model.Note)

var (
	NoteService noteServiceInterface = &noteService{}
)

type noteService struct{}

type noteServiceInterface interface {
	AddNote(model.Note)
}

func (notesService *noteService) AddNote(note model.Note) {
	note.ID = uuid.New().String()
	notes[note.ID] = note
}
