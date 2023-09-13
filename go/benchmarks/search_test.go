package benchmarks

import (
	"fmt"
	"testing"

	"github.com/vinicius-gregorio/algorithms_implementations/randomize"
	"github.com/vinicius-gregorio/algorithms_implementations/search"
)

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search.Search()
	}
}
func BenchmarkLinearSearch(b *testing.B) {
	qtdy := 1000
	arr := make([]string, qtdy)
	for i := 0; i < qtdy; i++ {
		arr[i] = fmt.Sprintf("element%d", i)
	}

	b.ResetTimer() // Reset the timer to ignore the setup time

	for i := 0; i < b.N; i++ {
		searchElement := arr[randomize.GetRandomNumber(0, len(arr)-1)]
		search.LinearSearch(arr, searchElement)
	}
}
