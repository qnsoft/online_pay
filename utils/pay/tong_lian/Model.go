package TongLian

import (
	"sort"
	"strconv"
	"time"
	"qnsoft/web_api/utils/DateHelper"
	"qnsoft/web_api/utils/php2go"
)

var (
	//基础接口地址
	BaseUrl string = "https://test.allinpaygd.com/ipayapiweb"
)

//公共请求参数
type Model_public struct {
	//平台分配的机构号
	Orgid string `json:"orgid"`
	//平台分配的机构的APPID
	Appid string `json:"appid"`
	//接口版本号
	Version string `json:"version"`
	//商户自行生成的随机字符串
	Randomstr string `json:"randomstr"`
	//请求IP
	Reqip string `json:"reqip"`
	//请求时间戳
	Reqtime string `json:"reqtime"`
	//签名
	//Sign string `json:"sign"`
}

//公共响应参数
type Return_public struct {
	//返回码
	Retcode string `json:"retcode"`
	//返回码说明
	Retmsg string `json:"retmsg"`
	//商户自行生成的随机字符串
	Randomstr string `json:"randomstr"`
	//平台分配的机构的APPID
	Appid string `json:"appid"`
	//签名
	Sign string `json:"sign"`
}

//商户进件
type Model_22 struct {
	Model_public
	//拓展代理商号
	Belongorgid string `json:"belongorgid"`
	//商户外部唯一标记
	Outcusid string `json:"outcusid"`
	//商户名称
	Cusname string `json:"cusname"`
	//商户简称
	Cusshortname string `json:"cusshortname"`
	//所在省
	Merprovice string `json:"merprovice"`
	//所在市
	Areacode string `json:"areacode"`
	//法人姓名
	Legal string `json:"legal"`
	//法人代表证件号
	Idno string `json:"idno"`
	//法人手机号码
	Phone string `json:"phone"`
	//注册地址
	Address string `json:"address"`
	//账户号
	Acctid string `json:"acctid"`
	//账户名
	Acctname string `json:"acctname"`
	//卡折类型(00-借记卡,01-存折)
	Accttp string `json:"accttp"`
	//拓展人
	Expanduser string `json:"expanduser"`
	//支付产品信息列表(产品列表的json)
	Prodlist []Prod_Model `json:"prodlist"`
	//提现手续费
	Settfee string `json:"settfee"`
}

//产品信息参数
type Prod_Model struct {
	//交易类型
	Trxcode string `json:"trxcode"`
	//费率
	Feerate string `json:"feerate"`
}

//商户进件返回
type Return_22 struct {
	Return_public
	//平台生成商户号
	Cusid string `json:"cusid"`
	//外部商户标识
	Outcusid string `json:"outcusid"`
}

//商户进件状态查询
type Model_23 struct {
	Return_public
	//商户外部唯一标识
	Outcusid string `json:"outcusid"`
}

//商户进件状态查询返回
type Return_23 struct {
	Return_public
	//商户状态
	State string `json:"state"`
	//平台商户号
	Cusid string `json:"cusid"`
	//商户外部唯一标识
	Outcusid string `json:"outcusid"`
}

//商户结算/费率信息修改
type Model_24 struct {
	Return_public
	//商户号
	Cusid string `json:"cusid"`
	//账户号
	Acctid string `json:"acctid"`
	//卡折类型
	Accttp string `json:"accttp"`
	//支付产品信息列表
	Prodlist string `json:"prodlist"`
	//实时到账手续费
	Settfee string `json:"settfee"`
}

//商户结算/费率信息修改返回
type Return_24 struct {
	Return_public
}

//商户绑定消费银行卡
type Model_25 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//商户用户号
	Meruserid string `json:"meruserid"`
	//银行卡
	Cardno string `json:"cardno"`
	//账户名
	Acctname string `json:"acctname"`
	//卡类型(00:借记卡02:信用卡)
	Accttype string `json:"accttype"`
	//有效期(MMyy)
	Validdate string `json:"validdate"`
	//Cvv2
	Cvv2 string `json:"cvv2"`
	//身份证号
	Idno string `json:"idno"`
	//预留手机号码
	Tel string `json:"tel"`
}

