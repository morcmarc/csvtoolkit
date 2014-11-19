package converter

import (
	"log"
	"strconv"

	"github.com/morcmarc/csv2json/types"
)

type Records struct {
	Fields  []string
	TypeMap map[string]interface{}
}

func NewRecords(fields []string, typeMap map[string]interface{}) *Records {
	r := &Records{
		Fields:  fields,
		TypeMap: typeMap,
	}
	return r
}

func (r *Records) Convert(line []string) map[string]interface{} {
	data := make(map[string]interface{}, len(r.Fields))

	for idx, f := range r.Fields {
		var typedv interface{}
		var err error
		switch typ := r.TypeMap[f]; {
		case line[idx] == "" && typ != types.STRING:
			typedv = nil
		case typ == types.INT:
			typedv, err = strconv.ParseInt(line[idx], 0, 64)
			if err != nil {
				log.Fatalf("Field in line did not match inferred type (int): %s", err)
			}
		case typ == types.BOOL:
			typedv, err = strconv.ParseBool(line[idx])
			if err != nil {
				log.Fatalf("Field in line did not match inferred type (bool): %s", err)
			}
		case typ == types.FLOAT:
			typedv, err = strconv.ParseFloat(line[idx], 64)
			if err != nil {
				log.Fatalf("Field in line did not match inferred type (float): %s", err)
			}
		default:
			typedv = line[idx]
		}

		data[f] = typedv
	}

	return data
}
