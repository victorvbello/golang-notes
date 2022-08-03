// Original code from https://github.com/aybabtme/orderedjson
package orderedjson

import (
	"bytes"
	"encoding/json"
	"errors"

	"gonotes/utils/json_translate/flatjson"
)

type Map []MapEntry

type MapEntry struct {
	Key   json.RawMessage
	Value json.RawMessage
}

func (m *Map) UnmarshalJSON(data []byte) error {
	_, found, err := flatjson.ScanObject(data, 0, &flatjson.Callbacks{
		OnRaw: func(name, value flatjson.Pos) {
			entry := MapEntry{
				Key:   json.RawMessage(name.Bytes(data)),
				Value: json.RawMessage(value.Bytes(data)),
			}
			(*m) = append((*m), entry)
		},
	})
	if err != nil {
		return err
	}
	if !found {
		return errors.New("expected an object but none found")
	}
	return nil
}

func (m Map) MarshalJSON() ([]byte, error) {
	out := bytes.NewBuffer(nil)
	out.WriteRune('{')
	for i, kv := range m {
		if i != 0 {
			out.WriteRune(',')
		}
		out.Write(kv.Key)
		out.WriteRune(':')
		out.Write(kv.Value)
	}
	out.WriteRune('}')
	return out.Bytes(), nil
}
