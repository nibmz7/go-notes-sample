package service

import (
	"github.com/google/uuid"
	"nibmz7/go-notes-sample/server/model"
)

var notes = make(map[string]model.Note)

type NotesService struct{}

func NewNotesService() *NotesService {
	return &NotesService{}
}

func (notesService *NotesService) AddNote(note model.Note) {
	note.ID = uuid.New().String()
	notes[note.ID] = note
}
