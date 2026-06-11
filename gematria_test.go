package gematria

import "testing"

func TestGematria(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		num  int
		want string
	}{
		{
			"Units",
			1,
			"א׳",
		},
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
