package main

///resign
///https://www.itread01.com/content/1548030066.html
///https://www.itread01.com/content/1546425722.html

import (
	"log"
	"net/http"
)

var (
	port = "8000"
)

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

// func broadcaster() {
// 	clients := make(map[client]bool)
// 	for {
// 		select {
// 		case msg := <-message:

// 		}
// 	}

// }

func main() {
	http.HandleFunc("/", myweb)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func myweb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
}
