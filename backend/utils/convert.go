/*
This is where everything about convert lies
*/
package utils

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

// EncodeBase64 function encode everything to base64 by convert the data to string then encode it to base64
func EncodeBase64(data any) string {
	byte := []byte(fmt.Sprint(data))
	return base64.StdEncoding.EncodeToString(byte)
}

// EncodeBase64 function decode base64  string to string
func DecodeBase64(data string) string {
	decodedValue, _ := base64.StdEncoding.DecodeString(data)
	return string(decodedValue)
}

/*
PrimitiveCrossConvert function can only be used accross non pointer primitives data type
Pass "to" argument with "string" | "int" | "float64". Otherwise, it'll return nil
*/
func PrimitiveCrossConvert(value any, to string) any {
	switch to {
	case "string":
		return fmt.Sprint(value)
	case "int":
		switch v := value.(type) {
		case string:
			if i, err := strconv.Atoi(v); err == nil {
				return i
			}
		case float64:
			return int(v)
		case int:
			return v
		}
	case "float64":
		switch v := value.(type) {
		case string:
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				return f
			}
		case int:
			return float64(v)
		case float64:
			return v
		}
	default:
		return nil
	}
	return nil
}
