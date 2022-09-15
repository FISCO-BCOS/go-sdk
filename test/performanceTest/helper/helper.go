package helper

import (
	"time"
)

// DiffNano
func DiffNano(startTime time.Time) (diff int64) {
	diff = int64(time.Since(startTime))
	return
}

func InArrayStr(str string, arr []string) (inArray bool) {
	for _, s := range arr {
		if s == str {
			inArray = true
			break
		}
	}
	return
}
