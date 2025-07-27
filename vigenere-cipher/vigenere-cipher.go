package main

import (
	"fmt"
)

func main() {
	cypher := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	p := 0
	var res string
	for i := 0; i < len(cypher); i++ {
		shifted := (cypher[i] - keyword[p] + 26) % 26 + 'A'
		res += string(shifted)
		p++
		if p == len(keyword) {
			p = 0
		}
	}
	fmt.Print(res)
}
