package utils

import (
	"fmt"
	"testing"
)

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		bytes int
		want  string
	}{
		{5, "5.00Bytes"},
		{1024, "1.00KB"},
		{2048, "2.00KB"},
		{1024 * 1024, "1.00MB"},
		{2 * 1024 * 1024, "2.00MB"},
		{2 * 1024 * 1024 * 1024, "2.00GB"},
		{2 * 1024 * 1024 * 1024 * 1024, "2.00TB"},
		{2 * 1024 * 1024 * 1024 * 1024 * 1024, "2048.00TB"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%s", tt.bytes, tt.want), func(t *testing.T) {
			if got := FormatBytes(tt.bytes); got != tt.want {
				t.Errorf("FormatBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
