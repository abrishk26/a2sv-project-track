package main

import (
	"github.com/abrishk26/a2sv-project-track/task4/router"
)


func main() {
	router := router.NewRouter()

	router.Run(":8080")
}
