package query

import (
	"log"

	"github.com/morcmarc/csvtoolkit/utils"
)

func Keys(r utils.CsvReader) []string {
	r.Reset()
	keys, err := r.Read()
	if err != nil {
		log.Fatalf("Could not read input: %s", err)
	}
	return keys
}
