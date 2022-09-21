package rand

import "testing"

func BenchmarkRandStrMaskSB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStrMaskSB(6)
	}
}

func BenchmarkRandStrSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStrSimple(6)
	}
}
func BenchmarkRandStrSimpleSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStrSimpleSlice(6)
	}
}
