package GongYiFuNew

import (
	"encoding/json"
	"fmt"
	"github.com/qnsoft/utils/ErrorHelper"
	"github.com/qnsoft/utils/WebHelper"

	"github.com/goinggo/mapstructure"
)

/*
业务处理公共函数封装
@_requestName 业务方法名
@_bll_Params 业务参数
*/
func Tool(_requestName string, _bll_Params map[string]interface{}) interface{} {
	var _rt interface{}
	//开始处理业务参数
	_send_json, _ := json.Marshal(_bll_Params)
	// //业务参数处理结束
	/******************************** 商户私钥对业务数据签名 */
	sign, _ := SignMsg([]byte(_send_json), PrivateKey)
	//******************************** GYF平台公钥加密业务数据 */
	encryptDataJson := EncryptData([]byte(_send_json), PublicKey)

	//fmt.Println("要传的数据sign", sign)
	//fmt.Println("要传的数据encryptDataJson", encryptDataJson)

	_Params := GetBaseParam()
	_Params["data"] = encryptDataJson //请求报文体
	_Params["sign"] = sign            //_sign数据签名
	//业务参数结束
	//_http_url := PayUrl + _requestName
	_http_url := "http://test.bantpay.com/api/sub_merchant_register"
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	//_HeaderData"Content-Type": "application/json"}
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	fmt.Printf("%#V", _req)
	ErrorHelper.CheckErr(err)
	return _rt
}

/*
子商户进件(这里统称为用户进件)
*/
func JinJian(_model Model_jinjian) Model_jinjian_Return {
	//开始组合业务参数
	_bll_Params := make(map[string]interface{})
	if _model.BankAccountNo != "" && len(_model.BankAccountNo) > 0 {
		_bll_Params["bankAccountNo"] = _model.BankAccountNo
	}
	if _model.CertificateNo != "" && len(_model.CertificateNo) > 0 {
		_bll_Params["certificateNo"] = _model.CertificateNo
	}
	if _model.Mobile != "" && len(_model.Mobile) > 0 {
		_bll_Params["mobile"] = _model.Mobile
	}
	if _model.Remark != "" && len(_model.Remark) > 0 {
		_bll_Params["remark"] = _model.Remark
	}
	if _model.SubMerchantFeeRate != "" && len(_model.SubMerchantFeeRate) > 0 {
		_bll_Params["subMerchantFeeRate"] = _model.SubMerchantFeeRate
	}
	if _model.SubMerchantName != "" && len(_model.SubMerchantName) > 0 {
		_bll_Params["subMerchantName"] = _model.SubMerchantName
	}
	if _model.SubMerchantWithdrawFee != "" && len(_model.SubMerchantWithdrawFee) > 0 {
		_bll_Params["subMerchantWithdrawFee"] = _model.SubMerchantWithdrawFee
	}
	_rt := Tool("sub_merchant_register", _bll_Params) //订单查询传参
	mapInstance := Jiemi_data(_rt)
	var rt_101 Model_jinjian_Return
	err1 := mapstructure.Decode(mapInstance, &rt_101)
	ErrorHelper.CheckErr(err1)
	return rt_101
}

/*
工易付S102商户信息修
*/
func Gyf_S102(_model Model_102) Model_Return_102 {
	//开始组合业务参数
	_bll_Params := make(map[string]interface{})
	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
		_bll_Params["subMerchId"] = _model.SubMerchId
	}
	if _model.Phone != "" && len(_model.Phone) > 0 {
		_bll_Params["phone"] = _model.Phone
	}
	if _model.CardId != "" && len(_model.CardId) > 0 {
		_bll_Params["cardId"] = _model.CardId
	}
	_rt := Tool("SdkMerchSettleModify", _bll_Params) //订单查询传参
	mapInstance := Jiemi_data(_rt)
	var rt_102 Model_Return_102
	err1 := mapstructure.Decode(mapInstance, &rt_102)
	ErrorHelper.CheckErr(err1)
	return rt_102
}

