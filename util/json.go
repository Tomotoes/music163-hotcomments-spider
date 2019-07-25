package util

import (
	"encoding/json"
)

func ToJSON(v interface{}) string{
	s,err := json.Marshal(v)
	Check(err)
	return ToString(s)
}