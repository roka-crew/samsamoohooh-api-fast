package main

import (
	"os"

	"github.com/roka-crew/cmd/gen/token"
)

func main() {
	switch os.Args[1] {
	case "token":
		token.Run()
	default:
		return
	}
}
