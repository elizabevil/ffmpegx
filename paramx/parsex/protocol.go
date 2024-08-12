package parsex

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func Options(r any) (string, error) {
	marshal, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	var m map[string]any
	err = json.Unmarshal(marshal, &m)
	if err != nil {
		return "", err
	}
	values := make(url.Values)
	for key, value := range m {
		values.Set(key, fmt.Sprintf("%v", value))
	}
	return values.Encode(), nil
}
