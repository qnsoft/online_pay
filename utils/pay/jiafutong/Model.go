package JiaFuTong

import (
	"sort"
	"strconv"
	php2go "github.com/qnsoft/utils/Php2go"
)

var (
	//基础接口地址测试环境
	//BaseUrl string = "http://test.man-opaydev.ncfgroup.com/fusionPosp/interface/memberReg"
	//基础接口地址生产环境
	BaseUrl string = "https://client-api.verajft.com/fusionPosp/interface"
)

//公共请求参数
type Model_public struct {
	//接口版本号 固定值：1001
	VerCode string `json:"verCode"`
}

//相应报文基础参数
type Return_public struct {
	//通道商户编号
	ChMerCode string `json:"chMerCode"`
	//应答码
	ResCode string `json:"resCode"`
	//应答描述
	ResMsg string `json:"resMsg"`
}

/**
 * 商户注册接口提交对象(接口4.11)
 */
type Model_411 struct {
	Model_public
	//商户编号(渠道的子商户编号，不能长于 16 位)
	MerCode string `json:"merCode"`
	//商户名称
	MerName string `json:"merName"`
	//真实姓名(真实姓名和结算户名必须一致)
	RealName string `json:"realName"`
	//商户详细地址(除省份、城市和区县外的地址信息。)
	MerAddress string `json:"merAddress"`
	//身份证号码
	IdCard string `json:"idCard"`
	//手机号
	Mobile string `json:"mobile"`
	//结算户名(真实姓名和结算户名必须一致)
	AccountName string `json:"accountName"`
	//结算账号(仅为储蓄卡，此处结算卡用于快捷支付业务。注：代偿业务入账卡在交易接口中上送。)
	AccountNo string `json:"accountNo"`
	//结算卡预留手机号(结算卡的预留手机号)
	ReservedMobile string `json:"reservedMobile"`
	//联行号(必输项，标准行号)
	SubBankCode string `json:"subBankCode"`
	//结算户类型(1：公户 2：私户)
	SettleAccType string `json:"settleAccType"`
}

/**
 * 商户注册接口提交返回对象
 */
type Return_411 struct {
	Return_public
	//商户编号
	MerCode string `json:"merCode"`
}

/*
获取基础参数
*/
func GetBaseParam() map[string]interface{} {
	_rt := make(map[string]interface{})
	//接口版本号
	_rt["verCode"] = "1001"
	return _rt
}

/*
参数集排序Sign签名 根据key排序
mp参数列表
is_url是否URLEncode加密
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
		}
		if len(_srt) > 0 {
			if is_url {
				_str += v + "=" + php2go.URLEncode(_srt) + "&"
			} else {
				_str += v + "=" + _srt + "&"
			}
		}
	}
	return _str
}

/*
参数集排序Sign签名 根据key排序
*/
func SignParametersMap(mp map[string]string, is_url bool) map[string]string {
	_map := make(map[string]string)
	var newMp = make([]string, 0)
	for k, _ := range mp {
		newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	for _, v := range newMp {
		if is_url {
			_map[v] = php2go.URLEncode(mp[v])
		} else {
			_map[v] = mp[v]
		}
	}
	return _map
}
