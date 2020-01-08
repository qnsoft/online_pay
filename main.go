package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"

	_ "github.com/qnsoft/online_pay/routers"

	"github.com/astaxie/beego"
	"github.com/wenzhenxi/gorsa"
)

func main() {
	//----------------------------工易付S101商户信息提交--------------------
	// _model_101 := GongYiFu.Model_101{
	// 	MerchName:    "河南联九舟网络科技有限公司",
	// 	Name:         "宋晓辉",
	// 	Phone:        "13938202388",
	// 	IdNo:         "410185198209190514",
	// 	MerchAddress: "郑州市郑汴路环球大厦B座2303",
	// 	CardId:       "6214623121001278490",
	// 	FeeRate:      "45",  //费率0.68%
	// 	ExternFee:    "100", //手续费单位分
	// 	Remark:       "",
	// }
	// _model_rt_101 := GongYiFu.Gyf_S101(_model_101)
	// _str_json_101, _ := json.Marshal(_model_rt_101)
	// fmt.Print("%s", string(_str_json_101))
	// // ----------------------------工易付S102商户信息提交--------------------
	// _model_102 := GongYiFu.Model_102{
	// 	SubMerchId: "SMCH10470212909113",
	// 	Phone:      "15638905677",
	// 	CardId:     "6226223001466633",
	// }
	// _model_rt_102 := GongYiFu.Gyf_S102(_model_102)
	// _str_json_102, _ := json.Marshal(_model_rt_102)
	// fmt.Print("%s", string(_str_json_102))
	// // ----------------------------工易付S103.商户修改费率--------------------
	// _model_103 := GongYiFu.Model_103{
	// 	SubMerchId: "SMCH10470212909113",
	// 	FeeRate:    "58",
	// 	ExternFee:  "100",
	// }
	// _model_rt_103 := GongYiFu.Gyf_S103(_model_103)
	// _str_json_103, _ := json.Marshal(_model_rt_103)
	// fmt.Print("%s", string(_str_json_103))
	// ----------------------------工易付S201.银联绑卡--------------------
	// _OrderId := date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4)
	// _model_201 := GongYiFu.Model_201{
	// 	SubMerchId: "SMCH10470212929804",                                         //商户编号
	// 	Name:       "宋晓辉",                                                        //法人姓名
	// 	Phone:      "13938202388",                                                //法人电话
	// 	IdNo:       "410185198209190514",                                         //身份证号
	// 	CardId:     "6214623121001278490",                                        //交易卡号
	// 	NotifyUrl:  "http://cgi.weiyifu123.com/cgi-bin/nps/cps_pay_callback.cgi", //异步通知地址
	// 	FrontUrl:   "www.baidu.com",                                              //页面通知地址
	// 	OrderId:    _OrderId,                                                     //请求流水号
	// 	DeviceId:   "358588070801255",                                            //设备ID 安卓:IMEI，iOS:IDFV，PC:硬盘序列号（若不填大额交易限额会被银联风控）
	// 	IpAddress:  "39.97.111.217",                                              //请求IP地址 公网IP地址（若不填大额交易限额会被风控）（付款客户端IP）
	// }
	// _model_rt_201 := GongYiFu.Gyf_S201(_model_201)
	// _str_json_201, _ := json.Marshal(_model_rt_201)
	// fmt.Print("%s", string(_str_json_201))
	//--------------------------S301订单查询接口-------------------
	// _model_301 := GongYiFu.Model_301{
	// 	OrderId: "1909290916278874",
	// 	//SubMerchId: "SMCH99990212558313",
	// }
	// str_301 := GongYiFu.Gyf_S301(_model_301)

	// // 	if _b_ret != nil && len(_b_ret.Code) > 0 {
	// // 		if _b_ret.Code == "0000" {
	// // 			return map[string]interface{}{"status": 0, "msg": _b_ret.Msg, "ret_data": _b_ret.Data, "json": _b_json}
	// // 		}
	// // 		return map[string]interface{}{"status": 1, "msg": _b_ret.Msg, "ret_data": _b_ret.Data, "json": _b_json}
	// // 	}
	// // }
	// // return map[string]interface{}{"status": 1, "msg": "解密失败", "json": _json}

	// fmt.Print("%#v", str_301)
	//--------------------------S302.落地云闪付支付接口-------------------
	// _OrderId := date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4)
	// _model_302_A := GongYiFu.Model_302_A{
	// 	SubMerchId: "SMCH10470212929804",
	// 	OrderId:    _OrderId,
	// 	Name:       "宋晓辉",
	// 	Phone:      "13938202388",
	// 	IdNo:       "410185198209190514",
	// 	CardId:     "6225768311530149",
	// 	NotifyUrl:  "http://cgi.weiyifu123.com/cgi-bin/nps/cps_pay_callback.cgi",
	// 	Amount:     "5",
	// 	GoodsName:  "淘客商城订单",
	// 	CardType:   "2",
	// 	Cvv:        "999",
	// 	ExpDate:    "1120",
	// 	DeviceId:   "358588070801255", //设备ID 安卓:IMEI，iOS:IDFV，PC:硬盘序列号（若不填大额交易限额会被银联风控）
	// 	IpAddress:  "39.97.111.217",   //请求IP地址 公网IP地址（若不填大额交易限额会被风控）（付款客户端IP）
	// }
	// _model_rt_302_A := GongYiFu.Gyf_S302_A(_model_302_A)
	// _str_json_302_A, _ := json.Marshal(_model_rt_302_A)
	// fmt.Print("%s", string(_str_json_302_A))

	//----------------------------S501.地区编码查询--------------------
	//str_501 := GongYiFu.Gyf_S501()
	//fmt.Print("%+v", str_501)

	// beego.ErrorHandler("404", page_not_found)
	// beego.ErrorHandler("401", page_note_permission)

	beego.Run()
}

