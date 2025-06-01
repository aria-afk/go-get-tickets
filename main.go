package main

import (
	"fmt"
	"os"

	"github.com/aria-afk/go-get-tickets/pg"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "migrate" {
		fmt.Println("Running migrations...")
		pgConn, err := pg.NewPG()
		if err != nil {
			fmt.Printf("Error getting DB connection:\n%s", err)
			os.Exit(1)
		}
		err = pg.MigrateUp(pgConn)
		if err != nil {
			os.Exit(1)
		}
	}
}
