package types

import (
	"encoding/json"
	"errors"
)

type Headers struct {
	headersMap map[string]string
}

func NewHeaders(headersMap map[string]string) Headers {
	return Headers{headersMap: headersMap}
}

func (headers Headers) GetValueByKey(key string) (string, error) {
	value, ok := headers.headersMap[key]
	if ok {
		return value, nil
	} else {
		return "None", errors.New("no key found")
	}
}

func (headers Headers) Json() ([]byte, error) {
	result, err := json.Marshal(headers.headersMap)
	return result, err
}
