package service

import (
	"github.com/google/uuid"
	"nibmz7/go-notes-sample/server/model"
)

type NoteEvent struct {
	Type string
	Data interface{}
}

type NoteClient struct {
	Channel chan NoteEvent
}

func NewNoteClient() *NoteClient {
	return &NoteClient{Channel: make(chan NoteEvent)}
}

type noteService struct {
	notes   map[string]model.Note
	clients map[*NoteClient]bool
}

type NoteService interface {
	AddNote(note *model.Note)
	Subscribe(client *NoteClient)
	Unsubscribe(client *NoteClient)
}

var MakeNoteService func() NoteService = newNoteService

func newNoteService() NoteService {
	return &noteService{notes: make(map[string]model.Note), clients: make(map[*NoteClient]bool)}
}

func (ns *noteService) AddNote(note *model.Note) {
	note.ID = uuid.New().String()
	ns.notes[note.ID] = *note
	go func() {
		for client := range ns.clients {
			client.Channel <- NoteEvent{Type: "Added", Data: *note}
		}
	}()
}

func (ns *noteService) Subscribe(client *NoteClient) {
	ns.clients[client] = true
}

func (ns *noteService) Unsubscribe(client *NoteClient) {
	delete(ns.clients, client)
}
