package main

import "fmt"

/*
逆序对
分治方法，归并排序思路
input:
5
3 1 4 5 2
output: 4
*/
func main() {
	var n int
	fmt.Scan(&n)
	iData := make([]int, n)
	iBuffer := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&iData[i])
	}
	fmt.Println(iData)
	fmt.Println(reverseOrderPairs(iData, iBuffer, 0, n-1))
	fmt.Println(iData)
}
func reverseOrderPairs(iData, iBuffer []int, iLow, iHigh int) int {
	if iLow == iHigh { return 0 }//只有一个元素
	iMid := (iLow+iHigh)/2
	c1 := reverseOrderPairs(iData, iBuffer, iLow, iMid)//分左
	c2 := reverseOrderPairs(iData, iBuffer, iMid+1, iHigh)//分右
	c3 := reverseOrderMerge(iData, iBuffer, iLow, iMid, iHigh)//合
	for i:=iLow; i<=iHigh; i++ {//复制升序列
		iData[i] = iBuffer[i]
	}
	return c1+c2+c3
}

func reverseOrderMerge(iData, iBuffer []int, iLow, iMid, iHigh int) int {
	i, j, k := iLow, iMid+1, iLow
	count := 0
	for i<=iMid && j<=iHigh {
		if iData[i]>iData[j] {
			count += iMid-i+1//对于当前j, i～iMid所有的元素都是逆序对
			iBuffer[k] = iData[j]
			k++
			j++
		} else {
			iBuffer[k] = iData[i]
			k++
			i++
		}
	}
	/**
	复制剩余部分
	 */
	for i<=iMid {
		iBuffer[k] = iData[i]
		k++
		i++
	}
	for j<=iHigh {
		iBuffer[k] = iData[j]
		k++
		j++
	}
	return count
}
