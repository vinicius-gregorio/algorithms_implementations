package search

func LinearSearch(arr []string, element string) {

	for i := 0; i < len(arr); i++ {
		if arr[i] == element {

			break
		}
	}
}
