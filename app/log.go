package main

import "log"

func logProduct(source string, id int) {
	log.Printf("source: %s, product: %d", source, id)
}

func logImportStart(source string) {
	log.Printf("start: %s import", source)
}

func logImportEnd(source string) {
	log.Printf("end: %s import", source)
}
