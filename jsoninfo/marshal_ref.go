package jsoninfo

import (
	"encoding/json"
	"reflect"
)

func MarshalRef(value string, otherwise interface{}) ([]byte, error) {
	if reflect.ValueOf(otherwise).IsNil() {
		return json.Marshal(&refProps{
			Ref: value,
		})
	}
	return json.Marshal(otherwise)
}

func UnmarshalRef(data []byte, destRef *string, destOtherwise interface{}) error {
	refProps := &refProps{}
	if err := json.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			*destRef = ref
			return nil
		}
	}
	return json.Unmarshal(data, destOtherwise)
}

type refProps struct {
	Ref string `json:"$ref,omitempty"`
}