// func page_not_found(rw http.ResponseWriter, r *http.Request) {
// 	t, _ := template.New("404.html").ParseFiles("views/404.html")
// 	data := make(map[string]interface{})

// 	//data["content"] = "page not found"
// 	t.Execute(rw, data)
// }

// func page_note_permission(rw http.ResponseWriter, r *http.Request) {
// 	t, _ := template.New("401.html").ParseFiles("views/401.html")
// 	data := make(map[string]interface{})
// 	//data["content"] = "你没有权限访问此页面，请联系超级管理员。或去<a href='/'>开启我的OPMS之旅</a>。"
// 	t.Execute(rw, data)
// }

// /*
// 私钥加密
// */
// func Rsa_Encode(data, _privateKey []byte) (string, error) {
// 	_srt, errrt := gorsa.PriKeyEncrypt(string(data), string(_privateKey))
// 	return _srt, errrt
// 	//return []byte(_srt), errrt

// }

// /*
// 公钥解密
// */
// func Rsa_Decode(encrypted, _publicKey []byte) ([]byte, error) {
// 	_srt, errrt := gorsa.PublicDecrypt(string(encrypted), string(_publicKey))
// 	return []byte(_srt), errrt
// }
/*
私钥签名
*/
func Rsa_Sign(data, _privateKey []byte) (string, error) {
	_str_content := string(data)
	_str_privateKey := string(_privateKey)
	fmt.Println("rsa原始内容：", _str_content)
	_srt, errrt := gorsa.SignSha256WithRsa(_str_content, _str_privateKey)
	return _srt, errrt
}

func Sign(data string, _privateKey []byte) (string, error) {
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	//获取私钥
	block, _ := pem.Decode(_privateKey)
	//解析PKCS1格式的私钥
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	signByte, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
	sign := base64.StdEncoding.EncodeToString(signByte)
	return sign, err
	// block, _ := pem.Decode(_privateKey)
	// privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	// sha256Hash := sha256.New()
	// s_data := []byte(data)
	// sha256Hash.Write(s_data)
	// hashed := sha256Hash.Sum(nil)

	// signByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	// sign := base64.StdEncoding.EncodeToString(signByte)
	// // decoded, _ := base64.StdEncoding.DecodeString(sign)
	// // fmt.Println(string(decoded))
	// return sign, err
}

func Base64UrlEncode(s string) string {
	s = strings.Split(s, "=")[0]
	s = strings.Replace(s, "+", "-", -1)
	s = strings.Replace(s, "/", "_", -1)
	return s
}
