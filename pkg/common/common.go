package common

func StrInList(s string, l []string) bool {
	for _, e := range l {
		if s == e {
			return true
		}
	}
	return false
}
