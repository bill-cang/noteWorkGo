package encode

import (
	"crypto/cipher"
	"crypto/des"
)

//使用des进行对称加密
func EncryptDES(src []byte, private []byte) []byte {
	//1、创建并返回一个使用DES算法的cipher.Block接口
	block, _ := des.NewCipher(private)
	//2、对最后一个明文分组进行数据填充
	src = padding(src, block.BlockSize())
	//3、创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//4、加密连续的数据块
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

//使用des进行解密
func DecryptDES(src, private []byte) []byte {
	block, err := des.NewCipher(private)
	if err != nil {
		print(err)
	}
	//2、创建一个密码分组位链接模式的，底层使用DES解密的BlockMode接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//3、数据库解密
	blockMode.CryptBlocks(src, src)
	//4、去掉最后一组的填充数据
	newText := unPadding(src)
	return newText
}

//使用des进行对称加密
func Encrypt3DES(src []byte, private []byte) []byte {
	//1、创建并返回一个使用DES算法的cipher.Block接口
	block, _ := des.NewTripleDESCipher(private)
	//2、对最后一个明文分组进行数据填充
	src = padding(src, block.BlockSize())
	//3、创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, private[:block.BlockSize()])
	//4、加密连续的数据块
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

//使用des进行解密
func Decrypt3DES(src, private []byte) []byte {
	block, err := des.NewTripleDESCipher(private)
	if err != nil {
		print(err)
	}
	//2、创建一个密码分组位链接模式的，底层使用DES解密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, private[:block.BlockSize()])
	//3、数据库解密
	blockMode.CryptBlocks(src, src)
	//4、去掉最后一组的填充数据
	newText := unPadding(src)
	return newText
}