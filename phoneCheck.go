package process

import "regexp"

//	@CheckRealPhone
// 	@description  	check phoneNumber 	true/false
// 	@Author 		CyBen
// 	@param   		phoneNumber 		string
//	@return   		true/false  		bool
func CheckRealPhone(phonenumber string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	valid := rgx.MatchString(phonenumber)
	if valid {
		return true
	} else {
		return false
	}

}
