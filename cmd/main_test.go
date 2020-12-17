package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	main()
}

func BenchmarkMain(b *testing.B) {
	b.StartTimer()
	main()
	b.StopTimer()
	b.ReportAllocs()
}