//商户绑定消费银行卡返回对象
type Return_25 struct {
	Return_public
	//交易的状态,0000:绑卡成功,流程完成其余详看说明,0000:绑卡成功,流程完成,3051: 协议已存在,请勿重复签约,1999: 短信验证码已发送,请继续调用签约确认接口完成绑卡操作,其他,交易失败
	Trxstatus string `json:"trxstatus"`
	//绑卡完成返回协议编号
	Agreeid string `json:"agreeid"`
	//通联用户号
	Memid string `json:"memid"`
	//交易透传信息
	Thpinfo string `json:"thpinfo"`
}

//重新获取绑卡验证码
type Model_26 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//商户用户号
	Meruserid string `json:"meruserid"`
	//银行卡
	Cardno string `json:"cardno"`
	//账户名
	Acctname string `json:"acctname"`
	//卡类型(00:借记卡02:信用卡)
	Accttype string `json:"accttype"`
	//有效期(MMyy)
	Validdate string `json:"validdate"`
	//Cvv2
	Cvv2 string `json:"cvv2"`
	//身份证号
	Idno string `json:"idno"`
	//预留手机号码
	Tel string `json:"tel"`
	//交易透传信息(签约申请时返回,如果不为空,则原样带上)
	Thpinfo string `json:"thpinfo"`
}

//重新获取绑卡验证码返回对象
type Return_26 struct {
	Return_public
}

//商户绑定银行卡确认
type Model_27 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//商户用户号
	Meruserid string `json:"meruserid"`
	//银行卡
	Cardno string `json:"cardno"`
	//账户名
	Acctname string `json:"acctname"`
	//卡类型(00:借记卡02:信用卡)
	Accttype string `json:"accttype"`
	//有效期(MMyy)
	Validdate string `json:"validdate"`
	//Cvv2
	Cvv2 string `json:"cvv2"`
	//身份证号
	Idno string `json:"idno"`
	//预留手机号码
	Tel string `json:"tel"`
	//短信验证码
	Smscode string `json:"smscode"`
	//交易透传信息(签约申请时返回,如果不为空,则原样带上)
	Thpinfo string `json:"thpinfo"`
}

//商户绑定银行卡确认返回对象
type Return_27 struct {
	Return_public
	//交易状态(0000:绑卡成功,流程完成其余详看说明)
	Trxstatus string `json:"trxstatus"`
	//绑卡完成返回协议编号
	Agreeid string `json:"agreeid"`
	//通联用户号
	Memid string `json:"memid"`
}

//商户解绑消费银行卡
type Model_28 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//银行卡
	Cardno string `json:"cardno"`
}

//商户解绑消费银行卡返回对象
type Return_28 struct {
	Return_public
}

//快捷交易支付申请
type Model_29 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//商户订单号
	Orderid string `json:"orderid"`
	//协议编号
	Agreeid string `json:"agreeid"`
	//交易类型(QUICKPAY_OF_HP QUICKPAY_OF_NP QUICKPAY_OL_HP)
	Trxcode string `json:"trxcode"`
	//订单金额
	Amount string `json:"amount"`
	//手续费
	Fee string `json:"fee"`
	//币种
	Currency string `json:"currency"`
	//订单内容
	Subject string `json:"subject"`
	//有效时间
	Validtime string `json:"validtime"`
	//市别
	City string `json:"city"`
	//行业
	Mccid string `json:"mccid"`
	//交易备注
	Trxreserve string `json:"trxreserve"`
	//交易结果通知地址
	Notifyurl string `json:"notifyurl"`
}

//快捷交易支付申请返回对象
type Return_29 struct {
	Return_public
	//商户订单号
	Orderid string `json:"orderid"`
	//交易状态
	Trxstatus string `json:"trxstatus"`
	//错误原因
	Errmsg string `json:"errmsg"`
	//交易单号
	Trxid string `json:"trxid"`
	//交易完成时间
	Fintime string `json:"fintime"`
	//交易透传信息
	Thpinfo string `json:"thpinfo"`
}

//快捷交易支付确认
type Model_210 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//交易单号(平台的交易流水号)
	Trxid string `json:"trxid"`
	//协议编号
	Agreeid string `json:"agreeid"`
	//短信验证码
	Smscode string `json:"smscode"`
	//交易透传信息(支付申请或者错误码为1999返回的thpinfo原样带上)
	Thpinfo string `json:"thpinfo"`
}

