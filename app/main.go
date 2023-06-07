package main

import (
	// "project-B/app/controller"
	"project-B/app/controller"
	"project-B/app/repo/db"
)

func main() {
	db.Init()
	controller.Start()
}
