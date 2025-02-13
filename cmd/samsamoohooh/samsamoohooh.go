package main

import (
	"fmt"
	"log"

	"github.com/roka-crew/pkg/config"
)

func main() {
	cfg, err := config.New("./configs/config.yaml")
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	fmt.Println("cfg: ", cfg)
}
