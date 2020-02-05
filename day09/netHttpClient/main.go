package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http Client

// 后台发请求太频繁，导致无法新建http连接，可以复用连接，在全局声明Client变量，共用一个client，设置为KeepAlive长连接
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func main() {
	// res, err := http.Get("http://127.0.0.1:9090/posts/hello/?name=kim&age=18")
	// if err != nil {
	// 	fmt.Println("get url failed", err)
	// 	return
	// }
	// b, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("read res.Body failed", err)
	// 	return
	// }
	// fmt.Println(string(b))

	data := url.Values{} // url value
	urlObj, _ := url.Parse("http://127.0.0.1:9090/posts/hello/")
	data.Set("name", "艾伦")
	data.Set("age", "19")
	urlObj.RawQuery = data.Encode() // url encode 之后的url
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println("get url failed, error:", err)
	// 	return
	// }
	// 禁用KeepAlive的client，后台发请求并不频繁时，用完就关闭连接，禁用KeepAlive长连接
	// tr := &http.Transport{
	// 	DisableKeepAlives: true, 
	// }
	// client := &http.Client{
	// 	Transport: tr,
	// }
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("get url failed, eror:", err)
		return
	}	
	defer res.Body.Close() // 一定要关闭res.Body，太多无效http连接一直存在，无法建立新的连接
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read res.Body failed", err)
		return
	}
	fmt.Println(string(b))
}
