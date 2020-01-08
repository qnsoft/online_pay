package paytool

/*
用户卡实体
*/
type CardModel struct {
	AccountNumber string `json:"accountNumber"` //卡号
	Phone         string `json:"phone"`         //手机号
	HolderName    string `json:"holderName"`    //卡主姓名
	IdCard        string `json:"idCard"`        //身份证号
	Cvv2          string `json:"cvv2"`          //Cvv2
	Expired       string `json:"expired"`       //有效日期
}
