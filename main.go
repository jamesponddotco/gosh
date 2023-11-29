package main

import (
	"os"

	"git.sr.ht/~jamesponddotco/gosh/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args))
}
