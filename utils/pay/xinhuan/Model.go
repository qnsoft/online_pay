package xinhuan

var (
	PayUrl     = "https://www.dh0102.com/oss-transaction/gateway" //基础接口地址
	Mid        = "000010000000001"                                //商户号 //测试账号 000010000000001
	EncryptId  = "000010000000001"
	PublicKey  = "96ea552bcd55253ba90bbcffcf81e654"
	AgencyType = "hxkjdh"
)

/*
信还接口content  通用的请求参数
*/
type Base_Model struct {
	//默认字符串
	EncryptId string `json:"encryptId"`
	//Api 版本号
	ApiVersion string `json:"apiVersion"`
	//时当前间戳
	TxnDate string `json:"txnDate"`
}

/*
信还接口result  通用的响应参数：
*/
type Rt_Base_Model struct {
	//响应码
	Code string `json:"code"`
	//返回内容 以下返回参数均为 data 里的内容
	Data string `json:"data"`
	//描述
	Message string `json:"message"`
}

//————————————————以下是业务实体对象————————————————————//

/*
信还2.2
*/
type Rt22_Model struct {
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
信还3.2
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
代付-信还3.3 结算接口
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
查询-信还3.4 结算状态查询接口 返回
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
余额查询-信还3.5 余额查询接口
*/
type Rt35_Model struct {
	//基础对象
	Rt_Base_Model
	//业务对象
	Content Rt35_Model_Content `json:"content"`
}

/*
返回签约状态
*/
type Rt35_Model_Content struct {
	//商户订单号
	MerOrderNumber string `json:"merOrderNumber"`
}

/*
 签约- 信还7.1. 支付卡开通发短信接口——请求响应对象
*/
type Rt71_Model struct {
	//基础对象
	Rt_Base_Model
	//业务对象
	Content Rt71_Model_Content `json:"content"`
}

/*
返回签约状态
*/
type Rt71_Model_Content struct {
	//状态码
	IsSign string `json:"isSign"`
	//商户订单号
	MerOrderNumber string `json:"merOrderNumber"`
}

/*
签约确认-信还8.1支付卡开通确认短信接口——请求响应对象
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
