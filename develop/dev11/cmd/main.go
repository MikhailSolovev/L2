package main

import (
	"github.com/MikhailSolovev/L2/develop/dev11/internal"
	"log"
)

func main() {
	handlers := new(internal.Handler)
	calendar := internal.NewCalendar()
	handlers.InitRoutes(calendar)

	srv := new(internal.Server)
	err := srv.Run("8000")
	if err != nil {
		log.Fatalf("server start error: %s", err.Error())
	}
}
