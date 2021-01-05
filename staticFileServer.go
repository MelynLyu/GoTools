package main

import (
	"flag"
	"log"
	"net/http"
)

var port, dir String

func Init() {
	flag.StringVar(&port, "p", "9709", "port")
	flag.StringVar(&dir, "d", "/home/melyn/", "directory")
}

func main() {
	Init()
	flag.Parse()

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/home/melyn/", http.StripPrefix("/home/melyn/", fs))
	
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
