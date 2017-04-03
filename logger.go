package main

import (
	"log"
)

func logFileMessage(message string, filename string) {
	log.Printf(message+"\n", filename)
}

func logStringError(message string, e error) {
	log.Printf(message+"\n", e)
}
