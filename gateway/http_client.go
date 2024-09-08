package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-lambda-sagemaker/types"
	"io/ioutil"
	"net/http"
)

// Post makes an HTTP POST request and returns the response body as a string
func Post(url string, data []byte) ([]types.ItemProduct, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(body)
	// Unmarshal the JSON response into a slice of ItemProduct
	var products []types.ItemProduct
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
