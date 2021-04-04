package service

import (
	"github.com/google/uuid"
	"nibmz7/go-notes-sample/server/model"
)

type NoteEvent struct {
	Type string
	Data interface{}
}

type noteService struct {
	notes   map[string]model.Note
	channel chan NoteEvent
}

type NoteService interface {
	AddNote(note *model.Note)
	Subscribe() chan NoteEvent
}

var MakeNoteService func() NoteService = newNoteService

func newNoteService() NoteService {
	return &noteService{notes: make(map[string]model.Note), channel: make(chan NoteEvent)}
}

func (ns *noteService) AddNote(note *model.Note) {
	note.ID = uuid.New().String()
	ns.notes[note.ID] = *note
	ns.channel <- NoteEvent{Type: "Added", Data: *note}
}

func (ns *noteService) Subscribe() chan NoteEvent {
	return ns.channel
}
