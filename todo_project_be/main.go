package main

import (
	"todo_project_be/config"
	"todo_project_be/routes"
)

func main() {
	config.InitDB()

	router := routes.SetupRouter()

	router.Run(":8080")
}
