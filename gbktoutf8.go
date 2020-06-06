package process

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// @title    		ByteUtf8ToGbk
// @description   	Utf8 Transform To Gbk
// @auth      		CyBen             		(2020-06-06 11:10:35)
// @param     		utf8Byte        		[]byte
// @return    		gbkByte,error       	[]byte,error
func ByteUtf8ToGbk(utf8byte []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(utf8byte), simplifiedchinese.GBK.NewEncoder())
	gbkByte, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return gbkByte, nil
}

// @title    		ByteGbkToUtf8
// @description   	Gbk Transform To Utf8
// @auth      		CyBen             		(2020-06-06 11:14:10)
// @param     		gbkByte        			[]byte
// @return    		utf8Byte,error       	[]byte,error
func ByteGbkToUtf8(gbkByte []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(gbkByte), simplifiedchinese.GBK.NewDecoder())
	utf8Byte, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return utf8Byte, nil
}
