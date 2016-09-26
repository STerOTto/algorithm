package main

import "fmt"

/*
最长递增子序列
*/
func main() {
	var n int
	fmt.Scan(&n)
	l := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&l[i])
	}
	fmt.Println(LIS(l))

}

func LIS(l[]int) (max int) {
	lis := make([]int, len(l))
	max = 0
	for i:=0; i<len(l); i++ {
		lis[i] = 1
		for k:=0; k<i; k++ {
			if l[i]>l[k] && lis[k]+1>lis[i] {
				lis[i] = lis[k] + 1
			}
		}
		if max < lis[i] {
			max = lis[i]
		}
	}
	return
}
