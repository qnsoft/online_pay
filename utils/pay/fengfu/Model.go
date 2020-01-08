package fengfu

var (
	PayUrl     = "https://47.96.91.41/service/gateway" //基础接口地址 //测试地址 https://47.92.100.165/service/gateway
	MerchantId = "ff19000011"                          //商户号 //测试账号 ff19000001
	//	ChannelTypeDe = "ffkj"                                  //通道类型小额channelType：ffqr大额channelType：ffkj
	//	ChannelTypeXe = "ffqr"                                  //小额
	Key       = "bb9bdb31322b639b7d0e376170625bd8" //测试key 96ea552bcd55253ba90bbcffcf81e654
	PublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDd/CY4uZqtE2bE4SjGzVfYePvN
TSuLxcTGpkpm5fiN0sCK7J9FXlG3Vw6nGDMeWm/dX4fj8vY/PVxlf6JGJObTC5pZ
7/sQjbREh5jFmJlbdNgQrlyCLa2pUgfQ22023Lvk/ssWTn7/sL0oy+vswbSuM8Fo
sZnKeKN9uLrMeRKhlwIDAQAB
-----END PUBLIC KEY-----
`)
	PrivateKey = []byte(`
-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAI9wDyyXSQfojHHO
CLnW2tX9yr+PHuCiiU449tEcDnsRFVpv4STGFS/toyYuOiJJqhcAOSd42HzQo8JG
4lQ+RmfZyHVswF6WRfDdcBGaLOjrwEGOwwGm22QK8D6XBHZZfy772fIAboA4b/2d
7J0YpwUZOXVKRE9kPr+bX+XuPvaxAgMBAAECgYBdylNRv/aUAj5bZ/KSuefMBzBc
w5PYJM+NY3gYUbUySYBh60gXiKJv/X4xyAZhRzO96Hp20fEU5ERAE6OInFPiNcx4
1CgF/xHB0/bBdNZlc4wTDrH1dxcyVe0uiKJefz6lnnkZ3cjwgc9V+jESRH9bIgjV
kq5GQgn+y9XcyvDQQQJBANe1d+RwAlVZ1TycN5AJNTASg4xY5XxGyaAYRhes4lm3
JXHiJzm7kpmNY4Q83A8yiNTrtApIji1Z+TdIRJfHUekCQQCqOs8HHfUjKu+n68rq
hTr9IC5KW7MyzXdfqEpIB73zQ+OiPdoiwhB5VAAACk3I6832s+xNA3AD4YOiL+qY
1XmJAkEA1i9YNrmJd5pVg3NSMHEUIQGhIVB7vaTEO3Ue1A9USeTzB1uZu1emv1WH
BR0xHN3+w/yurq9QjXOTY/McOOEvUQJBAJ60bUdJWr4KzZxUbL3wsGpFkmKo04zN
mjmIgJhe/2zJInSkDbbSDuHOuqFO+e4USdVzMR8r0UcM/Ng8rA8JF2ECQA92LAP0
fNc8fRVguf/2M2p1rC8VWi2oloxZR1lSiZAxnDpHHVtrwP0dn+oM0ZnxBKjVsyLZ
bhwP9kIRdiTEGL4=
-----END PRIVATE KEY-----
`)
)

/*
丰付接口请求返回基础对象
*/
type Rt_Base_Model struct {
	//状态码
	Code string `json:"code"`
	//状态信息
	Message string `json:"message"`
}

/*
丰付3.1
*/
type Rt31_Model struct {
	//基础对象
	Rt_Base_Model
	//业务对象
	Content Rt31_Model_Content `json:"content"`
}

/*
返回签约状态
*/
type Rt31_Model_Content struct {
	//商户订单号
	MerOrderNumber string `json:"merOrderNumber"`
	Tn             string `json:"tn"`
}

/*
丰付3.2
*/
type Rt32_Model struct {
	//基础对象
	Rt_Base_Model
	//业务对象
	Content Rt32_Model_Content `json:"content"`
}

/*
返回签约状态
*/
type Rt32_Model_Content struct {
	//商户订单号
	MerOrderNumber string `json:"merOrderNumber"`
	//描述
	DataDescribe string `json:"dataDescribe"`
	//订单状态
	OrderStatus string `json:"orderStatus"`
}

/*
代付-丰付3.3 结算接口
*/
type Rt33_Model struct {
	//基础对象
	Rt_Base_Model
	//业务对象
	Content Rt33_Model_Content `json:"content"`
}

/*
返回签约状态
*/
type Rt33_Model_Content struct {
	//商户订单号
	MerOrderNumber string `json:"merOrderNumber"`
	OrderStatus    string `json:"orderStatus"`
}

/*
查询-丰付3.4 结算状态查询接口 返回
*/
type Rt34_Model struct {
	//基础对象
	Rt_Base_Model
	//业务对象
	Content Rt34_Model_Content `json:"content"`
}

/*
返回签约状态
*/
type Rt34_Model_Content struct {
	//商户订单号
	MerOrderNumber string `json:"merOrderNumber"`
	OrderStatus    string `json:"orderStatus"`
}

/*
余额查询-丰付3.5 余额查询接口
*/
type Rt35_Model struct {
	//基础对象
	Rt_Base_Model
	//业务对象
	Content string `json:"content"`
}

/*
 签约- 丰付3.6. 支付卡开通发短信接口——请求响应对象
*/
type Rt36_Model struct {
	//基础对象
	Rt_Base_Model
	//业务对象
	Content Rt36_Model_Content `json:"content"`
}

/*
返回签约状态
*/
type Rt36_Model_Content struct {
	//状态码
	IsSign string `json:"isSign"`
	//商户订单号
	MerOrderNumber string `json:"merOrderNumber"`
}

/*
签约确认-丰付3.7支付卡开通确认短信接口——请求响应对象
*/
type Rt37_Model struct {
	//基础对象
	Rt_Base_Model
	//业务对象
	Content Rt37_Model_Content `json:"content"`
}

/*
返回确认短信状态
*/
type Rt37_Model_Content struct {
	//商户订单号
	MerOrderNumber string `json:"merOrderNumber"`
}
