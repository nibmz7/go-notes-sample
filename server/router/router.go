package router

import (
	"github.com/gin-gonic/gin"
	"nibmz7/go-notes-sample/server/controller"
)

func SetupServer() *gin.Engine {
	router := gin.Default()
	noteController := controller.NewNoteController()
	router.POST(controller.ApiNote, noteController.PostNote)
	router.GET(controller.ApiNote, noteController.ListenNote)
	return router
}
