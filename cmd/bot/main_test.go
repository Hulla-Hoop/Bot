package main_test

import (
	"strconv"
	"testing"

	"github.com/shamil/bot/internal/service/product"
)

func BenchmarkXxx(b *testing.B) {
	producns := product.NewService()
	for i := 0; i < b.N; i++ {
		producns.New(strconv.Itoa(i))
	}
}
