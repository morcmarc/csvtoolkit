package query

import (
	"log"
)

func Keys(r CsvReader) []string {
	keys, err := r.Read()
	if err != nil {
		log.Fatalf("Could not read input: %s", err)
	}
	return keys
}
