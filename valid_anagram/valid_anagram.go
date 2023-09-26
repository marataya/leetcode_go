package valid_anagram

import "fmt"

//Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//Example 1:
//Input: s = "anagram", t = "nagaram"
//Output: true
//Example 2:
//Input: s = "rat", t = "car"
//Output: false
//
//Constraints:
//1 <= s.length, t.length <= 5 * 104
//s and t consist of lowercase English letters.
//
//Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := map[rune]int{}
	for _, c := range s {
		sMap[c]++
	}
	for _, c := range t {
		if _, ok := sMap[c]; !ok {
			return false
		}
		sMap[c]--
		if sMap[c] == 0 {
			delete(sMap, c)
		}
	}
	return len(sMap) == 0
}

func Run() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
