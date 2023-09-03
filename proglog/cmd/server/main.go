package main

import (
	"log"

	"github.com/MoriTomo7315/go_distribute_oreilly/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
