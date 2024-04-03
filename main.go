package main

import (
	// db "youtube-data/config"

	"log"
	"os"
	config "youtubedata/config/configs"
	"youtubedata/handlers"
)

func main() {

	// db.Init(&db.Config{
	//  URL:       "",
	// })

	//reading json file for env variable
	file, err := os.Open("dev.json")
	if err != nil {
		log.Println("Unable to open file. Err:", err)
		os.Exit(1)
	}
	//setting for env variable
	var cnf *config.Config
	config.ParseJSON(file, &cnf)

	r := handlers.GetRouter()
	r.Run(":8080")
}
