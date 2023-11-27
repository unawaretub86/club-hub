package main

import (
	"log"

	"github.com/unawaretub86/club-hub/server"
)

func main() {
	initiator := server.NewInitiator()

	initiator.InitDB()
	initiator.InitService()
	initiator.InitRouter()
	initiator.GetDB()

	router := initiator.GetRouter()

	port := "8080"

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Server init error:", err.Error())
	}
}
