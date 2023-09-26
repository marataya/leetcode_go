package happy_number

import "fmt"

/*
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.
*/

// if number of iterations i == 10 then beats 100%
func isHappy(n int) bool {
	for i := 0; i < 10000; i++ {
		if n == 1 {
			return true
		}
		m := n
		sumSquared := 0
		for m > 0 {
			sumSquared += (m % 10) * (m % 10)
			m /= 10
		}
		n = sumSquared
	}
	return false
}

func Run() {
	fmt.Println(isHappy(19))
}
