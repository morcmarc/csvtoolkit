package types

/* Credit goes to nathanwdavis */

import (
	"encoding/csv"
	"io"
	"strconv"
)

const (
	STRING = "string"
	INT    = int64(999)
	FLOAT  = float64(1.2)
	BOOL   = false
)

func Infer(csv *csv.Reader, fields []string, numLines int) (map[string]interface{}, error) {

	template := make(map[string]interface{})
	learners := make([]*Learner, len(fields))
	for i := 0; i < len(learners); i++ {
		learners[i] = NewLearner()
	}
	for i := 0; i < numLines; i++ {
		line, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		for j, _ := range fields {
			learners[j].Feed(line[j])
		}
	}
	for i, f := range fields {
		exampleVal := learners[i].BestGuess()
		template[f] = exampleVal
	}
	return template, nil
}

type Learner struct {
	lastGuess interface{}
}

func NewLearner() *Learner {
	return new(Learner)
}

func (s *Learner) Feed(str string) {
	typedv := GuessString(str)
	exampl := exampleVal(typedv)

	switch {
	case typedv == nil:
		//do nothing
	case exampl == STRING || s.lastGuess == STRING:
		s.lastGuess = STRING
	case s.lastGuess != FLOAT && s.lastGuess != BOOL && exampl == INT:
		s.lastGuess = INT
	case s.lastGuess != FLOAT && s.lastGuess != INT && exampl == BOOL:
		s.lastGuess = BOOL
	case (s.lastGuess != BOOL && exampl == FLOAT) ||
		(exampl != s.lastGuess &&
			(exampl == FLOAT || exampl == INT) &&
			(s.lastGuess == FLOAT || s.lastGuess == INT)):
		s.lastGuess = FLOAT
	default:
		s.lastGuess = STRING
	}
}

func (s *Learner) FeedAll(strings []string) {
	for _, str := range strings {
		s.Feed(str)
	}
}

func (s *Learner) BestGuess() interface{} {
	return s.lastGuess
}

func GuessString(v string) interface{} {
	if typedv, err := strconv.ParseInt(v, 0, 64); err == nil {
		return typedv
	} else if typedv, err := strconv.ParseBool(v); err == nil {
		return typedv
	} else if typedv, err := strconv.ParseFloat(v, 64); err == nil {
		return typedv
	} else if v == "" {
		return nil
	}
	return v
}

func GuessStrings(strings []string) []interface{} {
	typedVals := make([]interface{}, len(strings))
	for idx, v := range strings {
		typedVals[idx] = GuessString(v)
	}
	return typedVals
}

func exampleVal(orig interface{}) interface{} {
	switch orig.(type) {
	case string:
		return STRING
	case int64:
		return INT
	case float64:
		return FLOAT
	case bool:
		return BOOL
	default:
		return nil
	}
}
