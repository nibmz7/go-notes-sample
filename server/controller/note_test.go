package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"nibmz7/go-notes-sample/server/model"
	"nibmz7/go-notes-sample/server/service"
	"testing"
)

var (
	addNote   func(note *model.Note)
	subscribe func() chan service.NoteEvent
)

type mockNoteService struct {
	channel chan service.NoteEvent
}

func newMockNoteService() service.NoteService {
	return &mockNoteService{channel: make(chan service.NoteEvent)}
}

func (ns *mockNoteService) AddNote(note *model.Note) {
	addNote(note)
}

func (ns *mockNoteService) Subscribe() chan service.NoteEvent {
	return subscribe()
}

func TestPostNote(t *testing.T) {
	service.MakeNoteService = newMockNoteService
	addNote = func(note *model.Note) {
		note.ID = "123"
	}

	router := gin.Default()
	jsonBody := `{"title": "the title", "content": "Some content"}`
	req, _ := http.NewRequest(http.MethodPost, ApiNote, bytes.NewBufferString(jsonBody))
	rr := httptest.NewRecorder()
	noteController := NewNoteController()
	router.POST(ApiNote, noteController.PostNote)
	router.ServeHTTP(rr, req)

	var note model.Note
	err := json.Unmarshal(rr.Body.Bytes(), &note)

	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	if note.ID != "123" {
		t.Errorf("ID is not 124: ID value: %s\n", note.ID)

	}

}

