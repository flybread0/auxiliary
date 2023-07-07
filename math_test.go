package auxiliary

import (
	"fmt"
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	fmt.Println("test value")
	fmt.Println(reverse(17))
	fmt.Println(reverse(-117))
}

func Test_isPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(101))
	fmt.Println(isPalindrome(-101))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(11))
}

// leetcode 9
func isPalindrome(x int) bool {
	if x < 10 && x >= 0 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}

	tmp := 0
	xx := x
	for x > 0 {
		tmp = 10*tmp + x%10
		x = x / 10
	}
	return tmp == xx
}

// leetcode7
func reverse(i int) int {
	min := math.MinInt32
	max := math.MaxInt32

	if i < min || i > max {
		return 0
	}

	flag := false
	if i < 0 {
		i = -i
		flag = true
	}

	r := 0
	for i > 0 {
		if (max-i%10)/10 < r {
			return 0
		}
		r = 10*r + i%10
		i = i / 10
	}

	if flag {
		return -r
	} else {
		return r
	}
}
