package service

import (
	"github.com/google/uuid"
	"nibmz7/go-notes-sample/server/model"
)

type noteService struct {
	notes map[string]model.Note
}

type NoteService interface {
	AddNote(note *model.Note)
}

var MakeNoteService func() NoteService = newNoteService

func newNoteService() NoteService {
	return &noteService{notes: make(map[string]model.Note)}
}

func (ns *noteService) AddNote(note *model.Note) {
	note.ID = uuid.New().String()
	ns.notes[note.ID] = *note
}
