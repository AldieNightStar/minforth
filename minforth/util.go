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

func take2Numbers(arr []string, id1, id2 int) (int, int, bool) {
	a, aok := takeNumber(arr, id1)
	b, bok := takeNumber(arr, id2)
	if !aok || !bok {
		return 0, 0, false
	}
	return a, b, true
}
