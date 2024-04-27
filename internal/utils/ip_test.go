package utils_test

import (
	"testing"

	"github.com/ossan-dev/gotesting/internal/utils"
)

func BenchmarkIsValidIpOrCidr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.IsValidIpOrCidr("192.168.15.15/32")
	}
}
