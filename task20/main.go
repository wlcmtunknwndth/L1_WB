package main

import (
	"fmt"
	"strings"
	"time"
)

func reverseWordsOrder(str string) string {
	if len(str) <= 1 {
		return str
	}

	var ans strings.Builder
	var end int

	if str[len(str)-1] != 32 {
		end = len(str)
	} else {
		// look for last word in the string
		for i := len(str) - 1; i >= 0; i-- {
			if str[i] != 32 {
				end = i + 1
				break
			}
		}
	}

	for i := len(str) - 1; i >= 0; i-- {
		if i == 0 {
			if str[i] != 32 {
				ans.WriteString(str[i:end])
				ans.WriteByte(32)
				//fmt.Println(str[i:end], "|")
			}
			continue
		}
		if str[i-1] == 32 {
			ans.WriteString(str[i:end])
			ans.WriteByte(32)
			//fmt.Println(str[i:end], "|")
			end = i - 1
		}
	}
	return ans.String()
}

func main() {
	start := time.Now()
	var str = " snow a dog sun a"

	fmt.Println(reverseWordsOrder(str))

	fmt.Println(time.Since(start))
}
