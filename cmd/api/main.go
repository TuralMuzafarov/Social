package main

import (
	"log"
	"time"

	env "github.com/TuralMuzafarov/social/internal/env"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Printf("Could not load the environment variables... \n")
		log.Println(err)
		time.Sleep(time.Second * 3)
		log.Println("System is starting with default configurations")
	}
	cfg := Config{
		addr: env.GetString("ADDR", ":8081"),
	}
	app := &Application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
