package TongLian

import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/qnsoft/web_api/utils/ErrorHelper"
	"github.com/qnsoft/web_api/utils/WebHelper"
	"github.com/qnsoft/web_api/utils/php2go"
)

//商户进件(文档2.2)
func Addcus(_model Model_22) Return_22 {
	var _rt Return_22
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["belongorgid"] = _model.Belongorgid
	_Params["outcusid"] = _model.Outcusid
	_Params["cusname"] = _model.Cusname
	_Params["cusshortname"] = _model.Cusshortname
	_Params["merprovice"] = _model.Merprovice
	_Params["areacode"] = _model.Areacode
	_Params["legal"] = _model.Legal
	_Params["idno"] = _model.Idno
	_Params["phone"] = _model.Phone
	_Params["address"] = _model.Address
	_Params["acctid"] = _model.Acctid
	_Params["acctname"] = _model.Acctname
	_Params["accttp"] = _model.Accttp
	_Params["expanduser"] = _model.Expanduser
	//_Prodlist, _ := json.Marshal(_model.Prodlist)
	_Params["prodlist"] = "[{'trxcode':'QUICKPAY_OF_HP','feerate':'0.6'}]" //string(_Prodlist)
	_Params["settfee"] = _model.Settfee
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&")
	fmt.Println(_Sign)
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/org/addcus"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//商户进件状态查询(文档2.3)
func Cusquery(_model Model_23) Return_23 {
	var _rt Return_23
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["outcusid"] = _model.Outcusid
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/org/cusquery"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//商户结算/费率信息修改(文档2.4)
func Updatesettinfo(_model Model_24) Return_24 {
	var _rt Return_24
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["acctid"] = _model.Acctid
	_Params["accttp"] = _model.Accttp
	_Params["prodlist"] = _model.Prodlist
	_Params["settfee"] = _model.Settfee
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/org/updatesettinfo"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//商户绑定消费银行卡(文档2.5)
func Bindcard(_model Model_25) Return_25 {
	var _rt Return_25
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["meruserid"] = _model.Meruserid
	_Params["cardno"] = _model.Cardno
	_Params["acctname"] = _model.Acctname
	_Params["accttype"] = _model.Accttype
	_Params["validdate"] = _model.Validdate
	_Params["cvv2"] = _model.Cvv2
	_Params["idno"] = _model.Idno
	_Params["tel"] = _model.Tel
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/org/bindcard"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//重新获取绑卡验证码(文档2.6)
func Resendbindsms(_model Model_26) Return_26 {
	var _rt Return_26
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["meruserid"] = _model.Meruserid
	_Params["cardno"] = _model.Cardno
	_Params["acctname"] = _model.Acctname
	_Params["accttype"] = _model.Accttype
	_Params["validdate"] = _model.Validdate
	_Params["cvv2"] = _model.Cvv2
	_Params["idno"] = _model.Idno
	_Params["tel"] = _model.Tel
	_Params["thpinfo"] = _model.Thpinfo
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/org/resendbindsms"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//商户绑定银行卡确认(文档2.7)
func Bindcardconfirm(_model Model_27) Return_27 {
	var _rt Return_27
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["meruserid"] = _model.Meruserid
	_Params["cardno"] = _model.Cardno
	_Params["acctname"] = _model.Acctname
	_Params["accttype"] = _model.Accttype
	_Params["validdate"] = _model.Validdate
	_Params["cvv2"] = _model.Cvv2
	_Params["idno"] = _model.Idno
	_Params["tel"] = _model.Tel
	_Params["smscode"] = _model.Smscode
	_Params["thpinfo"] = _model.Thpinfo
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/org/bindcardconfirm"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//商户解绑消费银行卡(文档2.8)
func Unbindcard(_model Model_28) Return_28 {
	var _rt Return_28
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["cardno"] = _model.Cardno
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/org/unbindcard"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//快捷交易支付申请(文档2.9)
func Applypay(_model Model_29) Return_29 {
	var _rt Return_29
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["orderid"] = _model.Orderid
	_Params["agreeid"] = _model.Agreeid
	_Params["trxcode"] = _model.Trxcode
	_Params["amount"] = _model.Amount
	_Params["fee"] = _model.Fee
	_Params["currency"] = _model.Currency
	_Params["subject"] = _model.Subject
	_Params["validtime"] = _model.Validtime
	_Params["city"] = _model.City
	_Params["mccid"] = _model.Mccid
	_Params["trxreserve"] = _model.Trxreserve
	_Params["notifyurl"] = _model.Notifyurl
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/qpay/applypay"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//快捷交易支付确认(文档2.10)
func Confirmpay(_model Model_210) Return_210 {
	var _rt Return_210
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["trxid"] = _model.Trxid
	_Params["agreeid"] = _model.Agreeid
	_Params["smscode"] = _model.Smscode
	_Params["thpinfo"] = _model.Thpinfo
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/qpay/confirmpay"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//快捷支付短信重新获取(文档2.11)
func Paysms(_model Model_211) Return_211 {
	var _rt Return_211
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["trxid"] = _model.Trxid
	_Params["agreeid"] = _model.Agreeid
	_Params["thpinfo"] = _model.Thpinfo
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/qpay/paysms"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//快捷小额免短信支付(文档2.12)
func Quickpass(_model Model_212) Return_212 {
	var _rt Return_212
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["orderid"] = _model.Orderid
	_Params["agreeid"] = _model.Agreeid
	_Params["trxcode"] = _model.Trxcode
	_Params["amount"] = _model.Amount
	_Params["fee"] = _model.Fee
	_Params["currency"] = _model.Currency
	_Params["subject"] = _model.Subject
	_Params["validtime"] = _model.Validtime
	_Params["city"] = _model.City
	_Params["mccid"] = _model.Mccid
	_Params["trxreserve"] = _model.Trxreserve
	_Params["notifyurl"] = _model.Notifyurl
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/qpay/quickpass"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//快捷交易查询(文档2.13)
func Query(_model Model_213) Return_213 {
	var _rt Return_213
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["orderid"] = _model.Orderid
	_Params["trxid"] = _model.Trxid
	_Params["date"] = _model.Date
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/qpay/query"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//余额查询(文档2.14)
func Balance(_model Model_214) Return_214 {
	var _rt Return_214
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/acct/balance"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//提现(文档2.15)
func Withdraw(_model Model_215) Return_215 {
	var _rt Return_215
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["orderid"] = _model.Orderid
	_Params["amount"] = _model.Amount
	_Params["isall"] = _model.Isall
	_Params["fee"] = _model.Fee
	_Params["trxreserve"] = _model.Trxreserve
	_Params["notifyurl"] = _model.Notifyurl
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/acct/withdraw"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//付款(文档2.16)
func Pay(_model Model_216) Return_216 {
	var _rt Return_216
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["orderid"] = _model.Orderid
	_Params["amount"] = _model.Amount
	_Params["isall"] = _model.Isall
	_Params["fee"] = _model.Fee
	_Params["agreeid"] = _model.Agreeid
	_Params["trxreserve"] = _model.Trxreserve
	_Params["notifyurl"] = _model.Notifyurl
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/acct/pay"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//提现(付款)交易查询返回对象(文档2.17)
func Querypay(_model Model_217) Return_217 {
	var _rt Return_217
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["orderid"] = _model.Orderid
	_Params["trxid"] = _model.Trxid
	_Params["date"] = _model.Date
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/acct/querypay"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

/* 219 是异步回传，且参数是变参 所以不定义实体 */

//对账文件下载(文档2.19)
func Download(_model Model_219) Return_219 {
	var _rt Return_219
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["trxdate"] = _model.Trxdate
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/checkacct/download"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//查询绑卡协议号(文档2.20)
func Getagree(_model Model_220) Return_220 {
	var _rt Return_220
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["cusid"] = _model.Cusid
	_Params["cardno"] = _model.Cardno
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/org/getagree"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//查询地区对应可用行业分类代码(文档2.21)
func FindMccByAreacode(_model Model_221) Return_221 {
	var _rt Return_221
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["areacode"] = _model.Areacode
	_Params["trxcode"] = _model.Trxcode
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") //参数排序
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["sign"] = _SignParameters
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/dict/findMccByAreacode"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}
