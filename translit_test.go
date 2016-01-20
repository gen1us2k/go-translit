package go_translit

import "testing"

func TestToLatin(t *testing.T) {
	tests := map[string]string{
		"Prived": "Привед",
//		"Jackie Chan": "Йацкие Цхан",
//		"Jessie Jane": "Йессие Йане",
		"Impossibru": "Импоссибру",
		"Hello": "Хелло",
		"Red Center": "Ред Центер",
	}
	for k, v := range tests {
		tr := Translit(k)
		if tr != v {
			t.Errorf("\"%s\" expected %q got %q", k, v, tr)
		}
	}
}