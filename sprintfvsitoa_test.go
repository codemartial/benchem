package benchem_test

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkFmtSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", 6785624)
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(6785624)
	}
}
