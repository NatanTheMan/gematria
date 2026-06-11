package gematria

import "testing"

func TestGematria_Units(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		num  int
		want string
	}{
		{"Alef", 1, "א׳"},
		{"Beit", 2, "ב׳"},
		{"Guimel", 3, "ג׳"},
		{"Dalet", 4, "ד׳"},
		{"He", 5, "ה׳"},
		{"Vav", 6, "ו׳"},
		{"Zayn", 7, "ז׳"},
		{"Chet", 8, "ח׳"},
		{"Têt", 9, "ט׳"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Gematria(tt.num)

			if got != tt.want {
				t.Errorf("Gematria() = %v, want %v", got, tt.want)
			}
		})
	}
}
