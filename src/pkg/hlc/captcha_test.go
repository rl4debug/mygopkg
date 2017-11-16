package hlc

import (
	"testing"
)

func TestDownloadImage(t *testing.T) {
	captcha := &captcha{
		LinkCheck:   "https://www.hellochao.vn/capcha.php?type=5&st=1&rf=0.46294297312567165",
		LinkCaptcha: "https://www.hellochao.vn/capcha.php?type=5&st=2&rf=0.46294397312567165",
		Cookie:      "G_ENABLED_IDPS=google; CUSER_INFO=a%28.3a3.f2mc3amelg4.sbce%2A%2A52c2o2h0m0a1af69b98n3cufnd%4001e271%291ivAU; PHPSESSID=In6PWkkYHCU%2C9v8dsrTx3rNabH5; TawkConnectionTime=0; __tawkuuid=e::hellochao.vn::2Wf3kUPwQS96pwBerC//AY4pmLcrRyUbjY8Ovkw+zYPnBEFcy0uJie1f5jO2Gp1W::2; Tawk_57d7cf81cccb3b470ce023e1=vs58.tawk.to::0; _ga=GA1.2.338720649.1505846065; _gid=GA1.2.137572233.1510725443; _gat=1; __utmt=1; __utma=48272063.338720649.1505846065.1510782924.1510809363.32; __utmb=48272063.1.10.1510809363; __utmc=48272063; __utmz=48272063.1505846065.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)",
	}

	// var err = captcha.resolve()
	// if err != nil {
	// 	t.Error(err)
	// }

	_ = captcha
}
