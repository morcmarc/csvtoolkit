package query

import (
	"testing"
)

type csvReaderMock struct {
	CsvReader
	i int
}

func (c *csvReaderMock) Read() ([]string, error) {
	var csv map[int][]string = make(map[int][]string)
	csv[0] = []string{"A", "B", "C"}
	csv[1] = []string{"1", "3", "Hello"}
	csv[2] = []string{"2", "3", "World"}
	return csv[c.i], nil
}

func TestItReturnsKeys(t *testing.T) {
	reader := &csvReaderMock{}
	keys := Keys(reader)
	if keys[0] != "A" {
		t.Errorf("Got: %s", keys)
	}
	if keys[1] != "B" {
		t.Errorf("Got: %s", keys)
	}
	if keys[2] != "C" {
		t.Errorf("Got: %s", keys)
	}
}
