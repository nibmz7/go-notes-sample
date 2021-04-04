package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nibmz7/go-notes-sample/server/model"
	"nibmz7/go-notes-sample/server/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type NoteController struct {
	service service.NoteService
}

func NewNoteController() *NoteController {
	return &NoteController{service: service.MakeNoteService()}
}

func (controller *NoteController) PostNote(context *gin.Context) {
	var note model.Note
	if err := context.ShouldBindJSON(&note); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.service.AddNote(&note)
	context.JSON(http.StatusOK, note)
}

func (controller *NoteController) ListenNote(context *gin.Context) {
	// println("sddssd");
	// var note model.Note
	// if err := context.ShouldBindJSON(&note); err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	wshandler(controller, context.Writer, context.Request)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(controller *NoteController, w http.ResponseWriter, r *http.Request) {
	println("\nCONNECTED\n")
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	notesChannel := controller.service.Subscribe()
	conn.WriteMessage(websocket.PingMessage, []byte("ping"))

	println("\nCONNECTED\n")

	for {
		select {
		case event := <-notesChannel:
			b, err := json.Marshal(event)
			if err != nil {
				fmt.Println(err)
				return
			}
			conn.WriteMessage(websocket.TextMessage, b)
		}
	}
}
