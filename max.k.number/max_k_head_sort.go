package main

import (
	"fmt"
)
/*
求最大的K个数，用小顶堆的方式，
每次新来的数据，都与当前根节点比较，如果比根节点大，则代替根节点，重新调整堆
*/
func main() {
	var(
		k int
		all, head []int
	)
	fmt.Scan(&k)
	all = make([]int, 0)
	head = make([]int, 0)
	for i:=0; ;i++  {
		var key int
		fmt.Scan(&key)
		all = append(all, key)
		if i<k-1 {//0~k-1个 直接append
			head = append(head, all[i])
		} else if i==k-1 {//第k个开始建堆
			head = append(head, all[i])
			buildHead(head, k)
		} else {//大于K个时调整堆
			if all[i] > head[0] {
				head[0] = all[i]
				headAdjust(head, 0, k)

			}
		}
		fmt.Println("all:", all, "\nhead:", head)
	}

}
func headAdjust(a []int, cur, size int)  {
	lChild := cur*2
	rChild := cur*2 +1
	min := cur
	if cur<=size/2 {//非叶子节点，进行调整
		if lChild < size && a[lChild] < a[min] {
			min = lChild
		}
		if rChild < size && a[rChild] < a[min] {
			min = rChild
		}
		if min!=cur {
			a[cur], a[min] = a[min], a[cur]
			headAdjust(a, min, size)//交换位置后调整以min为根节点的子树是否平衡
		}
	}
}
func buildHead(a []int, size int)  {
	for i:=size/2; i>=0; i-- {
		headAdjust(a, i, size)
	}
}
