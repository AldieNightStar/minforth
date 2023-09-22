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

func Lex(src string) []string {
	return removeEmty(
		strings.Split(strings.ReplaceAll(src, "\t", " "), " "),
	)
}

func LexNumber(tok string) (n int64, ok bool) {
	n, err := strconv.ParseInt(tok, 10, 32)
	if err != nil {
		return 0, false
	}
	return n, true
}
