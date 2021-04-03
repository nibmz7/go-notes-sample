package router

import (
	"github.com/gin-gonic/gin"
	"nibmz7/go-notes-sample/server/controller"
)

func SetupServer() *gin.Engine {
	router := gin.Default()
	router.POST(controller.ApiNote, controller.NoteController.PostNote)
	return router
}
