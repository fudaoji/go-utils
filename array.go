package utils

//函数Contains用于判断数组$array中是否包含元素$search
func ContainsStr(array []string, search string) (exists bool, index int) {
	exists = false
	index = -1

	for i, v := range array {
		if search == v {
			index = i
			exists = true
			return exists, index
		}
	}
	return exists, index
}

//函数Contains用于判断数组$array中是否包含元素$search
func ContainsInt(array []int, search int) (exists bool, index int) {
	exists = false
	index = -1

	for i, v := range array {
		if search == v {
			index = i
			exists = true
			return exists, index
		}
	}
	return exists, index
}
