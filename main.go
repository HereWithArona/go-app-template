package main

import (
	"fmt"
	"inter-knot-server/config"
)

func main() {
	config := config.Get()

	fmt.Printf("Hello %v!\n", config.Name)
}
