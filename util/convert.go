package util

import (
	"strconv"
)

func ToInt(v interface{}) int {
	switch v.(type) {
	case int:
		return v.(int)
	case int32:
		return int(v.(int32))
	case int64:
		return int(v.(int64))
	case string:
		s, err := strconv.Atoi(v.(string))
		Check(err)
		return s
	case bool:
		s := v.(bool)
		if s {
			return 1
		}
		return 0
	case float64:
		return int(v.(float64))
	case float32:
		return int(v.(float32))
	//TODO: add other type...
	default:
		return 0
	}
}

func ToFloat64(v interface{}) float64 {
	switch v.(type) {
	case float64:
		return v.(float64)
	case float32:
		return float64(v.(float32))
	case int:
		return float64(v.(int))
	case string:
		s, err := strconv.ParseFloat(v.(string), 64)
		Check(err)
		return s
	//TODO: add other type...
	default:
		return 0
	}
}

func ToHex(v interface{}) int {
	switch v.(type) {
	case int:
		return ToInt(strconv.FormatInt(int64(v.(int)), 16))
	case string:
		return ToInt(strconv.FormatInt(int64(ToInt(v.(string))), 16))
	case float64:
		return ToInt(strconv.FormatInt(int64(v.(float64)), 16))
	//TODO: add other type...
	default:
		return 0
	}
}

func ToString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return strconv.Itoa(v.(int))
	case []byte:
		return string(v.([]byte))
	default:
		return ""
	}
}
