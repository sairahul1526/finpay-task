package util

import (
	"fmt"
	"strconv"
)

// GetPageNumber - get page number for pagination
func GetPageNumber(pageStr string) int {
	page, _ := strconv.Atoi(pageStr)
	fmt.Println("page number ", page)
	if page <= 0 {
		return 0
	}
	return page - 1
}
