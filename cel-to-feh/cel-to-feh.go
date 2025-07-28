package main

import (
	"fmt"
)
const (
	line = "===================================="
)
type celicuis float64

type fehrenheit float64

func drawtable (first string, second string) {
	fmt.Printf("%s\n", line)
	fmt.Printf("|      º%s       |       º%s         |\n", first, second)
	fmt.Printf("%s\n", line)
}
func celtofeh (c celicuis) fehrenheit {
	return fehrenheit((c * 1.8) + 32)
}
func main() {
	drawtable("C", "F")
	for i:= -40; i <= 100; i += 5 {
		feh := celtofeh(celicuis(i))
		fmt.Printf("|    º%3d       |       º%3.0f       |\n", i, feh)
	}
	fmt.Printf("%s\n", line)
}
