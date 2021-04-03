package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
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
	
	router := gin.Default()
	router.POST(ApiNote, NoteController.PostNote)
	ts := httptest.NewServer(router)

	defer ts.Close()

	values := map[string]string{"title": "Some title", "content": "Some content"}

	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(fmt.Sprintf("%s/note", ts.URL), "application/json; charset=utf-8", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}
