package main

import (
	"fmt"
	"math"
)

/*
计算字符串的相似度
*/
func main() {
	var strA, strB string
	fmt.Scan(&strA, &strB)
	fmt.Println(calculateStringDistance(strA, strB, 0, 0, len(strA)-1, len(strB)-1))
}

func calculateStringDistance(strA, strB string, startA, startB, endA, endB int) int {
	if startA > endA {
		if startB > endB {
			return 0
		} else {
			return endB - startB + 1
		}
	}
	if startB > endB {
		if startA > endA {
			return 0
		} else {
			return endA - startA + 1
		}
	}
	if strA[startA] == strB[startB] {
		return calculateStringDistance(strA, strB, startA+1, startB+1, endA, endB)
	} else {
		t1 := calculateStringDistance(strA, strB, startA, startB+1, endA, endB)
		t2 := calculateStringDistance(strA, strB, startA+1, startB, endA, endB)
		t3 := calculateStringDistance(strA, strB, startA+1, startB+1, endA, endB)
		return minValue(t1, t2, t3) + 1
	}
}
func minValue(args ...int) (min int) {
	min = math.MaxInt8
	for v := range args{
		if v < min {
			min = v
		}
	}
	return
}