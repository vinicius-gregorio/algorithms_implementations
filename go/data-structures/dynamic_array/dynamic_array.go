package main

func main() {

}

func dynamicArray(n int32, queries [][]int32) []int32 {
	// Write your code here
	answers := []int32{}
	arr := make([][]int32, n)
	lastAnswer := int32(0)

	for i := int32(0); i < n; i++ {
		arr[i] = []int32{}
	}

	for i := 0; i < len(queries); i++ {
		inpt := queries[i]
		q := inpt[0]
		x := inpt[1]
		y := inpt[2]

		idx := ((x ^ lastAnswer) % n)
		if q == 1 {
			arr[idx] = append(arr[idx], y)
		}
		if q == 2 {
			subIndex := (y % int32(len(arr[idx])))
			lastAnswer = arr[idx][subIndex]
			answers = append(answers, lastAnswer)
		}
	}

	return answers
}
