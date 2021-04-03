package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"nibmz7/go-notes-sample/server/controller"
)

func main() {
	SetupServer().Run(":3000")
}

func SetupServer() *gin.Engine {
	router := gin.Default()
	router.POST(controller.ApiNote, controller.NoteController.PostNote)
	return router
}
