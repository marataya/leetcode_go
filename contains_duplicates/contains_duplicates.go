package contains_duplicates

import "fmt"

func containsDuplicate(nums []int) bool {
	s := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if _, ok := s[nums[i]]; ok {
			return true
		}
		s[nums[i]] = true
	}
	return false
}

func Run() {
	nums1 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums1))
	fmt.Println(containsDuplicate(nums2))
}
