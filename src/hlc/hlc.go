/*
Package hlc is using to get related sentences or words of Text in both language Vietnamese-English
This package is implemented based on website hellochao.vn

	For example: the input text is "he study", so we expected to get some sentences related to that text, like "He is studying" or "He is a puple"
*/
package hlc

// Request is a request information for translating
type Request struct {
	Text   string //Input text to translate
	Cookie string //Must manually login and get the cookie from browser
}

//Response is a response of Request
type Response struct {
}

//Translator the translator
type Translator interface {
	Translate(info Request) Response //Send Request and get Response from hellochao
}
