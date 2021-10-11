package utils

import (
	"encoding/json"
	"strconv"
)

// MapToStruct convert map to struct
func MapToStruct(m map[string]interface{}, val interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, val)
}

// StringToInt convert string to int
func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
