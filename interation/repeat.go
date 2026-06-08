package interation

import "strings"

func Repeat(s string, repeatedCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatedCount; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
