package GongYiFuNew

import (
	"strconv"
	"time"
)

///------------------------------------基础参数-------------------------------------------------///
var (
	Merno     string = "MCH104702"                        //接入商户编号
	Key       string = "db7770445fdde3a663f6b93d7a08382a" //商户蜜钥KEY
	Orgno     string = "PT1047"                           //接入机构编号
	Version   string = "1.0"                              //版本号
	PayUrl    string = "https://quick.weiyifu88.com/api/" //接口url
	PublicKey        = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDnpRd6Rp23Gu6tUzdZ0eSEgmJ+
qySzg75b3A7y+63IogoPJu6R0Sh/vne9yn+8rrOmQuPLNdMWo4M+ntGpf1HAdARU
bqjNjWKJDt6d3tia+xTPlBuzrja6RjbpM1OJ86/qoHOZFaOi7zJgeRG8/WGJIEv3
u3lwjCP8QM5GvrDnWwIDAQAB
-----END PUBLIC KEY-----
`)
	PrivateKey = []byte(`
-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALUy+MNSXWGOgZnH
l8V9rye1z4nRs4HWFDmvWWTj3VISgLeVLls+xpL4MajDzK+z0LH9nqf8vuoIKwUw
hifSsCqypm5/APHF8XI1KH7K2+l+IYy6SVaoUCiJyrUKBcUcyOgJvJYF1SVw7xJs
Jo63X/6dk6clSebsWZBfR69Eyw5FAgMBAAECgYBnunc38VWtvFOqwdy7XLjBZc4a
GmbFg9TuNNha7irLifYPoiH4cBZjGhvrfbMWPjzRN9v0VLbB6M0f2mhiMbVQtLLg
KXrJbXpVnBKxzutEiZ34f/md30fL1GtWavZj0PuY2Y/8F6y13CcLPkLLvk6j7o9h
0xMhYso5D08kmPPkIQJBAOltpO2bFIUKeDqHaNrbCjbanlQbWZQ+gPGyOzDcbS4f
WWWCQd6o7yeylADk4852xO3opAHQm/31peRKUy10THkCQQDGuG4aND5u3FMEXUmY
N7XKuRF//g+T6u4dQaoHCaZFa9bQOHYfgtxoKxBPzFYUFMdQns1hHIennK6pToUb
ZUUtAkA4Vxh5qPaY7d/68Hfkav3aI4YXcsp6N2PT8lrK/kjz2uku0POpFEk04atL
U/OP/6akbYQ4U+tyrnmt0iqlS+6xAkEApFqHHX8WH+RzeMmbA50X6smz0pMS2TjV
pTbY5Ccz8HinWuFHuPonRrRPMmCC1Or2ihQ9MtNA0vzAbGD3r9fLJQJAWvbu22Dl
ThJuy1D8/NxwWdBrIqntU4jjBM5Ad9fM8bIXb8bOjU4ENHlInbUP6XmOPXRa08JK
8KS67pdSBrRWng==
-----END PRIVATE KEY-----
`)
)

///------------------------------------公用参数-------------------------------------------------///
/*
获取基础参数
*/
func GetBaseParam() map[string]interface{} {
	_rt := make(map[string]interface{})
	_rt["merchantNo"] = "100101"                            //商户编号
	_rt["msgId"] = strconv.FormatInt(time.Now().Unix(), 10) //
	return _rt
}

///------------------------------------商户进件提交接口对象-------------------------------------------------///
type Model_jinjian struct {
	BankAccountNo          string `json:"bankAccountNo"`          //结算卡号
	CertificateNo          string `json:"certificateNo"`          //身份证号码
	Mobile                 string `json:"mobile"`                 //子商户银行卡预留手机号
	Remark                 string `json:"remark"`                 //备注
	SubMerchantFeeRate     string `json:"subMerchantFeeRate"`     //子商户交易费率,单位万1
	SubMerchantName        string `json:"subMerchantName"`        //子商户姓名
	SubMerchantWithdrawFee string `json:"subMerchantWithdrawFee"` //子商户提现手续费,单位分
}

/*
商户进件提交返回对象
*/
type Model_jinjian_Return struct {
	Code string                    `json:"code"` //返回CODE
	Msg  string                    `json:"msg"`  //返回说明
	Data Model_jinjian_Return_data `json:"data"` //子商户对象
}
type Model_jinjian_Return_data struct {
	MerchantNo    string `json:"merchantNo"`    //商户号
	SubMerchantNo string `json:"subMerchantNo"` //子商户号
}

///------------------------------------S102.商户信息修改接口对象-------------------------------------------------///
type Model_102 struct {
	SubMerchId string `json:"subMerchId"` //商户编号
	Phone      string `json:"phone"`      //法人电话
	CardId     string `json:"cardId"`     //绑定结算卡号
}

/*
S102商户信息修改返回对象
*/
type Model_Return_102 struct {
	Code string `json:"code"` //返回CODE
	Msg  string `json:"msg"`  //返回说明
}

///------------------------------------S103.商户信息费率接口对象-------------------------------------------------///
type Model_103 struct {
	SubMerchId string `json:"subMerchId"` //商户编号
	FeeRate    string `json:"feeRate"`    //交易费率
	ExternFee  string `json:"externFee"`  //附加手续费
}

/*
S103.商户信息费率返回对象
*/
type Model_Return_103 struct {
	Code string `json:"code"` //返回CODE
	Msg  string `json:"msg"`  //返回说明
}

///------------------------------------S201.银联页面绑卡接口对象-------------------------------------------------///
type Model_201 struct {
	SubMerchId string `json:"subMerchId"` //商户编号
	Name       string `json:"name"`       //法人姓名
	Phone      string `json:"phone"`      //法人电话
	IdNo       string `json:"idNo"`       //身份证号
	CardId     string `json:"cardId"`     //交易卡号
	NotifyUrl  string `json:"notifyUrl"`  //异步通知地址
	FrontUrl   string `json:"frontUrl"`   //页面通知地址
	OrderId    string `json:"orderId "`   //请求流水号
	DeviceId   string `json:"deviceId "`  //设备ID 安卓:IMEI，iOS:IDFV，PC:硬盘序列号（若不填大额交易限额会被银联风控）
	IpAddress  string `json:"ipAddress "` //请求IP地址 公网IP地址（若不填大额交易限额会被风控）（付款客户端IP）
}

/*
S201.银联页面绑卡返回对象
*/
type Model_Return_201 struct {
	Code string                  `json:"code"` //返回CODE
	Msg  string                  `json:"msg"`  //返回说明
	Data Model_Return_201_RtHtml `json:"data"` //返回data
}

/*
S201.银联页面绑卡返回html对象
*/
type Model_Return_201_RtHtml struct {
	Html string `json:"html"` //银联入网页面
}

///------------------------------------S202.绑卡查询接口对象-------------------------------------------------///
type Model_202 struct {
	SubMerchId string `json:"subMerchId"` //商户编号
	CardId     string `json:"cardId"`     //需查询的卡号
}

/*
S202.绑卡查询返回对象
*/
type Model_Return_202 struct {
	Code string                `json:"code"` //返回CODE
	Msg  string                `json:"msg"`  //返回说明
	Data Model_Return_202_Data `json:"data"` //返回data
}
type Model_Return_202_Data struct {
	CardNo string `json:"cardNo"` //查询卡号
	Status string `json:"status"` //状态
}

///------------------------------------S301.订单查询接口对象-------------------------------------------------///
/*
S301订单查询对象
*/
type Model_301 struct {
	OrderId    string `json:"orderId"`    //订单号
	SubMerchId string `json:"subMerchId"` //子商户号非必填
}

/*
S301订单查询返回对象
*/
type Model_Return_301 struct {
	Code string         `json:"code"`
	Data Model_Data_301 `json:"data"`
	Msg  string         `json:"msg"`
}

/*
S301订单查询返回对象data
*/
type Model_Data_301 struct {
	Amount       int    `json:"amount"`
	RespDesc     string `json:"respDesc"`
	OrderID      string `json:"orderId"`
	SettleAmount int    `json:"settleAmount"`
	OrderStatus  string `json:"orderStatus"`
	RespCode     string `json:"respCode"`
}

///------------------------------------S302.快捷支付接口对象-------------------------------------------------///
type Model_302 struct {
	SubMerchId string `json:"subMerchId"` //子商户号
	OrderId    string `json:"orderId"`    //订单号
	Name       string `json:"name"`       //法人姓名
	Phone      string `json:"phone"`      //法人电话
	IdNo       string `json:"idNo"`       //身份证号
	CardId     string `json:"cardId"`     //交易卡号
	NotifyUrl  string `json:"notifyUrl"`  //异步通知地址
	Amount     string `json:"amount"`     //交易金额
	GoodsName  string `json:"goodsName"`  //订单名称
	CardType   string `json:"cardType"`   //卡类型
	Cvv        string `json:"cvv"`        //安全码
	ExpDate    string `json:"expDate"`    //有效期
	DeviceId   string `json:"deviceId "`  //设备ID 安卓:IMEI，iOS:IDFV，PC:硬盘序列号（若不填大额交易限额会被银联风控）
	IpAddress  string `json:"ipAddress "` //请求IP地址 公网IP地址（若不填大额交易限额会被风控）（付款客户端IP）
}

/*
S302.快捷支付返回对象
*/
type Model_Return_302 struct {
	Code string                `json:"code"` //返回CODE
	Msg  string                `json:"msg"`  //返回说明
	Data Model_Return_302_Data `json:"data"` //返回数据
}
type Model_Return_302_Data struct {
	OrderId     string `json:"orderId"`     //订单号
	Amount      string `json:"amount"`      //交易金额
	OrderStatus string `json:"orderStatus"` //订单状态 02 交易成功 03 失败 其他处理中
	RespCode    string `json:"respCode"`    //响应码
	RespDesc    string `json:"respDesc"`    //响应描述
}

///------------------------------------S302.落地云闪付支付接口对象-------------------------------------------------///
type Model_302_A struct {
	SubMerchId string `json:"subMerchId"` //子商户号
	OrderId    string `json:"orderId"`    //订单号
	Name       string `json:"name"`       //法人姓名
	Phone      string `json:"phone"`      //法人电话
	IdNo       string `json:"idNo"`       //身份证号
	CardId     string `json:"cardId"`     //交易卡号
	NotifyUrl  string `json:"notifyUrl"`  //异步通知地址
	Amount     string `json:"amount"`     //交易金额
	GoodsName  string `json:"goodsName"`  //订单名称
	CardType   string `json:"cardType"`   //卡类型
	Cvv        string `json:"cvv"`        //安全码
	ExpDate    string `json:"expDate"`    //有效期
	DeviceId   string `json:"deviceId "`  //设备ID 安卓:IMEI，iOS:IDFV，PC:硬盘序列号（若不填大额交易限额会被银联风控）
	IpAddress  string `json:"ipAddress "` //请求IP地址 公网IP地址（若不填大额交易限额会被风控）（付款客户端IP）
	CityCode   string `json:"cityCode"`   //商户地区编码
}

/*
SS302.落地云闪付支付返回对象
*/
type Model_Return_302_A struct {
	Code string                  `json:"code"` //返回CODE 0000 请求成功 0100 处理中 其他失败
	Msg  string                  `json:"msg"`  //返回说明
	Data Model_Return_302_A_Data `json:"data"` //返回数据

}
type Model_Return_302_A_Data struct {
	OrderId     string `json:"orderId"`     //订单号
	Amount      string `json:"amount"`      //交易金额
	OrderStatus string `json:"orderStatus"` //订单状态 02 交易成功 03 失败 其他处理中
	RespCode    string `json:"respCode"`    //响应码
	RespDesc    string `json:"respDesc"`    //响应描述
}

///------------------------------------S302.子商户结算接口对象-------------------------------------------------///
type Model_302_B struct {
	SubMerchId string `json:"subMerchId"` //子商户号
	OrderId    string `json:"orderId"`    //订单号
	Name       string `json:"name"`       //法人姓名
	Phone      string `json:"phone"`      //法人电话
	IdNo       string `json:"idNo"`       //身份证号
	CardId     string `json:"cardId"`     //交易卡号
	NotifyUrl  string `json:"notifyUrl"`  //异步通知地址
	Amount     string `json:"amount"`     //交易金额
	Remark     string `json:"remark"`     //备注
}

/*
S302.子商户结算返回对象
*/
type Model_Return_302_B struct {
	Code string                  `json:"code"` //返回CODE 0000 请求成功 0100 处理中 其他失败
	Msg  string                  `json:"msg"`  //返回说明
	Data Model_Return_302_B_Data `json:"data"` //返回数据
}
type Model_Return_302_B_Data struct {
	OrderId     string  `json:"orderId"`     //订单号
	Amount      float64 `json:"amount"`      //交易金额
	OrderStatus string  `json:"orderStatus"` //订单状态 02 交易成功 03 失败 其他处理中
	RespCode    string  `json:"respCode"`    //响应码
	RespDesc    string  `json:"respDesc"`    //响应描述
}

///------------------------------------S302.子商户余额查询接口对象-------------------------------------------------///
type Model_302_C struct {
	SubMerchId string `json:"subMerchId"` //子商户号
}

/*
S302.子商户余额查询返回对象
*/
type Model_Return_302_C struct {
	Code string                  `json:"code"` //返回CODE 0000 请求成功 0100 处理中 其他失败
	Msg  string                  `json:"msg"`  //返回说明
	Data Model_Return_302_C_Data `json:"data"` //返回数据

}
type Model_Return_302_C_Data struct {
	/*
		"data":{"acctId":"SMCH999902338719801","acctType":"00","availBalance":702,"frozenBalace":0,"insertTime":1556358269000,"updateTime":1559618545000}
	*/
	AcctId       string `json:"acctId"`       //
	AcctType     string `json:"acctType"`     //
	AvailBalance int    `json:"availBalance"` //可用余额
	FrozenBalace int    `json:"frozenBalace"` //冻结余额
	InsertTime   int    `json:"insertTime"`   //
	UpdateTime   int    `json:"updateTime"`   //
}

///------------------------------------S401.绑卡异步通知接口对象-------------------------------------------------///
type Model_401 struct {
}

///------------------------------------S402.交易异步通知接口对象-------------------------------------------------///
type Model_402 struct {
}
