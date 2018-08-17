package translit

import "testing"

func TestToCyrillic(t *testing.T) {
	tests := map[string]string{
		"Prived": "Привед",
		//		"Jackie Chan": "Йацкие Цхан",
		//		"Jessie Jane": "Йессие Йане",
		"Impossibru": "Импоссибру",
		"Hello":      "Хелло",
		"Red Center": "Ред Центер",
		"Andrew":     "Андрей",
		"Anatoliy":   "Анатолий",
		"Vasilii":    "Василий",
	}
	for k, v := range tests {
		tr := Translit(k)
		if tr != v {
			t.Errorf("\"%s\" expected %q got %q", k, v, tr)
		}
	}
}
