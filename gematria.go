package gematria

import "fmt"

func Gematria(num int) string {
	units := []string{
		"",
		"א",
		"ב",
		"ג",
		"ד",
		"ה",
		"ו",
		"ז",
		"ח",
		"ט",
	}
	return fmt.Sprintf("%s׳", units[num])
}
