package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	amount := flag.Int("a", 1, "amount of links to generate")
	top := flag.String("t", "onion", "Top domain")
	list := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUWXYZ1234567890"

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	for j := 0; j < *amount; j++ {
		domain := make([]byte, 16)
		for i := range domain {
			domain[i] = list[rand.Intn(len(list))]
		}
		fmt.Printf("http://%s.%s\n", domain, *top)
	}
}
