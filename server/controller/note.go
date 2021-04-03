package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nibmz7/go-notes-sample/server/model"
	"nibmz7/go-notes-sample/server/service"
)

var (
	NoteController noteControllerInterface = &noteController{}
)

type noteController struct{}

type noteControllerInterface interface {
	PostNote(*gin.Context)
}

func (controller *noteController) PostNote(context *gin.Context) {
	var note model.Note
	if err := context.ShouldBindJSON(&note); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.NoteService.AddNote(note)
	context.Status(http.StatusOK)
}
