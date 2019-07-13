package stacks

func match(a, b string) bool {
	c := a + b
	return (c == "()" || c == "[]" || c == "{}")
}


func isValid(s string) bool {
	str := ""
	for _, t := range s {
		v := string(t)
		if v == "(" || v == "[" || v == "{" {
			str += v
			continue
		}
		if len(str) == 0 {
			return false
		}
		if !match(string(str[len(str)-1]), v) {
			return false
		}
		str = str[:len(str)-1]
	}
	return len(str) == 0
}
