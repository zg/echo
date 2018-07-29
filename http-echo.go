package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func EchoHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.RemoteAddr + " requested " + request.URL.Path)
	body, err := ioutil.ReadAll(request.Body)
	if err == nil {
		panic(err.Error())
	} else {
		log.Println("Body: " + string(body))
	}
	request.Write(writer)
}

func main() {
	http.HandleFunc("/", EchoHandler)
	http.ListenAndServe(":3540", nil)
}
