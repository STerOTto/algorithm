package main

import (
	"math"
	"fmt"
)

/**
矩阵连乘
Input:
3
10 100 5 50
Output:
7500
 */
func main() {
	var (
		matrixNum int
	)
	fmt.Scan(&matrixNum)
	for i:=0; i<=matrixNum; i++ {
		fmt.Scan(&dim[i])
	}
	fmt.Println(matrixChain(matrixNum))
}

const MAX_NUM int = 1000
var(
	dim [MAX_NUM]int
	memoTable [MAX_NUM][MAX_NUM]int
	bestK [MAX_NUM][MAX_NUM]int
)

func matrixChain(matrixNum int) int {
	for i:=1; i<=matrixNum; i++ {
		memoTable[i][i] = 0
	}
	for l:=2; l<=matrixNum; l++ {
		for i:=1; i<=matrixNum-l+1; i++ {
			j:=i+l-1
			memoTable[i][j] = math.MaxInt8
			for k:=i; k<j; k++ {
				ans := memoTable[i][k] + memoTable[k+1][j] + dim[i-1]*dim[k]*dim[j]
				if ans<memoTable[i][j] {
					bestK[i][j] = k
					memoTable[i][j] = ans
				}
			}
		}
	}
	return memoTable[1][matrixNum]
}