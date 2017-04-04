package main

import (
	"fmt"
	"os"
)

func findOpenConnections() {
	var tika = "http://127.0.0.1:9998/"
	var resp int8

	resp = testConnection(tika)
	if resp == connBAD {
		fmt.Fprintln(os.Stderr, "INFO: Tika connection not available to connect to. Check localhost:9998.")
		os.Exit(1)
	}
}
