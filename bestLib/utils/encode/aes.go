package encode

import (
	"crypto/aes"
	"crypto/cipher"
)

//使用AES加密
func EncryptAES(src, private []byte) []byte {
	//1、创建并返回一个使用AES算法的cipher.Block接口
	block, err := aes.NewCipher(private)
	if err != nil {
		panic(err)
	}
	//2、数据填充
	src = padding(src, block.BlockSize())
	//3、创建一个密码分组位链接模式的，底层使用AES解密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, private)
	//4、数据加密
	blockMode.CryptBlocks(src, src)
	return src
}

//使用AES加密
func DecryptAES(src, private []byte) []byte {
	//1、创建并返回一个使用AES算法的cipher.Block接口
	block, err := aes.NewCipher(private)
	if err != nil {
		panic(err)
	}
	//2、创建一个密码分组位链接模式的，底层使用AES解密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, private)
	//3、数据解密
	blockMode.CryptBlocks(src, src)
	//4、去掉填充数据
	src = unPadding(src)
	return src
}

