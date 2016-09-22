package main

import "fmt"

/**
最优二叉搜索树
input：
3
0.5 0.1 0.05
0.15 0.1 0.05 0.05
out:
1.5
 */
func main() {
	var n int
	fmt.Scan(&n)
	pProb := make([]float32, n+1)
	for i:=1; i<n; i++ {
		fmt.Scan(&pProb[i])
	}
	qProb := make([]float32, n+1)
	for i:=0; i<=n; i++ {
		fmt.Scan(&qProb[i])
	}
	fmt.Println(optimalBST(pProb, qProb, n))
}

const (
	MaxNum int = 10000
)
var(
	C, w [MaxNum][MaxNum]float32
	root [MaxNum][MaxNum]int
)
func optimalBST(pProb, qProb []float32, n int) float32 {
	//初始化,空树，只有虚节点
	for i:=0; i<=n; i++ {
		C[i+1][i] = 0
		w[i+1][i] = qProb[i]
	}
	for r:=1; r<=n; r++ { //实节点数
		for i:=1; i<=n-r+1; i++ {//实节点开始编号
			j := i+r-1 //实节点的结束编号
			root[i][j] = i//第一个root节点为i
			C[i][j] = C[i+1][j]
			w[i][j] = w[i][j-1] + pProb[j] + qProb[j]
			for k:=i+1; k<=j; k++ {
				tmp := C[i][k-1] + C[k+1][j]
				if tmp<C[i][j] {
					C[i][j] = tmp
					root[i][j] = k
				}
			}
			C[i][j] += w[i][j]
		}

	}
	return C[1][n]
}

