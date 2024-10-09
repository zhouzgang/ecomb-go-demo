package main


func main() {
	// 对一个数组，根据指定数组升序操作
	var keys  = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var excludeIndexs  = []int{3, 7, 8}
	// 升序结果：{1, 2, 4, 5, 6, 9, 10, 11, 12, 13, 14}

	index := len(excludeIndexs)
	for i := len(keys) - 1; i >= 0; i-- {

		if keys[i] > excludeIndexs[index - 1] {
			keys[i] = keys[i] + index
		}

		if keys[i] == excludeIndexs[index - 1] {
			keys[i] = keys[i] + index
			index--
		}

		if index == 0{
			break
		}
	}

	println("%v", keys)
	println("%v", excludeIndexs)
}