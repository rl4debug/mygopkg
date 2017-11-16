/*
Package hlc is using to get related sentences or words of Text in both languages Vietnamese-English
	This package is implemented base on crawling data on website helloch4o.vn
	This package automatically bypass captcha on this website.

	Problems: this website allows anonymous users use limit number of requests, so we need to provide a cookie to comsume this library (maybe resolved next time)

	For example: the input text is "he study", so we expect to get some sentences related to that text, like "They study so much" or "He is a pupil" or etc.
*/
package hlc

import (
	"fmt"
)

// MaxTextSearchLength is maximum length of Text
const MaxTextSearchLength = 50

// Request is a request information for translating
type Request struct {
	Text string //Input text to translate
}

//Response is a response of Request
//includes a word, setences and conversations related to Request.Text
type Response struct {
	Word          Word
	Sentences     []Sentense
	Conversations []Conversation
}

//Word includes information of one couple equivalent Vi-En word
type Word struct {
	Word      string
	Pronounce string
	Mean      string
}

//Sentense includes information of one couple equivalent Vi-En sentense
type Sentense struct {
	Ori  string
	Mean string
}

//Conversation which related to Request
type Conversation struct {
	Title                string
	Description          string
	ConversationSetences []ConversationSentense
}

//ConversationSentense is a sentense in a Conversation
type ConversationSentense struct {
	From     string //Indicate name of people in the conversation
	Sentense string
	Mean     string // Mean of sentense, maybe empty
}

//CaptchaError occurs when the captcha arise in http response.
//Should we name 'CaptchaError' or 'ErrorCaptcha'?
// This named way can make 'go doc' tool show messy things
type CaptchaError struct {
	LinkCheck   string
	LinkCaptcha string
}

// Error method is implemented to satisfy Error interface
func (e CaptchaError) Error() string {
	return "Need captcha!"
}

// UnknownError is an error which should be returned when we don't know what happened. Such as the format of returned HTML goes wrong with our expected.
type UnknownError struct {
	Message string
}

// Error method is implemented to satisfy Error interface
func (e UnknownError) Error() string {
	return fmt.Sprint("UnknownError: ", e.Message)
}

// Translator is the HLC translater. (Why don't we define an interface? Because this package is aim only to the website HLC)
type Translator struct {
	Cookie string //Must manually login and get the cookie from browser, this may be ugly, better we should allows users provide username/password, but it's hard now because this site allow user login with Google (OAuth2)
}

//MakeTranslator is a helper function to make a Translator
func MakeTranslator(cookie string) *Translator {
	return &Translator{
		Cookie: cookie,
	}
}

// Translate is the primary method of Translator which processes Request and return Response.
//  Errors:
//   CaptchaError indicate that captcha is required (see 'go doc hlc.CaptchaError').
//   UnknownError indicate that we don't know what kind of error is (see 'go doc hlc.UnknownError').
func (t *Translator) Translate(r *Request) (*Response, error) {
	return nil, UnknownError{}
}

func get(text string) {
}
