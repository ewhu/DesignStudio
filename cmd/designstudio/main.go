// cmd/designstudio/main.go
package main

import (
	"flag"
	"log"
	"os"

	"designstudio/internal/designstudio"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := designstudio.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
