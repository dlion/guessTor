package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Url    string
	Status string
}

func main() {
	amount := flag.Int("a", 1, "amount of links to generate")
	top := flag.String("t", "onion", "Top domain")
	up := flag.Bool("u", false, "Check if the onion domain is up")
	list := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUWXYZ1234567890"

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	for j := 0; j < *amount; j++ {
		domain := make([]byte, 16)
		for i := range domain {
			domain[i] = list[rand.Intn(len(list))]
		}
		if *up == true {
			urlUP := fmt.Sprintf("http://ishsup.in/?a=%s.%s&format=json", domain, *top)
			resp, err := http.Get(urlUP)
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			var data Response
			err = json.Unmarshal([]byte(body), &data)
			if err != nil {
				panic(err)
			}
			var status string
			if data.Status == "200" {
				status = "up"
			} else {
				status = "down"
			}
			fmt.Printf("http://%s -- STATUS: %s\n", data.Url, status)
		} else {
			fmt.Printf("http://%s.%s\n", domain, *top)
		}
	}
}
