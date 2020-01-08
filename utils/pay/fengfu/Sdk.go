package fengfu

import (
	"encoding/json"
	"fmt"
	"time"
	_ "github.com/qnsoft/online_pay/routers"
	"github.com/qnsoft/online_pay/utils/pay/paytool"
	date "github.com/qnsoft/web_api/utils/DateHelper"
	"github.com/qnsoft/web_api/utils/ErrorHelper"
	"github.com/qnsoft/web_api/utils/StringHelper"
)

/*
丰付业务流程：
大额： 支付卡开通发短信 - 确认短信 - 下单 - 下单查询  - 结算 - 结算查询 - 查询余额
小额： 下单 - 下单查询  - 结算 - 结算查询 - 查询余额
*/

type FengFuPay struct {
	//接口名称
	Name string `json:"name"`
	//接口版本
	Version string `json:"version"`
}

/*
*出金 对应（丰付接口3.1-下单接口）
@_model 要出金的卡对象
@_channelType 通道标示小额channelType：ffqr大额channelType：ffkj
@_money出金额（元）
@_fee交易手续费 0.5%传0.5 （余额按手续费四舍五入）
*/
func (pay *FengFuPay) CardOut(_model paytool.CardModel, _channelType, _city string, _money, _fee float64) Rt31_Model {
	var _rt Rt31_Model
	httpURL := PayUrl + "/preOrderCreate"
	paramsData := GetBaseParam()
	paramsData["method"] = "preOrderCreate"                                                                            //调用的方法
	paramsData["merchantId"] = MerchantId                                                                              //商户号
	paramsData["money"] = _money                                                                                       //金额（元）
	paramsData["merOrderNumber"] = "DK" + date.FormatDate(time.Now(), "yyyyMMddHHmmss") + StringHelper.GetRandomNum(4) //商户订单号
	paramsData["accountNumber"] = _model.AccountNumber                                                                 //卡号
	paramsData["tel"] = _model.Phone
	paramsData["fee"] = _fee //交易手续费 0.5%传0.5 （余额按手续费四舍五入）
	if _channelType == "" {
		_channelType = "ffqr"
	}
	paramsData["channelType"] = _channelType     //通道标示小额channelType：ffqr大额channelType：ffkj
	paramsData["holderName"] = _model.HolderName //姓名
	paramsData["idcard"] = _model.IdCard         //身份证
	paramsData["city"] = _city                   //落地城市 如上海市
	//paramsData["notifyUrl"] = ""        //异步通知地址
	paramsData["cvv2"] = _model.Cvv2       //cvv
	paramsData["expired"] = _model.Expired //有效期YYMM格式
	//paramsData["mcc"] = ""      //Mcc码 多个用逗号隔开
	_json := WebRequest(httpURL, paramsData)
	err := json.Unmarshal([]byte(_json), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

/*
*出金订单查询 对应（丰付接口3.2-查询订单状态接口）
@_merOrderNumber 商户订单号
*/
func (pay *FengFuPay) CardOutOrder(_model paytool.CardModel, _merOrderNumber string) Rt32_Model {
	var _rt Rt32_Model
	httpURL := PayUrl + "/queryOrderStatus"
	paramsData := GetBaseParam()
	paramsData["method"] = "queryOrderStatus"      //调用的方法
	paramsData["merchantId"] = MerchantId          //商户号                                                                            //金额（元）
	paramsData["merOrderNumber"] = _merOrderNumber //商户订单号
	_json := WebRequest(httpURL, paramsData)
	err := json.Unmarshal([]byte(_json), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

/*
*入金 对应（丰付接口3.3-结算接口）
@_model 要入金的卡对象
@_channelType 通道标示小额channelType：ffqr大额channelType：ffkj
@_money出金额（元）
@_extraFee 结算手续费
*/
func (pay *FengFuPay) CardIn(_model paytool.CardModel, _channelType string, _money, _extraFee float64) Rt33_Model {
	var _rt Rt33_Model
	httpURL := PayUrl + "/checkOutOrder"
	paramsData := GetBaseParam()
	paramsData["method"] = "checkOutOrder"                                                                             //调用的方法
	paramsData["merchantId"] = MerchantId                                                                              //商户号
	paramsData["money"] = _money                                                                                       //金额（元）
	paramsData["merOrderNumber"] = "DH" + date.FormatDate(time.Now(), "yyyyMMddHHmmss") + StringHelper.GetRandomNum(4) //商户订单号
	paramsData["accountNumber"] = _model.AccountNumber                                                                 //卡号
	paramsData["tel"] = _model.Phone
	paramsData["extraFee"] = _extraFee //交易手续费 0.5%传0.5 （余额按手续费四舍五入）
	if _channelType == "" {
		_channelType = "ffqr"
	}
	paramsData["channelType"] = _channelType     //通道标示小额channelType：ffqr大额channelType：ffkj
	paramsData["holderName"] = _model.HolderName //姓名
	paramsData["idcard"] = _model.IdCard         //身份证
	//paramsData["notifyUrl"] = ""        //异步通知地址
	_json := WebRequest(httpURL, paramsData)
	err := json.Unmarshal([]byte(_json), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

/*
*入金订单查询 对应（丰付接口3.4- 结算状态查询接口）
@_merOrderNumber 商户订单号
*/
func (pay *FengFuPay) CardInOrder(_merOrderNumber string) Rt34_Model {
	var _rt Rt34_Model
	httpURL := PayUrl + "/queryCheckOutOrderStatus"
	paramsData := GetBaseParam()
	paramsData["method"] = "queryCheckOutOrderStatus" //调用的方法
	paramsData["merchantId"] = MerchantId             //商户号                                                                            //金额（元）
	paramsData["merOrderNumber"] = _merOrderNumber    //商户订单号
	_json := WebRequest(httpURL, paramsData)
	err := json.Unmarshal([]byte(_json), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

/*
*余额查询 对应（丰付接口3.5- 余额查询接口）
@_channelType 通道标示小额channelType：ffqr大额channelType：ffkj
@_accountNumber 卡号
*/
func (pay *FengFuPay) CardBanlance(_channelType, _accountNumber string) Rt35_Model {
	var _rt Rt35_Model
	httpURL := PayUrl + "/queryUserBanlance"
	paramsData := GetBaseParam()
	paramsData["method"] = "queryUserBanlance"   //调用的方法
	paramsData["merchantId"] = MerchantId        //商户号                                                                            //金额（元）
	paramsData["accountNumber"] = _accountNumber //卡号
	paramsData["channelType"] = _channelType     //通道标示小额channelType：ffqr大额channelType：ffkj
	_json := WebRequest(httpURL, paramsData)
	err := json.Unmarshal([]byte(_json), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

/*
*卡片签约 对应（丰付接口3.6-支付卡开通发短信接口）
@_model 要签约的卡对象
@_channelType 通道标示小额channelType：ffqr大额channelType：ffkj
*/
func (pay *FengFuPay) CardSign(_model paytool.CardModel, _channelType string) Rt36_Model {
	var _rt Rt36_Model
	httpURL := PayUrl + "/sendSms2OpenToken"
	paramsData := GetBaseParam()
	paramsData["method"] = "sendSms2OpenToken"         //调用的方法
	paramsData["merchantId"] = MerchantId              //商户号
	paramsData["accountNumber"] = _model.AccountNumber //卡号
	paramsData["tel"] = _model.Phone                   //预留手机号
	if _channelType == "" {
		_channelType = "ffqrdh"
	}
	paramsData["channelType"] = _channelType                                                                    //通道标示小额channelType：ffqr大额channelType：ffkj
	paramsData["holderName"] = _model.HolderName                                                                //姓名
	paramsData["idcard"] = _model.IdCard                                                                        //身份证
	paramsData["cvv2"] = _model.Cvv2                                                                            //cvv
	paramsData["expired"] = _model.Expired                                                                      //有效期YYMM格式
	paramsData["merOrderNumber"] = date.FormatDate(time.Now(), "yyyyMMddHHmmss") + StringHelper.GetRandomNum(4) //订单号
	fmt.Print(paramsData)
	_json := WebRequest(httpURL, paramsData)
	err := json.Unmarshal([]byte(_json), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

/*
*卡片签约确认 对应（丰付接口3.7-支付卡开通确认短信接口）
@_model 要签约的卡对象
@_channelType 通道标示小额channelType：ffqr大额channelType：ffkj
*/
func (pay *FengFuPay) CardSignConfirm(_model paytool.CardModel, _channelType, _merOrderNumber, _smsCode string) Rt37_Model {
	var _rt Rt37_Model
	httpURL := PayUrl + "/checkSms2OpenToken"
	paramsData := GetBaseParam()
	paramsData["method"] = "checkSms2OpenToken"        //调用的方法
	paramsData["merchantId"] = MerchantId              //商户号
	paramsData["accountNumber"] = _model.AccountNumber //卡号
	paramsData["tel"] = _model.Phone                   //预留手机号
	if _channelType == "" {
		_channelType = "ffkj"
	}
	paramsData["channelType"] = _channelType       //通道标示小额channelType：ffqr大额channelType：ffkj
	paramsData["holderName"] = _model.HolderName   //姓名
	paramsData["idcard"] = _model.IdCard           //身份证
	paramsData["cvv2"] = _model.Cvv2               //cvv
	paramsData["expired"] = _model.Expired         //有效期YYMM格式
	paramsData["merOrderNumber"] = _merOrderNumber //订单号
	paramsData["smsCode"] = _smsCode               //根据3.6提交手机短信获取
	_json := WebRequest(httpURL, paramsData)
	err := json.Unmarshal([]byte(_json), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

// func FF_36() {
// 	//3.6. 支付卡开通发短信接口
// 	_http_url := PayUrl + "/sendSms2OpenToken"
// 	//定义主机头对象
// 	_HeaderData := map[string]string{"Accept-Charset": "utf-8", "Content-Type": "application/x-www-form-urlencoded"}
// 	_params_data := GetBaseParam()
// 	_params_data["method"] = "sendSms2OpenToken"
// 	_params_data["merchantId"] = MerchantId
// 	_params_data["accountNumber"] = "6258091711738406"
// 	_params_data["tel"] = "13938202388"
// 	_params_data["channelType"] = "ffkj" //类型小额channelType：ffqr大额channelType：ffkj
// 	_params_data["holderName"] = "宋晓辉"
// 	_params_data["idcard"] = "410185198209190514"
// 	_params_data["cvv2"] = "903"
// 	_params_data["expired"] = "0423"
// 	_params_data["merOrderNumber"] = date.FormatDate(time.Now(), "yyyyMMddHHmmss") + StringHelper.GetRandomNum(2)
// 	_data := SignParameters(_params_data, false)
// 	_data_rsa, _ := Rsa_Pub(_data, string(PublicKey))
// 	signStr := "data=" + _data_rsa + "&merchantId=" + MerchantId + "&key=" + Key
// 	sign := php2go.Md5(signStr)
// 	_Params := map[string]interface{}{"data": _data_rsa, "merchantId": MerchantId, "sign": sign}
// 	_get_data := ""
// 	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
// 	if strings.Contains(_req, "code=00&message=SUCCESS") {
// 		_get_data = StringHelper.GetBetweenStr(_req, "SUCCESS&data=", "&merchantId=")
// 	}
// 	_json, _ := Rsa_Pri(_get_data, string(PrivateKey))
// 	fmt.Println(string(_json))
// }

// func FF_37() {
// 	_http_url := PayUrl + "/checkSms2OpenToken"
// 	//定义主机头对象
// 	_HeaderData := map[string]string{"Accept-Charset": "utf-8", "Content-Type": "application/x-www-form-urlencoded"}
// 	_params_data := GetBaseParam()
// 	_params_data["method"] = "checkSms2OpenToken"
// 	_params_data["merchantId"] = MerchantId
// 	_params_data["accountNumber"] = "6258091711738406"
// 	_params_data["tel"] = "13938202388"
// 	_params_data["channelType"] = "ffkj" //类型小额channelType：ffqr大额channelType：ffkj
// 	_params_data["holderName"] = "宋晓辉"
// 	_params_data["idcard"] = "410185198209190514"
// 	_params_data["cvv2"] = "903"
// 	_params_data["expired"] = "0423"
// 	_params_data["merOrderNumber"] = "2019110510301919" //根据3.6获取
// 	_params_data["smsCode"] = "379854"                  //根据3.6提交手机短信获取
// 	_data := SignParameters(_params_data, false)
// 	_data_rsa, _ := Rsa_Pub(_data, string(PublicKey))
// 	signStr := "data=" + _data_rsa + "&merchantId=" + MerchantId + "&key=" + Key
// 	sign := php2go.Md5(signStr)
// 	_Params := map[string]interface{}{"data": _data_rsa, "merchantId": MerchantId, "sign": sign}
// 	_get_data := ""
// 	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
// 	if strings.Contains(_req, "code=00&message=SUCCESS") {
// 		_get_data = StringHelper.GetBetweenStr(_req, "SUCCESS&data=", "&merchantId=")
// 	}
// 	_json, _ := Rsa_Pri(_get_data, string(PrivateKey))
// 	fmt.Println(string(_json))
// }
