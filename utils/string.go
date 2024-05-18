package utils

import "strconv"

func StringToUint(s string) uint {
	result, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}

	return uint(result)
}
