package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aria-afk/go-get-tickets/pg"
	"github.com/aria-afk/go-get-tickets/utils"
)

func main() {
	utils.LoadEnv("dev.env")

	// flags
	migrate := flag.String("migrate", "", "Type of migration to run, valid options are [up/down]")
	flag.Parse()

	if *migrate != "" {
		fmt.Printf("Running migration %s\n", *migrate)
		db, err := pg.NewPG()
		if err != nil {
			fmt.Printf("Error getting DB connection:\n%s", err)
			os.Exit(1)
		}
		err = db.Migrate(*migrate)
		if err != nil {
			os.Exit(1)
		}
	}
}
