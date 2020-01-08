package JiaFuTong

import (
	"encoding/json"
	"fmt"
	"strings"
	"qnsoft/web_api/utils/ErrorHelper"
	php2go "qnsoft/web_api/utils/Php2go"
	"qnsoft/web_api/utils/StringHelper"
	"qnsoft/web_api/utils/WebHelper"
)

/*
佳付通sdk
*/
//商户注册(文档4.1.1)
func MemberReg(_model Model_411) Return_411 {
	var _rt Return_411
	_Params := GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["merCode"] = _model.MerCode
	_Params["merName"] = _model.MerName
	_Params["realName"] = _model.RealName
	_Params["merAddress"] = _model.MerAddress
	_Params["idCard"] = _model.IdCard
	_Params["mobile"] = _model.Mobile
	_Params["accountName"] = _model.AccountName
	_Params["accountNo"] = _model.AccountNo
	_Params["reservedMobile"] = _model.ReservedMobile
	_Params["subBankCode"] = _model.SubBankCode
	_Params["settleAccType"] = _model.SettleAccType
	/*-----------开始签名------------------*/
	_Sign := php2go.Rtrim(SignParameters(_Params, false), "&") + "&md5key=PB74PU3S6JFDN5M6"
	fmt.Println(_Sign)
	_SignParameters := strings.ToUpper(php2go.Md5(_Sign))
	_Params["orgCode"] = "LJZ001"
	_Params["signData"] = _SignParameters                           //数据签名                                 //机构编号
	_key := []byte("PB74PU3S6JFDN5M6")                              //加密的密钥
	_encryptData := StringHelper.AesEncryptCBC([]byte(_Sign), _key) //加密报文
	_Params["encryptData"] = _encryptData
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_http_url := BaseUrl + "/memberReg"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}
