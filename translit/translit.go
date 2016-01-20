package go_translit
import (
	"fmt"
	"regexp"
)


func Translit(s string) {
	cyr_string := s
	for k, v := range(en_translations) {
		r, _ := regexp.Compile(k)
		cyr_string = r.ReplaceAllString(cyr_string, v)
	}
	fmt.Println(cyr_string)
//	return str
}