/*
工易付S103.商户信息费率
*/
func Gyf_S103(_model Model_103) Model_Return_103 {
	//开始组合业务参数
	_bll_Params := make(map[string]interface{})
	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
		_bll_Params["subMerchId"] = _model.SubMerchId
	}
	if _model.FeeRate != "" && len(_model.FeeRate) > 0 {
		_bll_Params["feeRate"] = _model.FeeRate
	}
	if _model.ExternFee != "" && len(_model.ExternFee) > 0 {
		_bll_Params["externFee"] = _model.ExternFee
	}
	_rt := Tool("SdkMerchRateModify", _bll_Params) //订单查询传参
	mapInstance := Jiemi_data(_rt)
	var rt_103 Model_Return_103
	err1 := mapstructure.Decode(mapInstance, &rt_103)
	ErrorHelper.CheckErr(err1)
	return rt_103
}

/*
工易付S201.银联页面绑卡
*/
func Gyf_S201(_model Model_201) Model_Return_201 {
	//开始组合业务参数
	_bll_Params := make(map[string]interface{})
	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
		_bll_Params["subMerchId"] = _model.SubMerchId
	}
	if _model.Name != "" && len(_model.Name) > 0 {
		_bll_Params["name"] = _model.Name
	}
	if _model.Phone != "" && len(_model.Phone) > 0 {
		_bll_Params["phone"] = _model.Phone
	}
	if _model.IdNo != "" && len(_model.IdNo) > 0 {
		_bll_Params["idNo"] = _model.IdNo
	}
	if _model.CardId != "" && len(_model.CardId) > 0 {
		_bll_Params["cardId"] = _model.CardId
	}
	if _model.NotifyUrl != "" && len(_model.NotifyUrl) > 0 {
		_bll_Params["notifyUrl"] = _model.NotifyUrl
	}
	if _model.FrontUrl != "" && len(_model.FrontUrl) > 0 {
		_bll_Params["frontUrl"] = _model.FrontUrl
	}
	if _model.OrderId != "" && len(_model.OrderId) > 0 {
		_bll_Params["orderId"] = _model.OrderId
	}
	if _model.DeviceId != "" && len(_model.DeviceId) > 0 {
		_bll_Params["deviceId"] = _model.DeviceId
	}
	if _model.IpAddress != "" && len(_model.IpAddress) > 0 {
		_bll_Params["ipAddress"] = _model.IpAddress
	}
	_rt := Tool("SdkBindCardH5", _bll_Params) //订单查询传参
	mapInstance := Jiemi_data(_rt)
	var rt_201 Model_Return_201
	err1 := mapstructure.Decode(mapInstance, &rt_201)
	ErrorHelper.CheckErr(err1)
	return rt_201
}

/*
工易付S202.银联页面绑卡
*/
func Gyf_S202(_model Model_202) Model_Return_202 {
	//开始组合业务参数
	_bll_Params := make(map[string]interface{})
	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
		_bll_Params["subMerchId"] = _model.SubMerchId
	}
	if _model.CardId != "" && len(_model.CardId) > 0 {
		_bll_Params["cardId"] = _model.CardId
	}
	_rt := Tool("SdkQueryBindCard", _bll_Params) //订单查询传参
	mapInstance := Jiemi_data(_rt)
	var rt_202 Model_Return_202
	err1 := mapstructure.Decode(mapInstance, &rt_202)
	ErrorHelper.CheckErr(err1)
	return rt_202
}

