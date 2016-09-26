package main

import (
	"fmt"
	"math/rand"
)
/*
找出最大的K个数，快排的方法
*/
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&a[i])
	}
	midK := partitions(a, 0, n-1)
	for n-midK > k {
		fmt.Println("midK:", midK)
		midK = partitions(a, midK+1, n-1)
	}
	fmt.Print(a)
	fmt.Print(a[midK:])

}

func partitions(arr []int, start, end int) (midK int) {
	midK = rand.Intn(end-start) + start
	arr[start], arr[midK] = arr[midK], arr[start]
	key := arr[midK]
	for i, j:= start+1, end; i<j;  {
		if arr[i] < key {
			i++
		}
		if arr[j] > key {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
		midK = i
	}
	arr[start], arr[midK] = arr[midK], arr[start]
	return
}
