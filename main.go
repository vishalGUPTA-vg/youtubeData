package main

import (
	"database/sql"
	"embed"
	"fmt"
	"strings"
	db "youtubedata/config"

	"github.com/pressly/goose/v3"
	"go.uber.org/zap"

	"log"
	"os"
	config "youtubedata/config/configs"
	"youtubedata/handlers"
	job "youtubedata/job"
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
		os.Exit(0)
	}

	// runMigration := flag.String("migration", "", "Flag to check if Migrations need to Run")

	// flag.Parse()

	// if runMigration != nil && strings.ToUpper(*runMigration) == "ON" {
	runMigrations()
	// os.Exit(0)
	// }

	go job.YoutubeJob()
	r := handlers.GetRouter()
	r.Run(":8080")
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func runMigrations() {
	migrationUserDbUrl := config.Get().DatabaseURL

	if strings.TrimSpace(migrationUserDbUrl) == "" {
		log.Fatal("MIGRATIONS_USER_DB_URL is not provided")
		os.Exit(1)
	}
	db, err := sql.Open("postgres", migrationUserDbUrl)
	if err != nil {
		log.Fatal("PG DB Connection Failed", zap.Error(err))
		panic(err)
	}
	// setup database
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("Setting Goose Postgres Dialect Failed", zap.Error(err))
		panic(err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal("Goose Up Failed", zap.Error(err))
		panic(err)
	}
}
