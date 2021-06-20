package util

import (
	"strconv"
)

// GetPageNumber - get page number for pagination
func GetPageNumber(pageStr string) int {
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		return 0
	}
	return page - 1
}
