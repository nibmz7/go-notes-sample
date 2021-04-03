package main

import (
	_ "fmt"
	"nibmz7/go-notes-sample/server/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	noteController := controller.NewNoteController()
	router.POST(controller.ApiNote, noteController.PostNote)
	router.Run(":3000")
}
