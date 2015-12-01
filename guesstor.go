package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	amount := flag.Int("a", 1, "amount of links to generate")
	top := flag.String("t", "onion", "Top domain")
	up := flag.Bool("u", false, "Check if the onion domain is up")
	list := "abcdefghijklmnopqrstuvwxyz234567"

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	for j := 0; j < *amount; j++ {
		domain := make([]byte, 16)
		for i := range domain {
			domain[i] = list[rand.Intn(len(list))]
		}
		if *up == true {
			urlUP := fmt.Sprintf("https://ahmia.fi/address/%s/status", domain)
			resp, err := http.Get(urlUP)
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			var status string
			if string(body) == "up" {
				status = "up"
			} else {
				status = "down"
			}
			fmt.Printf("http://%s.onion -- STATUS: %s\n", urlUP, status)
		} else {
			fmt.Printf("http://%s.%s\n", domain, *top)
		}
	}
}
