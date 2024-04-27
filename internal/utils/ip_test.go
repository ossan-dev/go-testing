package utils_test

import (
	"fmt"
	"testing"

	"github.com/ossan-dev/gotesting/internal/utils"
)

func ExampleIsValidIpOrCidr() {
	fmt.Println(utils.IsValidIpOrCidr("192.166.111.20/2"))
	// Output: true
}

func ExampleIsValidIpOrCidr_error() {
	fmt.Println(utils.IsValidIpOrCidr("192.166.111"))
	// Output: false
}

func BenchmarkIsValidIpOrCidr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.IsValidIpOrCidr("192.168.15.15/32")
	}
}
