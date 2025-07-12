package pkg

import "regexp"

var phoneRegex = regexp.MustCompile(`^1[3-9][0-9]{9}$`)

func Regex(phone string) bool {
	return phoneRegex.MatchString(phone)
}
