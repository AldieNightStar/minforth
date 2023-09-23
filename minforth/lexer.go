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

func lexLabel(tok string) string {
	if len(tok) < 2 {
		return ""
	}
	if strings.HasSuffix(tok, ":") {
		return tok[0 : len(tok)-1]
	}
	return ""
}

func lexJumpingToken(tok string) string {
	if len(tok) < 2 {
		return ""
	}
	if strings.HasSuffix(tok, "!") {
		return tok[0 : len(tok)-1]
	}
	return ""
}

func lexVariableGetter(tok string) string {
	if len(tok) < 2 {
		return ""
	}
	if strings.HasPrefix(tok, "$") && !strings.HasPrefix(tok, "$$") {
		return tok[1:]
	}
	return ""
}

func lexVariableSetter(tok string) string {
	if len(tok) < 2 {
		return ""
	}
	if strings.HasPrefix(tok, "=") {
		return tok[1:]
	}
	return ""
}

// token:   :b=block1
func lexConstantSetter(tok string) (k, v string) {
	if len(tok) < 2 {
		return "", ""
	}
	if !strings.HasPrefix(tok, ":") {
		return "", ""
	}
	if !strings.Contains(tok, "=") {
		return "", ""
	}
	arr := strings.SplitN(tok, "=", 2)
	if len(arr) != 2 {
		return "", ""
	}
	return arr[0][1:], arr[1]
}
