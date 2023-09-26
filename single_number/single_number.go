package single_number

import "fmt"

func singleNumber(nums []int) int {
	hashMap := map[int]int{}
	fmt.Printf("%v\n", nums)
	for _, num := range nums {
		if _, ok := hashMap[num]; ok {
			hashMap[num] += 1
		} else {
			hashMap[num] = 1
		}
	}
	fmt.Printf("%v\n", hashMap)
	for k, v := range hashMap {
		if v == 1 {
			return k
		}
	}
	return 69
}

// faster solution using XOR
func singleNumber1(nums []int) (res int) {
	res = nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return
}

func Run() {
	fmt.Println(singleNumber([]int{1, 2, 4, 2, 1}))
	fmt.Println(singleNumber1([]int{1, 2, 4, 2, 1}))
}
