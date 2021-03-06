package main

import (
	_ "github.com/qnsoft/online_pay/routers"

	// "github.com/qnsoft/online_pay/utils/pay/fengfu"
	// "github.com/qnsoft/online_pay/utils/pay/paytool"

	"github.com/astaxie/beego"
)

func main() {
	//payTool := fengfu.FengFuPay{}
	//fmt.Printf("%v", payTool)
	//3.6. 支付卡开通发短信接口--------------
	// cardModel := paytool.CardModel{
	// 	AccountNumber: "6258091711738406",   //卡号
	// 	Phone:         "13938202388",        //手机号
	// 	HolderName:    "宋晓辉",                //姓名
	// 	IdCard:        "410185198209190514", //身份证号
	// 	Cvv2:          "903",                //Cvv2
	// 	Expired:       "0423",               //有效期
	// }
	// payTool := fengfu.FengFuPay{}
	// _channelType := "" //签约小额为空 //通道标示小额channelType：ffqr大额channelType：ffkj
	// rtModel := payTool.CardSign(cardModel, _channelType)
	// fmt.Println(rtModel)

	//支付卡开通发短信(大额必走接口)------------------------------
	// cardModel := paytool.CardModel{
	// 	AccountNumber: "6258091711738406",   //卡号
	// 	Phone:         "13938202388",        //手机号
	// 	HolderName:    "宋晓辉",                //姓名
	// 	IdCard:        "410185198209190514", //身份证号
	// 	Cvv2:          "903",                //Cvv2
	// 	Expired:       "0423",               //有效期
	// }
	// fmt.Printf("%v", cardModel)
	// payTool := fengfu.FengFuPay{}
	//rtModel := payTool.CardSign(cardModel, "ffkj")
	//确认短信(大额必走接口)
	/* _merOrderNumber :="2019110510301919"
	   _smsCode:="379854"
	rtModel := payTool.CardSignConfirm(cardModel, "ffkj", _merOrderNumber, _smsCode) */
	//信用卡出金-----------------------------------------
	// cardModelA := paytool.CardModel{
	// 	AccountNumber: "6225258803477017",   //卡号
	// 	Phone:         "13938202388",        //手机号
	// 	HolderName:    "宋晓辉",                //姓名
	// 	IdCard:        "410185198209190514", //身份证号
	// 	Cvv2:          "307",                //Cvv2
	// 	Expired:       "1222",               //有效期
	// }
	// fmt.Printf("%v", cardModelA)
	// payTool := fengfu.FengFuPay{}
	// _channelType := "ffqrdh" //通道标示小额channelType：ffqrdh大额channelType：ffkj
	// _city := "郑州市"
	// _money := 100.0
	// _fee := 0.3
	// rtModel := payTool.CardOut(cardModelA, _channelType, _city, _money, _fee)

	//信用卡出金订单查询
	// cardModelA := paytool.CardModel{
	// 	AccountNumber: "6225258803477017",   //卡号
	// 	Phone:         "13938202388",        //手机号
	// 	HolderName:    "宋晓辉",                //姓名
	// 	IdCard:        "410185198209190514", //身份证号
	// 	Cvv2:          "307",                //Cvv2
	// 	Expired:       "1222",               //有效期
	// }
	// fmt.Printf("信用卡出金卡号%#v", cardModelA)
	// _merOrderNumber := "201911151224328297"
	// rtModel := payTool.CardOutOrder(cardModel, _merOrderNumber)
	// fmt.Printf("%#v", rtModel)
	//信用卡入金

	// cardModelB := paytool.CardModel{
	// 	AccountNumber: "6225768311530149",   //卡号
	// 	Phone:         "13938202388",        //手机号
	// 	HolderName:    "宋晓辉",                //姓名
	// 	IdCard:        "410185198209190514", //身份证号
	// 	Cvv2:          "999",                //Cvv2
	// 	Expired:       "1120",               //有效期
	// }
	// fmt.Printf("%v", cardModelB)
	// _AmountB := 862.88
	// _channelType := "ffqrdh" //通道标示小额channelType：ffqrdh大额channelType：ffkj
	// if _AmountB > 1000 {
	// 	_channelType = "fk"
	// }
	// _extraFee := 0.5
	// rtModel := payTool.CardIn(cardModelB, _channelType, _AmountB, _extraFee)
	// fmt.Printf("%v", rtModel)
	//信用卡入金订单查询
	// _merOrderNumber := "201911051635117832"
	// rtModel := payTool.CardInOrder(_merOrderNumber)
	//信用卡余额查询
	// _channelType := "ffqrdh"
	// _accountNumber := "6225258803477017"
	// rtModel := payTool.CardBanlance(_channelType, _accountNumber)
	// fmt.Printf("%v", rtModel)
	beego.Run()
}
