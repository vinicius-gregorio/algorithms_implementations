package search

import "fmt"

func LinearSearch(arr []string, element string) {
	//Using linear search to find the element
	for i := 0; i < len(arr); i++ {
		if arr[i] == element {
			fmt.Println("Element found at index ", i)
			break
		}
	}
}
