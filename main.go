package main

import (
	"os"

	"github.com/dermidgen/BigQueryFlow/run"
)

func main() {
	if cmd := run.GetMain(); cmd == "" {
		os.Exit(1)
	}

}
