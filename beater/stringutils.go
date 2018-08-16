package beater

import (
    "strings"
)

//
// String exists in array ?
//
// case in-sensitive
//
func existsIn(str string, haystack []string) bool {
    needle := strings.ToLower(str)
	for _, elem := range haystack {
		if strings.ToLower(elem) == needle {
			return true
		}
	}
	return false
}
