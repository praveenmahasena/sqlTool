package main

import (
	"log"

	"github.com/praveenmahasena/sqlTool/internal"
)

func main() {
	if err := internal.Start(); err != nil {
		log.Fatalln(err)
	}
}
