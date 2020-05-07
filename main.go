package main

import (
	"beegodemo/server"

	_ "github.com/lib/pq"
)

func main() {
	server.Run()
}
