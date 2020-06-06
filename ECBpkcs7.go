package process

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func aesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return encrypted
}

// ecb加密方法
type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbDecryptEr ecb

func newECBDecryptEr(b cipher.Block) cipher.BlockMode {
	return (*ecbDecryptEr)(newECB(b))
}

func (x *ecbDecryptEr) BlockSize() int { return x.blockSize }

func (x *ecbDecryptEr) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func pKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// @title    		StringAesECBPKCS7Encrypted
// @description   	Aes ECB PKCS7 Encrypted
// @auth      		CyBen             		(2020-06-06 11:33:29)
// @param     		keyMSG, origDataMSG     string
// @return    		encryptedMSG       	    string
func StringAesECBPKCS7Encrypted(keyMSG, origDataMSG string) string {
	origData := []byte(origDataMSG) // 待加密的数据
	key := []byte(keyMSG)           // 加密的密钥
	encrypted := aesEncryptECB(origData, key)
	encryptedMSG := base64.StdEncoding.EncodeToString(encrypted)
	return encryptedMSG
}

// @title    		ByteAesECBDecrypt
// @description   	[]Byte Aes ECB PKCS7 Decrypt
// @auth      		CyBen             		(2020-06-06 11:37:10)
// @param     		dataMSG,key       		[]byte
// @return    		retData,error       	[]byte,error
func ByteAesECBDecrypt(dataMSG, key []byte) ([]byte, error) {
	data, _ := base64.StdEncoding.DecodeString(string(dataMSG))
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	ecb := newECBDecryptEr(block)
	retData := make([]byte, len(data))
	ecb.CryptBlocks(retData, data)
	retData = pKCS7UnPadding(retData)
	return retData, nil
}
