package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get ...
func Get(url string) []byte {
	// 请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http.NewRequest error:", err.Error())
		return nil
	}
	// 自定义header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:83.0) Gecko/20100101 Firefox/83.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do(GET) error:", err.Error())
		return nil
	}
	// 延迟关闭，等函数结束时关闭resp
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error:", err.Error())
		return nil
	}
	return body
}
