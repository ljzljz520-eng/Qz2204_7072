package main

import (
	"flag"
	"log"
	"net/http"
	"training41/api"
	"training41/service"
	"training41/store"
)

func main() {
	path := flag.String("db", "training41.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	db, e := store.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()
	log.Fatal(http.ListenAndServe(*addr, api.New(service.New(db))))
}
