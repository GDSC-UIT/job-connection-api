package main

import (
	"github.com/GDSC-UIT/job-connection-api/server"
)

func main() {
	server := server.New()
	server.ListenAndServe()
}