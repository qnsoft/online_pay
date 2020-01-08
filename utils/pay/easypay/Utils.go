package EasyPay

import (
	"fmt"
	"sort"
	"strings"
	"qnsoft/web_api/utils/ArryHelper"
	php2go "qnsoft/web_api/utils/Php2go"
)

/*
主机头集排序Sign签名 根据key排序
mp 参数集合
*/
func SignHeaders(mp map[string]string) (string, map[string]string) {
	rt_str := ""
	rtt := make(map[string]string)
	var newMp = make([]string, 0)
	for k, _ := range mp {
		newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	for _, v := range newMp {
		if v == "" {
			rt_str += ""
		} else {
			rt_str += strings.ToLower(v) + ";"
		}
		_value := mp[v]
		rtt[strings.ToLower(v)] = php2go.Trim(_value) //+ "\n"
	}
	return php2go.Rtrim(rt_str, ";"), rtt
}

/*
参数集排序Sign签名 根据key排序
mp 参数集合
*/
func GetCanonicalQueryString(mp map[string]interface{}) string {
	_str := ""
	var newMp = make([]string, 0)
	for k, v := range mp {
		v_a := fmt.Sprintf("%s", v)
		newMp = append(newMp, k+"="+php2go.URLEncode(v_a))
	}
	sort.Strings(newMp)
	for _, v := range newMp {
		_str += v + "&"
	}
	return php2go.Rtrim(_str, "&")
}

/*

 */
func GetHeadersToSign(headers map[string]string, headersToSign []string) map[string]string {
	_rt := make(map[string]string)
	if len(headersToSign) > 0 {
		_arry := make([]string, 0)
		for _, v := range headersToSign {
			_arry = append(_arry, strings.ToLower(v))
		}
		headersToSign = _arry
	}
	for _key, _value := range headers {
		if (_value == "" && isDefaultHeaderToSign(_key)) || (_value != "" && arry.Contains(headersToSign, strings.ToLower(_key)) && strings.ToLower(_key) != "authorization") {
			_rt[_key] = _value
		}
	}
	//开始排序
	newMp := make([]string, 0)
	for k, _ := range _rt {
		newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	_rt_a := make(map[string]string)
	for _, v := range newMp {
		_rt_a[v] = _rt[v]
	}
	return _rt_a
}
func isDefaultHeaderToSign(header string) bool {
	_rt := false
	_header_t := php2go.Trim(header, "")
	header = strings.ToLower(_header_t)
	defaultHeadersToSign := make([]string, 0)
	defaultHeadersToSign = append(defaultHeadersToSign, "host")
	defaultHeadersToSign = append(defaultHeadersToSign, "content-length")
	defaultHeadersToSign = append(defaultHeadersToSign, "content-type")
	defaultHeadersToSign = append(defaultHeadersToSign, "content-md5")
	if strings.Contains(header, "x-yop-") || arry.Contains(defaultHeadersToSign, header) {
		_rt = true
	}
	return _rt
}

/*

 */
func GetCanonicalHeaders(mp map[string]string) string {
	_str := ""
	var newMp = make([]string, 0)
	for k, _ := range mp {
		newMp = append(newMp, k+":"+php2go.URLEncode(mp[k]))
	}
	sort.Strings(newMp)
	for _, v := range newMp {
		if _str == "" {
			_str += ""
		} else {
			_str += "\n"
		}
		_str += v
	}
	return _str
}

/*
参数集排序Sign签名 根据key排序
mp 参数集合
*/
// func SignParameters(mp map[string]interface{}, secret string) string {
// 	_str := ""
// 	var newMp = make([]string, 0)
// 	for k, _ := range mp {
// 		newMp = append(newMp, k)
// 	}
// 	sort.Strings(newMp)
// 	for _, v := range newMp {
// 		_srt := ""
// 		switch _value := mp[v].(type) {
// 		case string:
// 			_srt = _value
// 		case int:
// 			_srt = strconv.Itoa(_value)
// 		}
// 		if len(_srt) > 0 {

// 			_str += v + _srt
// 		}
// 	}
// 	_str = secret + _str + secret
// 	_str_byte := []byte(_str)
// 	_rt_has := GetSHA256HashCode(_str_byte)
// 	return _rt_has
// }

// //SHA256生成哈希值
// func GetSHA256HashCode(message []byte) string {
// 	//方法一：
// 	//创建一个基于SHA256算法的hash.Hash接口的对象
// 	hash := sha256.New()
// 	//输入数据
// 	hash.Write(message)
// 	//计算哈希值
// 	bytes := hash.Sum(nil)
// 	//将字符串编码为16进制格式,返回字符串
// 	hashCode := hex.EncodeToString(bytes)
// 	//返回哈希值
// 	return hashCode

// 	//方法二：
// 	//bytes2:=sha256.Sum256(message)//计算哈希值，返回一个长度为32的数组
// 	//hashcode2:=hex.EncodeToString(bytes2[:])//将数组转换成切片，转换成16进制，返回字符串
// 	//return hashcode2
// }
