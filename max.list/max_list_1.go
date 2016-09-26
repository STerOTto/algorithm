package main

import "fmt"

/*
最长子序列，记录长度为l[i]的子序列的最大值的最小值maxV[l[i]]
*/
func main() {
	var n int
	fmt.Scan(&n)
	l := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&l[i])
	}
	fmt.Println(lis(l))
}

func lis(l[]int) (max int) {
	if len(l) == 0 {
		return
	}
	maxV := make([]int, len(l) + 1)
	maxV[1] = l[0]
	maxV[0] = min(l) - 1 //边界值
	lis := make([]int ,len(l))
	for i:=0; i<len(lis); i++ {
		lis[i] = 1
	}
	maxL := 1 //maxV的最大值的最小值，默认是1
	for i:=0; i<len(l); i++ {
		j:=maxL
		for ; j>=0; j-- {
			if l[i]>maxV[j] {
				lis[i] = j+1
				break	//j为从大到小，找到最大的j需break
			}
		}
		if lis[i] > maxL {
			maxL = lis[i]
			maxV[lis[i]] = l[i]
		} else if maxV[j] < l[i] && l[i] < maxV[j+1] {
			maxV[j+1] = l[i]
		}
	}
	max = maxL
	return
}
func min(l[]int) (min int) {
	min = l[0]
	for i := range l {
		if min > i {
			min = i
		}
	}
	return
}
