package search

import (
	"fmt"

	"github.com/vinicius-gregorio/algorithms_implementations/randomize"
)

func Search() {
	qtdy := 1000
	arr := make([]string, qtdy)
	for i := 0; i < qtdy; i++ {
		arr[i] = fmt.Sprintf("element%d", i)
	}
	searchElement := arr[randomize.GetRandomNumber(0, len(arr)-1)]
	LinearSearch(arr, searchElement)
}
