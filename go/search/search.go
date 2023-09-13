package search

import (
	"fmt"

	"github.com/vinicius-gregorio/algorithms_implementations/randomize"
)

func Search() {
	fmt.Println("Starting Search")
	//populate toSearch with 100 items
	arr := make([]string, 100)
	for i := 0; i < 100; i++ {
		arr[i] = fmt.Sprintf("element%d", i)
	}
	searchElement := arr[randomize.GetRandomNumber(0, len(arr))]
	LinearSearch(arr, searchElement)
}
