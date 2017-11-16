package hlc

import (
	"io/ioutil"
	"testing"
)

func TestParseFromByteSlice(t *testing.T) {
	data, _ := ioutil.ReadFile("notes/nocaptcha_response.html")
	parseFromByteSlice(data)
}
