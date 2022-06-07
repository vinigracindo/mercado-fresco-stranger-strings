package main

import "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server"

func main() {
	port := 8080
	server := server.NewAPIServer()
	server.Run(port)
}
