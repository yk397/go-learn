package pointer

//题目一
func Add10(a *int) {
	*a += 10
}

//题目二
func ArrMulti2(arr *[]int) {
	for i := range *arr {
		(*arr)[i] *= 2
	}
}
