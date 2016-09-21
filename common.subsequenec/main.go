package main

import (
	"fmt"
)

/*
最长公共子序列
input：
ABCBDAB
CBCEDB
output： 4
最长公共子序列为BCDB
*/
const MAX int = 100
func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	fmt.Println("without recursion:", commonSubSequenceWithoutRecursion(str1, str2))
	fmt.Println("with recursion:", commonSubSequenceWithRecursion(str1, str2, 0, 0, len(str1), len(str2)))

}
/*
非递归写法
*/
func commonSubSequenceWithoutRecursion(str1, str2 string) int {
	C, B := [MAX][MAX]int{}, [MAX][MAX]int{}
	m, n := len(str1)+1, len(str2)+1
	//边界初始化
	for i:=0; i<m; i++ {
		C[i][0] = 0
	}
	for i:=0; i<n; i++ {
		C[0][i] = 0
	}
	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			if str1[i-1]==str2[j-1] {
				C[i][j] = C[i-1][j-1] + 1
				B[i][j] = 1
			} else if C[i-1][j] >= C[i][j-1] {
				C[i][j] = C[i-1][j]
				B[i][j] = 2
			} else {
				C[i][j] = C[i][j-1]
				B[i][j] = 3
			}

		}

	}
	return C[m-1][n-1]
}
/*
递归写法
*/
func commonSubSequenceWithRecursion(str1, str2 string, i, j, endI, endJ int) int {
	if i>=endI || j>=endJ {
		return 0
	}
	if str1[i]==str2[j] {
		return commonSubSequenceWithRecursion(str1, str2, i+1, j+1, endI, endJ) + 1
	}
	res1 := commonSubSequenceWithRecursion(str1, str2, i+1, j, endI, endJ)
	res2 := commonSubSequenceWithRecursion(str1, str2, i, j+1, endI, endJ)
	if res1>res2 {
		return res1
	} else {
		return res2
	}
}
