package process

import (
	"fmt"
	"strconv"
)

//	@convertToBin
// 	@description  	DecimalNumberToBin
// 	@Author 		CyBen
// 	@param   		DecimalNumber 		int
//	@return   		BinNumber   Bin1count Bin0count		string int int
func convertToBin(num int) (string, int, int) {
	s := ""
	s1 := 0
	s0 := 0
	if num == 0 {
		return "0", 0, 1
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		fmt.Println(lsb)
		if lsb == 1 {
			s1++
		} else {
			s0++
		}
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s, s0, s1
}
