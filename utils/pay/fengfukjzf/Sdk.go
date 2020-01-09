package fengfukjzf

import (
	"encoding/json"
	"fmt"
	_ "github.com/qnsoft/online_pay/routers"
	"github.com/qnsoft/online_pay/utils/pay/paytool"
	date "github.com/qnsoft/utils/DateHelper"
	"github.com/qnsoft/utils/ErrorHelper"
	"github.com/qnsoft/utils/StringHelper"
	"time"
)

/*
丰付快捷支付业务流程：
支付卡开通发短信 - 确认短信 - 下单 - 下单查询
*/

type FengFuPay struct {
	//接口名称
	Name string `json:"name"`
	//接口版本
	Version string `json:"version"`
}

/*
*出金 （丰付快捷支付接口3.1-下单接口）
@_model 要出金的卡对象
@_model_to 要入金的卡对象
@_channelType 通道标示小额channelType：ffqr大额channelType：ffkj
@_money出金额（元）
@_fee交易手续费 0.5%传0.5 （余额按手续费四舍五入）
*/
func (pay *FengFuPay) CardPay(_model, _model_to paytool.CardModel, _channelType, _city string, _money, _fee, _ewfee float64) Rt31_Model {
	var _rt Rt31_Model
	httpURL := PayUrl + "/preOrderCreate2"
	paramsData := GetBaseParam()
	paramsData["method"] = "preOrderCreate2"                                                                           //调用的方法
	paramsData["merchantId"] = MerchantId                                                                              //商户号
	paramsData["money"] = _money                                                                                       //金额（元）
	paramsData["merOrderNumber"] = "DK" + date.FormatDate(time.Now(), "yyyyMMddHHmmss") + StringHelper.GetRandomNum(4) //商户订单号
	paramsData["accountNumber"] = _model.AccountNumber                                                                 //卡号
	paramsData["tel"] = _model.Phone
	paramsData["fee"] = _fee //交易手续费 0.5%传0.5 （余额按手续费四舍五入）
	if _channelType == "" {
		_channelType = "ffapi"
	}
	paramsData["channelType"] = _channelType                  //通道标示小额channelType：ffqr大额channelType：ffkj
	paramsData["holderName"] = _model.HolderName              //姓名
	paramsData["idcard"] = _model.IdCard                      //身份证
	paramsData["notifyUrl"] = ""                              //异步通知地址
	paramsData["settAccountNumber"] = _model_to.AccountNumber //结算卡
	paramsData["settAccountTel"] = _model_to.Phone            //结算卡预留手机号
	paramsData["extraFee"] = _ewfee                           //额外手续费    最终手续费 = money * fee% + extraFee
	paramsData["city"] = _city                                //落地城市 如上海市
	paramsData["expired"] = _model.Expired                    //有效期YYMM格式
	//paramsData["mcc"] = ""      //Mcc码 多个用逗号隔开
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
func (pay *FengFuPay) CardSms(_model paytool.CardModel, _channelType string) Rt36_Model {
	var _rt Rt36_Model
	httpURL := PayUrl + "/sendSms2OpenToken"
	paramsData := GetBaseParam()
	paramsData["method"] = "sendSms2OpenToken"         //调用的方法
	paramsData["merchantId"] = MerchantId              //商户号
	paramsData["accountNumber"] = _model.AccountNumber //卡号
	paramsData["tel"] = _model.Phone                   //预留手机号
	if _channelType == "" {
		_channelType = "ffapi"
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
func (pay *FengFuPay) CardSmsConfirm(_model paytool.CardModel, _channelType, _merOrderNumber, _smsCode string) Rt37_Model {
	var _rt Rt37_Model
	httpURL := PayUrl + "/checkSms2OpenToken"
	paramsData := GetBaseParam()
	paramsData["method"] = "checkSms2OpenToken"        //调用的方法
	paramsData["merchantId"] = MerchantId              //商户号
	paramsData["accountNumber"] = _model.AccountNumber //卡号
	paramsData["tel"] = _model.Phone                   //预留手机号
	if _channelType == "" {
		_channelType = "ffapi"
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
