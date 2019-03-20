package main

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
)

func main() {
	//hhtp Post
	conn, err := http.PostForm("http://client.6che.vip/client/user/user/detail", url.Values{"uid": {"2"}})
	//conn.Header.Add("CLIENT-LOGIN-ID", "138")
	//conn.Header.Add("CLIENT-LOGIN-TOKEN", "03e4368053d106d5eac6eaa0674f2da1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer conn.Body.Close()

	body, err := ioutil.ReadAll(conn.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}
