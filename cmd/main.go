package main

import (
	"orders/config"
	"orders/internal"
)

func main() {
	cfg := config.New()
	server := internal.NewServer(cfg)
	server.Run()
}
