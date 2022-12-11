package main

import (
	"inventory_app/database"
	"inventory_app/templates"
)

func main() {
	db := database.ConnectDatabase()
	templates.Menu(db)
}
