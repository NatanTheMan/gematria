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

func TestGematria_Dozens(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		num  int
		want string
	}{
		{"Yud", 10, "י׳"},
		{"Caf", 20, "כ׳"},
		{"Lamed", 30, "ל׳"},
		{"Mem", 40, "מ׳"},
		{"Nun", 50, "נ׳"},
		{"Samech", 60, "ס׳"},
		{"Ayin", 70, "ע׳"},
		{"Pei", 80, "פ׳"},
		{"Tsadi", 90, "צ׳"},
		{"dozens with units", 11, "י״א"},
		{"dozens with units", 39, "ל״ט"},
		{"dozens with units", 82, "פ״ב"},
		{"dozens with units", 46, "מ״ו"},
		{"dozens with units", 95, "צ״ה"},
		{"teenfive", 15, "ט״ו"},
		{"teensix", 16, "ט״ז"},
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

func TestGematria_Hundreds(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		num  int
		want string
	}{
		{"Kof", 100, "ק׳"},
		{"Resh", 200, "ר׳"},
		{"Shin", 300, "ש׳"},
		{"Tav", 400, "ת׳"},
		{"Kaf sofit", 500, "ך׳"},
		{"Mem sofit", 600, "ם׳"},
		{"Nun sofit", 700, "ן׳"},
		{"Pe sofit", 800, "ף׳"},
		{"Tsadi sofit", 900, "ץ׳"},
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
