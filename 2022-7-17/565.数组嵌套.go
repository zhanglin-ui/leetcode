package main

import "fmt"


func arrayNesting(nums []int) int {
	ret := 0
	exist := make([]bool, len(nums))

	for i := range nums {
		cnt := 0
		for {
			if exist[i] == true {
				break
			}
			exist[i] = true
			i = nums[i]
			cnt++
		}

		if cnt > ret {
			ret = cnt
		}
	}

	return ret
}

func main() {
	A := []int{5, 4, 0, 3, 1, 6, 2}

	ret := arrayNesting(A)

	fmt.Printf("ret is %d\n", ret)
}