package main

import "fmt"
/*
十进制数中1的个数
*/
func main() {
	var (
		number int64
		count, iFactor	int64 = 0, 1
	)
	fmt.Scan(&number)
	for number/iFactor != 0 {
		iHigh := number/(iFactor*10)
		iLow := number - (number/iFactor)*iFactor
		iCurrent := (number/iFactor) % 10
		switch iCurrent {
		case 0:
			count += iHigh * iFactor
		case 1:
			count += iHigh*iFactor + iLow + 1
		default:
			count += (iHigh+1)*iFactor
		}
		iFactor *= 10
	}
	fmt.Print(count)
}