//快捷交易支付确认返回对象
type Return_210 struct {
	Return_public
	//交易状态
	Trxstatus string `json:"trxstatus"`
	//错误原因
	Errmsg string `json:"errmsg"`
	//交易单号
	Trxid string `json:"trxid"`
	//交易完成时间
	Fintime string `json:"fintime"`
	//交易透传信息
	Thpinfo string `json:"thpinfo"`
}

//快捷支付短信重新获取
type Model_211 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//交易单号(平台的交易流水号)
	Trxid string `json:"trxid"`
	//协议编号
	Agreeid string `json:"agreeid"`
	//交易透传信息(支付申请或者错误码为1999返回的thpinfo原样带上)
	Thpinfo string `json:"thpinfo"`
}

//快捷支付短信重新获取返回对象
type Return_211 struct {
	Return_public
	//交易透传信息
	Thpinfo string `json:"thpinfo"`
}

//快捷小额免短信支付
type Model_212 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//商户订单号
	Orderid string `json:"orderid"`
	//协议编号
	Agreeid string `json:"agreeid"`
	//交易类型(QUICKPAY_OF_HP QUICKPAY_OF_NP QUICKPAY_OL_HP)
	Trxcode string `json:"trxcode"`
	//订单金额
	Amount string `json:"amount"`
	//手续费
	Fee string `json:"fee"`
	//币种
	Currency string `json:"currency"`
	//订单内容
	Subject string `json:"subject"`
	//有效时间
	Validtime string `json:"validtime"`
	//市别
	City string `json:"city"`
	//行业
	Mccid string `json:"mccid"`
	//交易备注
	Trxreserve string `json:"trxreserve"`
	//交易结果通知地址
	Notifyurl string `json:"notifyurl"`
}

//快捷小额免短信支付返回对象
type Return_212 struct {
	Return_public
	//商户订单号
	Orderid string `json:"orderid"`
	//交易状态
	Trxstatus string `json:"trxstatus"`
	//错误原因
	Errmsg string `json:"errmsg"`
	//交易单号
	Trxid string `json:"trxid"`
	//交易完成时间
	Fintime string `json:"fintime"`
	//交易透传信息
	Thpinfo string `json:"thpinfo"`
}

//快捷交易查询
type Model_213 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//商户订单号
	Orderid string `json:"orderid"`

	//平台交易流水
	Trxid string `json:"trxid"`
	//交易日期
	Date string `json:"date"`
}

//快捷交易查询返回对象
type Return_213 struct {
	Return_public
	//交易单号
	Trxid string `json:"trxid"`
	//商户订单号
	Orderid string `json:"orderid"`
	//交易类型(详见附录3.1)
	Trxcode string `json:"trxcode"`
	//交易金额(单位为分)
	Trxamt string `json:"trxamt"`
	//交易状态(详见附录3.3)
	Trxstatus string `json:"trxstatus"`
	//交易完成时间
	Fintime string `json:"fintime"`
	//错误原因
	Errmsg string `json:"errmsg"`
}

//余额查询
type Model_214 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
}

//余额查询返回对象
type Return_214 struct {
	Return_public
	//余额
	Balance string `json:"balance"`
}

//提现
type Model_215 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//商户订单号
	Orderid string `json:"orderid"`
	//订单金额
	Amount string `json:"amount"`
	//全额提现(1-代表全额提取)
	Isall string `json:"isall"`
	//手续费
	Fee string `json:"fee"`

	//交易备注
	Trxreserve string `json:"trxreserve"`
	//交易结果通知地址
	Notifyurl string `json:"notifyurl"`
}

//提现返回对象
type Return_215 struct {
	Return_public
	//商户订单号
	Orderid string `json:"orderid"`
	//交易状态
	Trxstatus string `json:"trxstatus"`
	//错误原因
	Errmsg string `json:"errmsg"`
	//交易单号
	Trxid string `json:"trxid"`
	//交易完成时间(yyyyMMddHHmmss)
	Fintime string `json:"fintime"`
	//到账帐号
	Acctno string `json:"acctno"`
	//提现金额
	Amount string `json:"amount"`
	//到账金额
	Actualamount string `json:"actualamount"`
	//手续费
	Fee string `json:"fee"`
}

