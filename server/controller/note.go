package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
