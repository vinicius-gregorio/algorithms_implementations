package left_rotation

func rotateLeft(d int32, arr []int32) []int32 {
	res := arr
	for i := 0; i < int(d); i++ {
		res = append(res[1:], []int32{res[0]}...)

	}
	return res
}
