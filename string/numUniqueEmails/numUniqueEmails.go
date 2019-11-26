package numUniqueEmails

import "strings"

func numUniqueEmails(emails []string) int {
	res := make(map[string]struct{})
	for _, email := range emails {
		domain := strings.Split(email, "@")[1]
		var local strings.Builder
		for _, c := range email {
			switch rune(c) {
			case '.':
				continue
			case '+', '@':
				goto done
			default:
				local.WriteRune(c)
			}
		}
	done:
		res[local.String()+"@"+domain] = struct{}{}
	}
	return len(res)
}
