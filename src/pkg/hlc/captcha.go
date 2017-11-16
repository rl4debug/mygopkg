package hlc

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	_ "image/jpeg" //This is use for Decode jpeg image data, we can treat it acts as a plugin of image.Image type
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const captchaBaseLink = "https://www.hellochao.vn/dictionary.php?isAJ=1&act=verify_security_code"

//Define captcha data structure correspond to HLC captcha
type captcha struct {
	LinkCheck   string
	LinkCaptcha string
	Cookie      string
}

func (c *captcha) resolve() (err error) {
	//Download captchas
	checkImg, captchaImg, err := c.downloadCaptchas()
	if err != nil {
		return
	}

	x, y, err := getVerifyPoint(checkImg, captchaImg)

	if err != nil {
		return
	}

	client := &simpleHTTPGetClient{}
	var queryParams = map[string]string{
		"x": strconv.Itoa(x),
		"y": strconv.Itoa(y),
	}
	response, err := client.Get(captchaBaseLink, queryParams, defaultHTTPHeaders, map[string]string{"Cookie": c.Cookie})
	if err != nil {
		return
	}

	bodySlice, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	body := string(bodySlice)
	//Success
	if strings.Contains(body, "<rs>1</rs>") {
		return nil
	}

	return errors.New("Resolving error failed")
}

func (c *captcha) downloadCaptchas() (checkImg, captchaImg image.Image, err error) {
	client := &simpleHTTPGetClient{}

	//Step download check image
	response, err := client.Get(c.LinkCheck, nil, defaultHTTPHeaders, map[string]string{"Cookie": c.Cookie})
	if err != nil {
		return nil, nil, err
	}
	bodySlice, _ := ioutil.ReadAll(response.Body)
	//Write file for debug
	// err = ioutil.WriteFile("notes/checkImg.jpg", bodySlice, 0777)

	img, _, err := image.Decode(bytes.NewReader(bodySlice))
	checkImg = img

	//Step download captcha image
	response, err = client.Get(c.LinkCaptcha, nil, defaultHTTPHeaders, map[string]string{"Cookie": c.Cookie})
	if err != nil {
		return nil, nil, err
	}
	bodySlice, _ = ioutil.ReadAll(response.Body)

	//Write file for debug
	// err = ioutil.WriteFile("notes/captchaImg.jpg", bodySlice, 0777)

	img, _, err = image.Decode(bytes.NewReader(bodySlice))
	captchaImg = img

	//Step to check
	if err != nil {
		return nil, nil, err
	}

	return
}

func getVerifyPoint(checkImg, captchaImg image.Image) (x, y int, err error) {
	err = nil

	var checkImgBound = checkImg.Bounds()
	var captchaImgBound = captchaImg.Bounds()
	var widthCheck = checkImgBound.Dx()
	var heightCheck = checkImgBound.Dy()
	var widthCaptcha = captchaImgBound.Dx()
	var heightCaptcha = captchaImgBound.Dx()

	var initX = widthCheck / 2
	var initY = heightCheck / 2
	x, y = initX, initY
	xInc, yInc := x/2, y/2
	var checkColor = checkImg.At(x, y)

	for x <= widthCaptcha || y <= heightCaptcha {
		captchaPointColor := captchaImg.At(x, y)

		if compareColor(checkColor, captchaPointColor) {
			return
		}

		if x > widthCaptcha {
			x = initX
			y += yInc
		} else {
			x += xInc
		}
	}

	return 0, 0, errors.New("Not found satify point")
}

func convertColorTo8BitRGBA(c color.Color) (uint32, uint32, uint32, uint32) {
	// return c.RGBA()
	r, g, b, a := c.RGBA()

	var convert = func(i uint32) uint32 {
		if i <= 255 {
			return 0
		}

		var mod = (i + 1) % 256
		var div = (i + 1) / 256

		if mod > 0 {
			return div
		}

		return div - 1
	}
	return convert(r), convert(g), convert(b), convert(a)
}

//This function may have risk
func compareColor(c1, c2 color.Color) bool {
	r1, g1, b1, _ := convertColorTo8BitRGBA(c1)
	r2, g2, b2, _ := convertColorTo8BitRGBA(c2)

	var sum = int(math.Abs(float64(r1)-float64(r2)) + math.Abs(float64(g1)-float64(g2)) + math.Abs(float64(b1)-float64(b2)))

	if sum/3 <= 30 {
		return true
	}

	return false

	// fmt.Println(sum / 3)
	// var cR = r1 - r2
	// var cG = g1 - g2
	// var cB = b1 - b2
	// var uR = r1 + r2
	// var dst = cR*cR*(2+uR/256) + cG*cG*4 + cB*cB*(2+(255-uR)/256)

	// _ = dst
	// fmt.Println(dst)

	// fmt.Println(r1, g1, b1, a1, "--", r2, g2, b2, a2)
}
