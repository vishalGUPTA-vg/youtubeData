package main

import (
	"fmt"
	db "youtubedata/config"

	"log"
	"os"
	config "youtubedata/config/configs"
	"youtubedata/handlers"
	// job "youtubedata/job"
)

func main() {

	//reading json file for env variable
	file, err := os.Open("dev.json")
	if err != nil {
		log.Println("Unable to open file. Err:", err)
		os.Exit(1)
	}
	//setting for env variable
	var cnf *config.Config
	config.ParseJSON(file, &cnf)
	config.Set(cnf)
	//starting youtube job
	fmt.Println("config ", config.Get())
	
	db.Init(&db.Config{
		URL: cnf.DatabaseURL,
	})
	if err != nil {
		log.Println("Unable to open postges connection. Err:", err)
		os.Exit(1) 
	}
	
	// go job.YoutubeJob()
	r := handlers.GetRouter()
	r.Run(":8080")
}
