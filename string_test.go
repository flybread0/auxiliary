package auxiliary

import (
	"fmt"
	"math"
	"testing"
)

func Test_s2i(t *testing.T) {
	fmt.Println(myAtoi("     -567"))
	fmt.Println(myAtoi("4193"))
	fmt.Println(myAtoi("4193 with words"))
}

func myAtoi(s string) int {
	//Read in and ignore any leading whitespace.
	index := 0
	for index < len(s) && string(s[index]) == " " {
		index++
	}

	//Check if the next character (if not already at the end of the string) is '-' or '+'.
	//Read this character in if it is either. This determines if the final result is negative or positive respectively.
	// Assume the result is positive if neither is present.

	flag := true
	r := 0
	if index< len(s) && (string(s[index]) == "+" || string(s[index]) == "-") && index != len(s)-1 {
		if string(s[index]) == "-" {
			flag = false
		}
		index++
	} else if index< len(s) && (string(s[index]) == "+" || string(s[index]) == "-") && index == len(s)-1 {
		return 0
	}

	for {
		if index < len(s) && s[index] >= "0"[0] && s[index] <= "9"[0] {
			if (math.MaxInt32-int((s[index]-"0"[0])))/10 < r {
				if flag {
					return math.MaxInt32
				}else{
					return math.MinInt32
				}
			}
			r = r*10 + int((s[index] - "0"[0]))
			index++
		} else {
			break
		}
	}

	if flag {
		return r
	} else {
		return -r
	}
}
