package minforth

import (
	"strconv"
	"strings"
)

func removeEmty(arr []string) (out []string) {
	for _, el := range arr {
		if len(el) > 0 {
			out = append(out, el)
		}
	}
	return out
}

func lex(src string) []string {
	src = strings.ReplaceAll(src, "\t", " ")
	src = strings.ReplaceAll(src, "\n", " ")
	src = strings.ReplaceAll(src, "\r", " ")
	return removeEmty(strings.Split(src, " "))
}

func lexNumber(tok string) (n int64, ok bool) {
	n, err := strconv.ParseInt(tok, 10, 32)
	if err != nil {
		return 0, false
	}
	return n, true
}
