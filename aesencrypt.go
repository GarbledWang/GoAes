package main

import (
	"crypto/aes"
	"crypto/cipher"
)

type AesEncrypt struct{

}

//获取key
func(AesEncrypt)GetKey(key string)[]byte{
	keyLen := len(key)
	if keyLen < 16 {
		for i := 0; i < 16-keyLen; i++ {
			key+="0"
		}
	}
	arrKey := []byte(key)
	if keyLen >= 32{
		return arrKey[:32]
	}
	if keyLen >= 24 {
		return arrKey[:24]
	}

	return arrKey[:16]
}

//加密
func(AesEncrypt)Encrypt(strMsg string,key []byte)([]byte,error){
	iv := []byte(key)[:aes.BlockSize]
	encrypted := make([]byte,len(strMsg))
	aesBlockEncrypter,err := aes.NewCipher(key)
	if err != nil {
		return nil ,err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter,iv)
	aesEncrypter.XORKeyStream(encrypted,[]byte(strMsg))
	return encrypted,nil
}

//解密
func(AesEncrypt)Decrypt(strMsg []byte,key []byte)(strDesc string,err error){
	defer func() {
		//错误处理
		if e := recover(); e!= nil{
			err = e.(error)
		}
	}()

	iv := []byte(key)[:aes.BlockSize]
	decrypted := make([]byte,len(strMsg))
	aseBlockDecrypter ,err := aes.NewCipher([]byte(key))
	if err != nil {
		return "",err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aseBlockDecrypter,iv)
	aesDecrypter.XORKeyStream(decrypted,strMsg)
	return string(decrypted),nil
}