/*
工易付S301订单查询接口
*/
func Gyf_S301(_model Model_301) Model_Return_301 {
	//var _rt interface{}
	//开始组合业务参数
	_bll_Params := make(map[string]interface{})
	_bll_Params["orderId"] = _model.OrderId //"ZCM2019092516250991"
	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
		_bll_Params["subMerchId"] = _model.SubMerchId //"SMCH99990212558313"
	}
	_rt := Tool("SdkQueryOrder", _bll_Params) //订单查询传参
	// //开始处理业务参数
	// _send_json, _ := json.Marshal(_bll_Params) //1.组成请求生成明文数据

	// //_send_json := "{\"subMerchId\":\"SMCH99990212558313\",\"orderId\":\"ZCM2019092516250991\"}"
	// //encryptkey := strconv.Itoa(php2go.Rand(10000000, 49999999)) + strconv.Itoa(php2go.Rand(50000000, 99999999)) //2.生成请求随机密码AesKey(由数字和组成的16位定长字符串)
	// AesKey := date.FormatDate(time.Now(), "yyyyMMddHHmmss") + StringHelper.GetRandomNum(2) //2.生成请求随机密码AesKey(由数字和组成的16位定长字符串)
	// //fmt.Println(encryptkey) //StringHelper.AesEncryptECB
	// dataStr := StringHelper.AesEncryptECB([]byte(_send_json), []byte(AesKey)) //3.用生成的随机密码AesKey对明文数据做AES加密（BCD编码）之后得到加密data（data就是报文传输中data的值）
	// dataStr_a := strings.ToUpper(hex.EncodeToString(dataStr))
	// _rsaEncryptkey, _ := Rsa_Encode([]byte(AesKey), PrivateKey) //4.用RSA私钥文件对AesKey进行RSA加密(Base64格式编码输出)之后得到encryptkey。
	// //fmt.Println(err2)
	// _sign := php2go.Md5(dataStr_a + Key) //数据签名
	// //业务参数处理结束
	// _Params := GetBaseParam()
	// _Params["requestName"] = "SdkQueryOrder" //接口服务名称
	// _Params["data"] = dataStr_a              //请求报文体（AES加密) 这一步已没问题
	// _Params["encryptkey"] = _rsaEncryptkey   //数据体加密KEY (Base64格式编码输出)之后得到encryptkey 这一步也没问题
	// _Params["sign"] = _sign                  //_sign//数据签名
	// //业务参数结束
	// _http_url := PayUrl
	// //fmt.Println("提交的业务参数：", string(_send_json))
	// //_send_json_all, _ := json.Marshal(_Params)
	// //fmt.Println("提交的所有参数：", string(_send_json_all))
	// _HeaderData := map[string]string{"Accept-Charset": "utf-8", "Content-Type": "application/x-www-form-urlencoded"}
	// _req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	// //fmt.Println("返回结果：", _req)
	// err := json.Unmarshal([]byte(_req), &_rt)
	// ErrorHelper.CheckErr(err)
	mapInstance := Jiemi_data(_rt)
	var rt_301 Model_Return_301
	//将 map 转换为指定的结构体
	err1 := mapstructure.Decode(mapInstance, &rt_301)
	ErrorHelper.CheckErr(err1)
	return rt_301
}

