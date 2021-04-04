package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"nibmz7/go-notes-sample/server/model"
	"nibmz7/go-notes-sample/server/service"
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
	conn, err := wsupgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	println("CONNECTION OPENED")

	noteClient := service.NewNoteClient()
	closeChannel := make(chan bool)

	go func() {
		controller.service.Subscribe(noteClient)
	L:
		for {
			select {
			case event := <-noteClient.Channel:
				b, err := json.Marshal(event)
				if err != nil {
					fmt.Println(err)
					break L
				}
				conn.WriteMessage(websocket.TextMessage, b)
			case <-closeChannel:
				break L
			}

		}
		println("CONNECTION CLOSED")
	}()

	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				controller.service.Unsubscribe(noteClient)
				closeChannel <- true
				break
			}
		}
		println("CONNECTION ERROR")
	}()

}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
