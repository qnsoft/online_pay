package xinhuan

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	php2go "github.com/qnsoft/utils/Php2go"
	"github.com/qnsoft/utils/StringHelper"
	"github.com/qnsoft/utils/WebHelper"

	"github.com/wenzhenxi/gorsa"
)

/*
接口请求方法
*/
func WebRequest(_http_url string, _params_data map[string]interface{}) string {
	//定义主机头对象
	_HeaderData := map[string]string{"Accept-Charset": "utf-8", "Content-Type": "application/x-www-form-urlencoded"}
	_data := SignParameters(_params_data, false)
	_data_rsa, _ := Rsa_Pub(_data, string(PublicKey))
	signStr := "data=" + _data_rsa + "&merchantId=" + MerchantId + "&key=" + Key
	sign := php2go.Md5(signStr)
	_Params := map[string]interface{}{"data": _data_rsa, "merchantId": MerchantId, "sign": sign}
	_get_data := ""
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	fmt.Print(_req)
	if strings.Contains(_req, "code=00&message=SUCCESS") {
		_get_data = StringHelper.GetBetweenStr(_req, "SUCCESS&data=", "&merchantId=")
	}
	_json, _ := Rsa_Pri(_get_data, string(PrivateKey))
	fmt.Println(string(_json))
	return _json
}

/*
获取基础参数
*/
func GetBaseParam() map[string]interface{} {
	_rt := make(map[string]interface{})
	_rt["apiVersion"] = 1
	_ts := time.Now().UnixNano() / 1e6
	_rt["curDate"] = strconv.FormatInt(_ts, 10)
	return _rt
}

/*
参数集排序Sign签名 根据key排序
*/
func SignParameters(mp map[string]interface{}, is_url bool) string {
	_str := ""
	var newMp = make([]string, 0)
	for k, _ := range mp {
		newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	for _, v := range newMp {
		_srt := ""
		switch _value := mp[v].(type) {
		case string:
			_srt = _value
		case int:
			_srt = strconv.Itoa(_value)
		case float64:
			_srt = fmt.Sprintf("%.2f", _value)
			//_srt = strconv.ParseFloat(fmt.Sprintf("%.2f", _value), 64)
		}
		if len(_srt) > 0 {
			if is_url {
				_str += v + php2go.URLEncode(_srt)
			} else {
				_str += v + "=" + _srt + "&"
			}
		}
	}
	fmt.Println(_str)
	return php2go.Rtrim(_str, "&")
}

/*
公钥加密
*/
func Rsa_Pub(data, _PublicKey string) (string, error) {
	_str, errrt := gorsa.PublicEncrypt(data, _PublicKey)
	return _str, errrt
}

/*
私钥解密
*/
func Rsa_Pri(data, _PrivateKey string) (string, error) {
	_str, errrt := gorsa.PriKeyDecrypt(data, _PrivateKey)
	fmt.Println(data)
	return _str, errrt
}