/*
工易付S302.快捷支付
*/
func Gyf_S302(_model Model_302) Model_Return_302 {
	//开始组合业务参数
	_bll_Params := make(map[string]interface{})
	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
		_bll_Params["subMerchId"] = _model.SubMerchId
	}
	if _model.OrderId != "" && len(_model.OrderId) > 0 {
		_bll_Params["orderId"] = _model.OrderId
	}
	if _model.Name != "" && len(_model.Name) > 0 {
		_bll_Params["name"] = _model.Name
	}
	if _model.Phone != "" && len(_model.Phone) > 0 {
		_bll_Params["phone"] = _model.Phone
	}
	if _model.IdNo != "" && len(_model.IdNo) > 0 {
		_bll_Params["idNo"] = _model.IdNo
	}
	if _model.CardId != "" && len(_model.CardId) > 0 {
		_bll_Params["cardId"] = _model.CardId
	}
	if _model.NotifyUrl != "" && len(_model.NotifyUrl) > 0 {
		_bll_Params["notifyUrl"] = _model.NotifyUrl
	}
	if _model.Amount != "" && len(_model.Amount) > 0 {
		_bll_Params["amount"] = _model.Amount
	}
	if _model.GoodsName != "" && len(_model.GoodsName) > 0 {
		_bll_Params["goodsName"] = _model.GoodsName
	}
	if _model.CardType != "" && len(_model.CardType) > 0 {
		_bll_Params["cardType"] = _model.CardType
	}
	if _model.Cvv != "" && len(_model.Cvv) > 0 {
		_bll_Params["cvv"] = _model.Cvv
	}
	if _model.ExpDate != "" && len(_model.ExpDate) > 0 {
		_bll_Params["expDate"] = _model.ExpDate
	}
	if _model.DeviceId != "" && len(_model.DeviceId) > 0 {
		_bll_Params["deviceId"] = _model.DeviceId
	}
	if _model.IpAddress != "" && len(_model.IpAddress) > 0 {
		_bll_Params["ipAddress"] = _model.IpAddress
	}
	_rt := Tool("SdkNoCardPay", _bll_Params) //订单查询传参
	mapInstance := Jiemi_data(_rt)
	var rt_302 Model_Return_302
	err1 := mapstructure.Decode(mapInstance, &rt_302)
	ErrorHelper.CheckErr(err1)
	return rt_302
}

/*
工易付S302.落地云闪付支付
*/
func Gyf_S302_A(_model Model_302_A) Model_Return_302_A {
	//开始组合业务参数
	_bll_Params := make(map[string]interface{})
	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
		_bll_Params["subMerchId"] = _model.SubMerchId
	}
	if _model.OrderId != "" && len(_model.OrderId) > 0 {
		_bll_Params["orderId"] = _model.OrderId
	}
	if _model.Name != "" && len(_model.Name) > 0 {
		_bll_Params["name"] = _model.Name
	}
	if _model.Phone != "" && len(_model.Phone) > 0 {
		_bll_Params["phone"] = _model.Phone
	}
	if _model.IdNo != "" && len(_model.IdNo) > 0 {
		_bll_Params["idNo"] = _model.IdNo
	}
	if _model.CardId != "" && len(_model.CardId) > 0 {
		_bll_Params["cardId"] = _model.CardId
	}
	if _model.NotifyUrl != "" && len(_model.NotifyUrl) > 0 {
		_bll_Params["notifyUrl"] = _model.NotifyUrl
	}
	if _model.Amount != "" && len(_model.Amount) > 0 {
		_bll_Params["amount"] = _model.Amount
	}
	if _model.GoodsName != "" && len(_model.GoodsName) > 0 {
		_bll_Params["goodsName"] = _model.GoodsName
	}
	if _model.CardType != "" && len(_model.CardType) > 0 {
		_bll_Params["cardType"] = _model.CardType
	}
	if _model.Cvv != "" && len(_model.Cvv) > 0 {
		_bll_Params["cvv"] = _model.Cvv
	}
	if _model.ExpDate != "" && len(_model.ExpDate) > 0 {
		_bll_Params["expDate"] = _model.ExpDate
	}
	if _model.DeviceId != "" && len(_model.DeviceId) > 0 {
		_bll_Params["deviceId"] = _model.DeviceId
	}
	if _model.IpAddress != "" && len(_model.IpAddress) > 0 {
		_bll_Params["ipAddress"] = _model.IpAddress
	}
	if _model.CityCode != "" && len(_model.CityCode) > 0 {
		_bll_Params["cityCode"] = _model.CityCode
	}
	_rt := Tool("SdkUnionQuickPay", _bll_Params) //订单查询传参
	mapInstance := Jiemi_data(_rt)
	var rt_302_A Model_Return_302_A
	err1 := mapstructure.Decode(mapInstance, &rt_302_A)
	ErrorHelper.CheckErr(err1)
	return rt_302_A
}

