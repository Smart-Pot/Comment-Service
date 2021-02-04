package main

import (
	"commentservice/cmd"
	"commentservice/config"
	"log"
)

func main() {
	config.ReadConfig()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
