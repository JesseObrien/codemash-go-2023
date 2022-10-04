package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jesseobrien/codemash-go-2023/logs"
)

// see: https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

var log = logs.New()

func main() {

	http.HandleFunc("/", NoTimeout)

	srv := http.Server{
		ReadTimeout:  time.Second * 1,
		WriteTimeout: time.Second * 1,
	}

	srv.ListenAndServe()
}

func NoTimeout(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "OK")
}

func ReadTimeout(rw http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Second * 2)

	fmt.Fprint(rw, "Read Timeout")
}

func WriteTimeout(rw http.ResponseWriter, req *http.Request) {
	payload, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err.Error())
		http.Error(rw, "Failed to read body", http.StatusInternalServerError)
		return
	}

	time.Sleep(time.Second * 2)

	log.Info(string(payload))

	fmt.Fprint(rw, "Write Timeout")
}