//付款
type Model_216 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//商户订单号
	Orderid string `json:"orderid"`
	//订单金额
	Amount string `json:"amount"`
	//全额提现(1-代表全额提取)
	Isall string `json:"isall"`
	//手续费
	Fee string `json:"fee"`
	//协议编号
	Agreeid string `json:"agreeid"`
	//交易备注
	Trxreserve string `json:"trxreserve"`
	//交易结果通知地址
	Notifyurl string `json:"notifyurl"`
}

//付款返回对象
type Return_216 struct {
	Return_public
	//商户订单号
	Orderid string `json:"orderid"`
	//交易状态
	Trxstatus string `json:"trxstatus"`
	//错误原因
	Errmsg string `json:"errmsg"`
	//交易单号
	Trxid string `json:"trxid"`
	//交易完成时间(yyyyMMddHHmmss)
	Fintime string `json:"fintime"`
	//到账帐号
	Acctno string `json:"acctno"`
	//提现金额
	Amount string `json:"amount"`
	//到账金额
	Actualamount string `json:"actualamount"`
	//手续费
	Fee string `json:"fee"`
}

//提现(付款)交易查询
type Model_217 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//商户订单号
	Orderid string `json:"orderid"`
	//平台交易流水号
	Trxid string `json:"trxid"`
	//交易日期
	Date string `json:"date"`
}

//提现(付款)交易查询返回对象
type Return_217 struct {
	Return_public
	//平台交易流水号
	Trxid string `json:"trxid"`
	//商户订单号
	Orderid string `json:"orderid"`
	//交易类型
	Trxcode string `json:"trxcode"`
	//到账帐号
	Acctno string `json:"acctno"`
	//提现金额
	Amount string `json:"amount"`
	//提现金额
	Actualamount string `json:"actualamount"`
	//手续费
	Fee string `json:"fee"`
	//交易状态
	Trxstatus string `json:"trxstatus"`
	//交易备注
	Trxreserve string `json:"trxreserve"`
	//交易完成时间(yyyyMMddHHmmss)
	Fintime string `json:"fintime"`
	//错误原因
	Errmsg string `json:"errmsg"`
}

/* 219 是异步回传，且参数是变参 所以不定义实体 */

//对账文件下载
type Model_219 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//交易日期
	Trxdate string `json:"trxdate"`
}

//对账文件下载返回对象
type Return_219 struct {
	Return_public
	//对账文件地址
	Reporturl string `json:"reporturl"`
}

//查询绑卡协议号
type Model_220 struct {
	Model_public
	//商户号
	Cusid string `json:"cusid"`
	//卡号
	Cardno string `json:"cardno"`
}

//查询绑卡协议号返回对象
type Return_220 struct {
	Return_public
	//商户号
	Cusid string `json:"cusid"`
	//卡号
	Cardno string `json:"cardno"`
	//协议号
	Agreeid string `json:"agreeid"`
}

//查询地区对应可用行业分类代码
type Model_221 struct {
	Model_public
	//地区代码
	Areacode string `json:"areacode"`
	//交易类型
	Trxcode string `json:"trxcode"`
}

//查询地区对应可用行业分类代码返回对象
type Return_221 struct {
	Return_public
	//行业列表
	Mcclist []AreaMccUnit_Model `json:"mcclist"`
}

//地区列表实体
type AreaMccUnit_Model struct {
	//地区代码
	Areacode string `json:"areacode"`
	//行业ID
	Mccid string `json:"mccid"`
	//行业名称
	Mccname string `json:"mccname"`
}

/*
获取基础参数
*/
func GetBaseParam() map[string]interface{} {
	_rt := make(map[string]interface{})
	_rt["orgid"] = "201000003530"
	_rt["appid"] = "6666678"
	_rt["key"] = "549935538f05ceaef400cbdb2dda4ebf"
	_rt["version"] = "11"
	_rt["randomstr"] = date.FormatDate(time.Now(), "yyMMddHHmmss") + strconv.Itoa(php2go.Rand(100, 999))
	// _rt["reqip"] = "127.0.0.1"
	// _rt["reqtime"] = date.FormatDate(time.Now(), "yyMMddHHmmss")
	return _rt
}

/*
参数集排序Sign签名 根据key排序
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
				_str += v + php2go.URLEncode(_srt)
			} else {
				_str += v + _srt
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
