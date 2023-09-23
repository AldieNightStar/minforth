package minforth

func getAt[T any](arr []T, id int, defval T) T {
	if id < 0 || id >= len(arr) {
		return defval
	}
	return arr[id]
}

func takeNumber(arr []string, id int) (int, bool) {
	s := getAt(arr, id, "")
	if s == "" {
		return 0, false
	}
	n, ok := lexNumber(s)
	if !ok {
		return 0, false
	}
	return int(n), true
}

func getAtMap[T any](m map[string]T, k string, defval T) T {
	v, ok := m[k]
	if !ok {
		return defval
	}
	return v
}
