package gematria

import "fmt"

var (
	units  = []string{"", "א", "ב", "ג", "ד", "ה", "ו", "ז", "ח", "ט"}
	dozens = []string{"", "י", "כ", "ל", "מ", "נ", "ס", "ע", "פ", "צ"}
)

func Gematria(num int) string {
	if num < 10 {
		return fmt.Sprintf("%s׳", units[num])
	}
	return fmt.Sprintf("%s׳", dozens[(num/10)])
}
