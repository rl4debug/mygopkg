package hlc

import (
	"bytes"
	"fmt"
	_ "fmt"

	// _ "github.com/antchfx/xpath"
	"github.com/antchfx/xquery/html"
)

func parseFromByteSlice(data []byte) (*Response, error) {
	var response *Response
	root, err := htmlquery.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	nodes := htmlquery.Find(root, "//*[contains(@class, 'c-a-v')]")

	fmt.Println("len =", len(nodes))
	for node := range nodes {
		_ = node
		// fmt.Println("gi vay")
		// fmt.Println(node)
	}

	_, _ = root, err
	return response, nil
}
