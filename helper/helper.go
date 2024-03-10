package helper

import "strings"

func QueryAdjuster(q string) string {
	nq := strings.Replace(q, " ", "+", -1)
	return nq
}
