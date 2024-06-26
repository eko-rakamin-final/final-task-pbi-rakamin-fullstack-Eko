package main

import (
	"log"
	"rakamin-go/database"
	"rakamin-go/router"
)
func main() {
	
	database.ConnectDB()
	r := router.SetupRouter()
	r.Run()
	if r == nil {
		log.Fatal("error")
	}
}
