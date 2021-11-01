package main

import (
	"GoCrud/api"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}
func run() error {
	server := api.NewServer()

	err := server.Run()
	if err != nil {
		return err
	}
	return nil
}
