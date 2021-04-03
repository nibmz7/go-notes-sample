package main

import (
	"nibmz7/go-notes-sample/server/router"
)

func main() {
	router.SetupServer().Run(":3000")
}
