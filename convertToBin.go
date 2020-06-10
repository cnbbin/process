package process

import (
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
		if lsb == 1 {
			s1++
		} else {
			s0++
		}
		// 三种 int  转 string
		//s = fmt.Sprintf("%d", lsb) +s
		s = strconv.FormatInt(int64(lsb), 2) + s
		//s = strconv.Itoa(lsb) + s

		//参考各大神，strconv.FormatInt()效率最高，fmt.Sprintf()效率最低 ,而 strconv.Itoa 是通过 strconv.FormatInt()转转

	}
	return s, s0, s1
}
