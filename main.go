package main

import (
	"commentservice/cmd"
	"commentservice/config"
	"commentservice/data"
	"log"
	"os"
	"os/signal"

	"github.com/Smart-Pot/pkg/adapter/amqp"
)

func main() {
	config.ReadConfig()
	amqp.Set("amqp://guest:guest@localhost:5672")
	data.DatabaseConnection()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		if err := cmd.Execute(); err != nil {
			log.Fatal(err)
		}
	}()
	sig := <-c
	log.Println("GOT SIGNAL: " + sig.String())
}
