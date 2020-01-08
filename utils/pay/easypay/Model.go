package EasyPay

/*
易宝(易付通接口参数))
*/
var (
	//Merno     string = "MCH104702"                                //接入商户编号
	//Key       string = "db7770445fdde3a663f6b93d7a08382a"         //商户蜜钥KEY
	//Orgno     string = "PT1047"                                   //接入机构编号
	CustomerNo string = ""                                      //易宝商户号
	AppKey     string = "OPR:10000466938"                       //"OPR:10015386666"
	Version    string = "2.0"                                   //版本号
	PayUrl     string = "https://openapi.yeepay.com/yop-center" //接口url
	PublicKey         = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6p0XWjscY+gsyqKRhw9M
eLsEmhFdBRhT2emOck/F1Omw38ZWhJxh9kDfs5HzFJMrVozgU+SJFDONxs8UB0wM
ILKRmqfLcfClG9MyCNuJkkfm0HFQv1hRGdOvZPXj3Bckuwa7FrEXBRYUhK7vJ40a
fumspthmse6bs6mZxNn/mALZ2X07uznOrrc2rk41Y2HftduxZw6T4EmtWuN2x4CZ
8gwSyPAW5ZzZJLQ6tZDojBK4GZTAGhnn3bg5bBsBlw2+FLkCQBuDsJVsFPiGh/b6
K/+zGTvWyUcu+LUj2MejYQELDO3i2vQXVDk7lVi2/TcUYefvIcssnzsfCfjaorxs
uwIDAQAB
-----END PUBLIC KEY-----
`)
	PrivateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAyi4CPuuJpiXTYNU8JQjwZzR9296n8Nk3R1XNrg3eC08kCPJ0
a9IXhYYXY9DmYumoKnpCMswB/15jvT3S604uGFyVzn4jv3WMCgOoEpLUXuhSmOAt
ls19cgOHxn6iTFe0WFaC+iJ+dB7q+qxp2rzUkGam4gyToMucqJCdeDHstYjKqqlr
wA7SjfUqsbW8iMWA/44/+qi1sJrvPMEsEmhnTt8DRYM2o7B4ESXvRvSIQ9M6+Xe9
h7yRqHnau1Wj5944SLM3D0BOqKm1wHXWURigXoQOFkgoyxJfO1x9982T68a0JQUx
1VRs38Y2weEL2M9IZRKFodZ7cXgx034MoXT0KQIDAQABAoIBAE5sTEi1ul3tPDHK
ZJsE/aM6tDsuAbyltRCBkteZ33BJkXO7ADe2dZNAcSAxrrvdYHKhm73ybi/tIAkO
yxqk57Jhb2Z6fUlNwjU8N9yEgcYSnb9Y7bYaHTe1cNwo0LGGL5k4bXOzH5PRTlpP
uHu1VUpx1vfF4xmzASBNgBdYXjYHmeNLm3YM7s/fuIOKaYY1p3D6cHGjH7A2Ztyb
Rcr1KdPizBNAplKq8bzjWQJKRUfz8noxh5yId6bR84Xv1YzXLiAuD6YZAjQWpAgQ
q7aIU8tGHnlnwwFEwDuLT1xfyVIZcYz7qkeq6UGP9LkHTALU6Ylh1SWq+3hnSvNU
J7rjhUUCgYEA65JVDzNKZLxhv8ERGtG4FCPpIvtf8L+zvF/eP/KJ6Jiaj6ZA/M+I
b/y3r/8A+YA3IJoCwt+rv1w7oRX1TTpsOHlFnp1Oqc6J66BI4YvdQCnEUX8w3zn1
oN8D5tSfVkDp+QNN97TUcnQqVXF/zreKg2ynAMC4jzT11zE5+C1EBP8CgYEA27Zh
NCkv5En47HFiTqsL9V28XJIqZMZ3KiG3YknVPg82i3flbmnKO0y0CURae0u1iVuE
8+igV9lIgOzu3s12PT8JYF8V4mQw7Z3lzQKpOwzcZk02R6xdgu6nVPfJJaaQhYhc
fW8rBmOpRVqH1poNz7uSIj6jt4i0T072YEDhPtcCgYEAvUgoJn2M8iJX3I7SIp2O
bkhohJ48+aHfVqYlW3aTjxxHXNM2dqx+sZT7N2QLsW5vXTfCrjiwtKL0mrudDb+6
wynWYdt+IHz/Xx+8T9ZlIbAtuVoct9DHZ8qqROhlWxXvqhEYrcdTyN9EQMFg84WQ
u4crLTStVwwz2QQJ0zrz45sCgYAxbro1+yJMUZxQzj3ZB0DchqdviG9DVyfOceGJ
mqUAnnmeLW6QRfr7FonXH+rfnKwOaGnkWq1gtoFKiIRB2qZEp18bsKkC05nyjDj8
xCGBKKaZ2bthqtUrNTiutEUsVGplsMVmb1GV0WxvLywqy2RaHZCGZeN90qMGyPRD
ubUz5wKBgCUbMtc/eUNRj5CqxbSor5MWctZO8yV4zNFVg3O5U/XwZ6HhlMrfr7qv
pQuYtmoliiSu6XvDDPUBQxpVj55cvifI28jH5vfZOaCbQlg3xdKK9aQEOJsNJIy/
U7w1V3aI4GAPq8UjENPYcLGLpEOG25a8HwVJkYyzeVom3EcyFuWQ
-----END RSA PRIVATE KEY-----
`)
)
