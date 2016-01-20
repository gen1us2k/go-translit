package translit

import (
	"regexp"
)


func Translit(s string) string {
	cyrillic_string := s
	for k, v := range(en_translations) {
		r, _ := regexp.Compile(k)
		cyrillic_string = r.ReplaceAllString(cyrillic_string, v)
	}
	return cyrillic_string
}
