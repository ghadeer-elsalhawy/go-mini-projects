package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Spaceline                  Days         Trip type           Price")
	fmt.Println("---------                  ----         ---------           -----")
	spaceline := []string{"SpaceX", "Space Adventures", "Virgin Galactic"}
	stat := []string{"Round Trip", "One Way"}
	const distance = 62100000
	const sectodays = 60 * 60 * 24
	for i := 0; i < 10; i++ {
		temp := rand.Intn(3)
		sp := spaceline[temp]
		temp = rand.Intn(2)
		st := stat[temp]
		speed := rand.Intn(16) + 15
		duration := distance / speed / sectodays
		price := speed + 20
		if st == "Round Trip" {
			price *= 2
		}
		fmt.Printf("%-16s	    %d          %-10s           $%2d\n", sp, duration, st, price)
	}
}
