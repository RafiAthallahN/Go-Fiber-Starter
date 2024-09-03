/*
This is where everything about convert lies
*/
package utils

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
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
ConvertStringToTime used for convert date-formatted string to time.Time
It'll return nil if the argument is invalid
*/
func ConvertStringToTime(date string) *time.Time {
	convertedTime, errConvertedTime := time.Parse("02-01-2006 15:04:05", date)
	if errConvertedTime != nil {
		logrus.Errorf("Can't convert \"%v\": %v", date, errConvertedTime)
		return nil
	}
	return &convertedTime
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
