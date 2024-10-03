package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cydave/gintemplate/internal/config"
	"github.com/cydave/gintemplate/internal/database"
	"github.com/cydave/gintemplate/internal/server"
)

func main() {
	// Prepare configuration.
	cfg, err := config.InitFrom(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Prepare database.
	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
	_ = database.Migrate(db)

	// Prepare server
	srv, err := server.Init()
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("%s:%d", cfg.GetString("server.host"), cfg.GetInt("server.port"))
	if err := srv.Run(addr); err != nil {
		log.Fatal(err)
	}
}
