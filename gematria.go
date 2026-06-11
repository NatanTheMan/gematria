package gematria

import (
	"fmt"
	"math"
)

var (
	units    = []string{"", "א", "ב", "ג", "ד", "ה", "ו", "ז", "ח", "ט"}
	dozens   = []string{"", "י", "כ", "ל", "מ", "נ", "ס", "ע", "פ", "צ"}
	hundreds = []string{"", "ק", "ר", "ש", "ת", "ך", "ם", "ן", "ף", "ץ"}
)

func Gematria(num int) string {
	if num < 10 {
		return fmt.Sprintf("%s׳", units[num])
	}
	if num < 100 && num%10 == 0 {
		return fmt.Sprintf("%s׳", dozens[num/10])
	}
	if num%100 == 0 {
		return fmt.Sprintf("%s׳", hundreds[num/100])
	}
	if num == 15 {
		return "ט״ו"
	}
	if num == 16 {
		return "ט״ז"
	}
	if num < 100 {
		return fmt.Sprintf("%s״%s", dozens[num/10], units[num%10])
	}

	return fmt.Sprintf("%s״%s", hundreds[int(math.Floor(float64(num/100)))], dozens[(num%100)/10])
}
