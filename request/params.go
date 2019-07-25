package request

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	. "music163-hotcomments-spider/util"
	"strings"
)

func aesEncrypt(_text, _key, _iv string) string {
	pad := 16 - len(_text)%16
	_text += strings.Repeat(string(pad), pad)
	text := []byte(_text)
	key := []byte(_key)
	iv := []byte(_iv)
	blockCipher, _ := aes.NewCipher(key)
	cbcEncrypter := cipher.NewCBCEncrypter(blockCipher, iv)
	cryted := make([]byte, len(text))
	cbcEncrypter.CryptBlocks(cryted, text)
	return base64.StdEncoding.EncodeToString(cryted)
}

func rsaEncrypt(randomText, key, fp string) string {
	encSecKey := "257348aecb5e556c066de214e531faadd1c55d814f9be95fd06d6bff9f4c7a41f831f6394d5a3fd2e3881736d94a02ca919d952872e7d0a50ebfa1769a7a62d512f5f1ca21aec60bc3819a9c3ffca5eca9a0dba6d6f7249b06f5965ecfff3695b54e1c28f3f624750ed39e7de08fc8493242e26dbc4484a01c76f739e135637c"
	return encSecKey
}

func generatorData(page int) string {

	firstParam := ToJSON(map[string]string{
		"rid":        "",
		"offset":     ToString(page * 20),
		"total":      "True",
		"limit":      "20",
		"csrf_token": "",
	})
	secondParam := "010001"
	thirdParam := "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	forthParam := "0CoJUm6Qyw8W8jud"
	iv := "0102030405060708"
	firstKey := forthParam
	secondKey := "FFFFFFFFFFFFFFFF"
	encText := aesEncrypt(firstParam, firstKey, iv)

	encText = aesEncrypt(encText, secondKey, iv)

	encSecKey := rsaEncrypt(secondKey, secondParam, thirdParam)

	return fmt.Sprintf("params=%s&encSecKey=%s",encText,encSecKey)
}
