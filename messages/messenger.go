package messages

import (
	"fmt"
	"http"
	"log"
	"io/ioutil"
)

var uniqueNames <-chan string

func init() {
	ch := make(chan string)
	uniqueNames = ch
	go func(ch chan<- string) {
		id := 0
		for {
			ids := fmt.Sprintf("n%d", id)
			ch <- ids
			id++
		}
	}(ch)
}

func GetMessenger() (name string, chs <-chan []byte) {
	name = <-uniqueNames
	ch := make(chan []byte)
	chs = ch

	h := func(rw http.ResponseWriter, rq *http.Request) {
		fmt.Printf("%s got message\n", name)
		data, err := ioutil.ReadAll(rq.Body)
		if err != nil {
			ch <- data
		}
	}

	pattern := fmt.Sprintf("/m/%s", name)
	http.DefaultServeMux.HandleFunc(pattern, h)
	log.Println("messenger listening at", pattern)

	return
}
