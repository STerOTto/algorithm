package main

import "fmt"
/**
找到多于一半的数字
 */
func main() {
	var n int
	fmt.Scan(&n);
	ids := make([]int, n)
	for i:=0; i<n;i++  {
		fmt.Scan(&ids[i])
	}
	candidate :=0
	nTimes := 0
	for i:=0; i<n;i++ {
		if nTimes == 0 {
			candidate = ids[i]
			nTimes ++
		} else {
			if candidate == ids[i] {
				nTimes ++
			} else {
				nTimes --
			}
		}
	}

	fmt.Print(candidate)
}


