package GongYiFu

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/qnsoft/utils/ErrorHelper"
	"github.com/qnsoft/utils/StringHelper"

	"github.com/wenzhenxi/gorsa"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

/*
私钥加密
*/
func Rsa_Encode(data, _privateKey []byte) ([]byte, error) {
	_srt, errrt := gorsa.PriKeyEncrypt(string(data), string(_privateKey))
	return []byte(_srt), errrt
}

/*
公钥解密
*/
func Rsa_Decode(encrypted, _publicKey []byte) ([]byte, error) {
	_srt, errrt := gorsa.PublicDecrypt(string(encrypted), string(_publicKey))
	return []byte(_srt), errrt
}

/*
针对返回的各项参数解密
*/
func Jiemi_data(_json interface{}) interface{} {
	_obj := _json.(map[string]interface{})
	if _obj["data"] != nil && _obj["encryptkey"] != nil {
		_data := fmt.Sprintf("%s", _obj["data"])
		_jm_encryptKey := fmt.Sprintf("%s", _obj["encryptkey"])
		//fmt.Println(_jm_encryptKey)
		_aes_key, _ := Rsa_Decode([]byte(_jm_encryptKey), PublicKey) //公钥解密
		_aes_key_str := string(_aes_key)
		//fmt.Println(_aes_key_str)
		_data_a, _ := hex.DecodeString(strings.ToLower(_data))                       //先小写 再解码
		_b_json := StringHelper.AesDecryptECB([]byte(_data_a), []byte(_aes_key_str)) //AES解密
		//fmt.Println(string(_b_json))
		//_b_ret := new(JModel)
		//	_b_ret := new(interface{})
		var _b_ret interface{}
		err := json.Unmarshal(_b_json, &_b_ret)
		ErrorHelper.CheckErr(err)
		return _b_ret
	}
	return nil
}

/* 以下为编码转换 */

func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}
func Utf8ToGbk(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}
func StrToUtf8(str *string) error {
	b, err := GbkToUtf8([]byte(*str))
	if err != nil {
		return err
	}
	*str = string(b)
	return nil
}

func StrToGBK(str *string) error {
	b, err := Utf8ToGbk([]byte(*str))
	if err != nil {
		return err
	}
	*str = string(b)
	return nil
}
