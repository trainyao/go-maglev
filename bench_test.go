package maglev

import (
	"fmt"
	"testing"
)

var total int

func BenchmarkGenerate(b *testing.B) {

	const size = 2000

	var names []string
	for i := 0; i < size; i++ {
		names = append(names, fmt.Sprintf("backend-%d", i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		//table := New(names, SmallM)
		offsets, _ := generatePermutations(names, BigM)
		total += len(offsets)
	}
}
