package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Chargement en cours")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("\nOpération terminée!")
}
