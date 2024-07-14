package main

import (
	"log"

	"github.com/hasssanezzz/multi-service-shortner/cmd"
)

var PORT = "80"

func main() {
	cmd.InitDB()
	server := cmd.NewAPIServer(":" + PORT)

	log.Println("server started, running on port:", PORT)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
