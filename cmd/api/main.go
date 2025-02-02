package main

import (
	"log"
	"time"

	"github.com/TuralMuzafarov/social/internal/db"
	env "github.com/TuralMuzafarov/social/internal/env"
	"github.com/TuralMuzafarov/social/internal/store"
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
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("Database Connection Established!")
	store := store.NewStore(db)

	app := &Application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
