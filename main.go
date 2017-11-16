package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	// _ "github.com/rl4debug/mygopkg/src/examples"
	// _ "github.com/rl4debug/mygopkg/src/examples/design-patterns/strategy"
)

func (m MyStruct)Method1 () {

}
func (m *MyStruct) Method1() {

}
func (m MyStruct) Method1() {

}

// Reader is an
type Reader interface {
	Read(p []byte) (n int, err error)
}

//AlertCounter is an unexported type that
//contains an integer counter for alerts.
type AlertCounter int

// Writer is and
type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
	Reader
	Writer
}

//MyType This is
//my type
type MyType struct {
	MyStruct1
	Name  string /*name of type*/
	Value string // struct of type
}

func (mt MyType) f1() {

}

func test(rw ReadWriter) {
	rw.Read(nil)
}

type intError struct {
	Value int
}

func (e intError) Error() string {
	return fmt.Sprintf("This is intError %v.", e.Value)
}

func throwError1() error {
	return intError{
		Value: 10,
	}
}

type Mutatable struct {
	a int
	b int
}

func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
}

func (m *Mutatable) Mutate() {
	m.a = 5
	m.b = 7
}

type I1 interface {
	i1()
}

type ImpI1 int

func (i ImpI1) i1() {

}

func (i ImpI1) i2() {

}

type I2 interface {
	i2()
}

type I interface {
	I1
	I2
}

func testVar1() int {
	var temp = 10
	return temp
}

func testInterface() {
	impi1 := ImpI1(10)
	func(i1 I) {
		var temp interface{} = i1
		switch temp.(type) {
		case I1:
			return
		case string:
			return
		}
		i1.i1()
	}(impi1)
}

type II interface {
	I1
	I2
}

func func1(slices *[]int) []int {
	*slices = append(*slices, 10)
	return *slices
}

func read(p []int) int {
	return len(p)
}

type MyStruct1 struct {
	int
	Hi int
}
type MyStruct struct {
	MyStruct1
	int
	Value int
}

//This is used for ...
func main() {
	var err interface{}
	err = intError{}
	err, ok := err.(string)
	if ok {
		fmt.Print("This is int error")
	}

	return
	// strategy.Test()
	// examples.Test()
	// examples.Test()

	arr := make([]int, 3)
	read(arr)
	fmt.Print(len(arr))
	slices := []int{1, 2, 3}
	fmt.Print(func1(&slices))
	fmt.Print(slices)
	return

	c := make(chan int)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recover in f", r)
				<-c
			}
		}()

		panic("sample panic")
		fmt.Print("Hehe Hihi")
	}()

	c <- 1

	return
	request, _ := http.NewRequest("GET", "https://www.hellochao.vn/tu-dien-tach-ghep-am/?lang=&type=sentence&act=search&sct=fall+through", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Cookie", "PHPSESSID=%2Ce943osf%2CpP4kxmcrV7vwaKHe01; IS_WARNING_PRE_CLOSE=1; _ga=GA1.2.2003095628.1510488725; _gid=GA1.2.1428047797.1510488725; __utma=48272063.2003095628.1510488725.1510488726.1510488726.1; __utmb=48272063.7.10.1510488726; __utmc=48272063; __utmz=48272063.1510488726.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Print("error roi")
		fmt.Print(err)
		return
	}

	body, _ := ioutil.ReadAll(response.Body)
	strBody := string(body)
	//_imgSecCodeIntro_Image
	if strings.Contains(strBody, "_imgSecCode_Image") {
		fmt.Print("Found")
	}
	ioutil.WriteFile("response", body, 0666)
}
