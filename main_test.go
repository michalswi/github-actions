package main

import (
	"testing"
)

func TestGetName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"John", "Hi John"},
		{"", "noname"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getName(tt.name); got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVersion(t *testing.T) {
	want := "1.0.0"
	if got := getVersion(); got != want {
		t.Errorf("getVersion() = %v, want %v", got, want)
	}
}
