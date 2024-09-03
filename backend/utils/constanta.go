/*
This is where every const you want to define lies
*/
package utils

import "time"

func GetStringTimeNow(format string) string {

	now := time.Now().Local()

	switch format {
	case "date":
		dateStr := now.Format("2006-01-02")
		return dateStr
	case "time":
		timeStr := now.Format("15:04:05")
		return timeStr
	case "year":
		dateStr := now.Format("2006")
		return dateStr
	default:
		datetime := now.Format("2006-01-02 15:04:05")
		return datetime
	}
}
