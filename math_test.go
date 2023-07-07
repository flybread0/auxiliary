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