/*
工易付S302.子商户结算
*/
func Gyf_S302_B(_model Model_302_B) Model_Return_302_B {
	//开始组合业务参数
	_bll_Params := make(map[string]interface{})
	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
		_bll_Params["subMerchId"] = _model.SubMerchId
	}
	if _model.OrderId != "" && len(_model.OrderId) > 0 {
		_bll_Params["orderId"] = _model.OrderId
	}
	if _model.Name != "" && len(_model.Name) > 0 {
		_bll_Params["name"] = _model.Name
	}
	if _model.Phone != "" && len(_model.Phone) > 0 {
		_bll_Params["phone"] = _model.Phone
	}
	if _model.IdNo != "" && len(_model.IdNo) > 0 {
		_bll_Params["idNo"] = _model.IdNo
	}
	if _model.CardId != "" && len(_model.CardId) > 0 {
		_bll_Params["cardId"] = _model.CardId
	}
	if _model.NotifyUrl != "" && len(_model.NotifyUrl) > 0 {
		_bll_Params["notifyUrl"] = _model.NotifyUrl
	}
	if _model.Amount != "" && len(_model.Amount) > 0 {
		_bll_Params["amount"] = _model.Amount
	}
	if _model.Remark != "" && len(_model.Remark) > 0 {
		_bll_Params["remark"] = _model.Remark
	}

	_rt := Tool("SdkSettleSubMerch", _bll_Params) //代付结算传参
	mapInstance := Jiemi_data(_rt)
	var rt_302_B Model_Return_302_B
	err1 := mapstructure.Decode(mapInstance, &rt_302_B)
	ErrorHelper.CheckErr(err1)
	return rt_302_B
}

/*
工易付S302.子商户余额查询
*/
func Gyf_S302_C(_model Model_302_C) Model_Return_302_C {
	//开始组合业务参数
	_bll_Params := make(map[string]interface{})
	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
		_bll_Params["subMerchId"] = _model.SubMerchId
	}
	_rt := Tool("SdkQuerySubMerchBalance", _bll_Params) //订单查询传参
	mapInstance := Jiemi_data(_rt)
	var rt_302_C Model_Return_302_C
	err1 := mapstructure.Decode(mapInstance, &rt_302_C)
	ErrorHelper.CheckErr(err1)
	return rt_302_C
}

/*
工易付S401.绑卡异步通知
*/
// func Gyf_S401(_model Model_401) Model_Return_401 {
// 	//开始组合业务参数
// 	_bll_Params := make(map[string]interface{})
// 	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
// 		_bll_Params["subMerchId"] = _model.SubMerchId
// 	}
// 	_rt := Tool("SdkQuerySubMerchBalance", _bll_Params) //订单查询传参
// 	mapInstance := Jiemi_data(_rt)
// 	var rt_401 Model_Return_401
// 	err1 := mapstructure.Decode(mapInstance, &rt_401)
// 	ErrorHelper.CheckErr(err1)
// 	return rt_401
// }

/*
工易付S402.交易异步通知
*/
// func Gyf_S401(_model Model_401) Model_Return_401 {
// 	//开始组合业务参数
// 	_bll_Params := make(map[string]interface{})
// 	if _model.SubMerchId != "" && len(_model.SubMerchId) > 0 {
// 		_bll_Params["subMerchId"] = _model.SubMerchId
// 	}
// 	_rt := Tool("SdkQuerySubMerchBalance", _bll_Params) //订单查询传参
// 	mapInstance := Jiemi_data(_rt)
// 	var rt_401 Model_Return_401
// 	err1 := mapstructure.Decode(mapInstance, &rt_401)
// 	ErrorHelper.CheckErr(err1)
// 	return rt_401
// }
/*
S501.地区编码查询
*/
func Gyf_S501() interface{} {
	var _rt interface{}
	_http_url := "http://pay.weiyifu123.com/gateway/getCityMerchConfigListName"
	_HeaderData := map[string]string{"Content-Type": "text/json"}
	_req := WebHelper.HttpGet(_http_url, _HeaderData, nil)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}
