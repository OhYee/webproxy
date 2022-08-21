package utils

import (
	"fmt"
)

func FormatBytes(bytes int) string {
	b := float64(bytes)

	unitArr := []string{"Bytes", "KB", "MB", "GB", "TB"}
	idx := 0
	for b >= 1024 && idx < len(unitArr)-1 {
		b /= 1024
		idx++
	}

	return fmt.Sprintf("%.2f%s", b, unitArr[idx])
}
