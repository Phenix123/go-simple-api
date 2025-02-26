package main

import (
	"orders/internal"
)

func main() {
	server := internal.New("a", "8080")
	server.Run()
}
