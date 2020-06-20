package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 11222, "Port to run the server on")
	flag.Parse()
	var strPort = strconv.Itoa(port)

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(CustomNotFoundHandler)
	r.HandleFunc("/getPasswords", GetPasswordsHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:"+strPort,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}
	log.Println("Listening on port: " + strPort)
	log.Fatal(srv.ListenAndServe())
}

func GetPasswordsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving request from: " + r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("admin:mysupercrazypa$$w0rd\n"))
}

func CustomNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.RequestURI))
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 page not found"))
}