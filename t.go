package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	ssWeb, err := getISSRCnt()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ssWeb)
}

func getISSRCnt() (string, error) {
	ssUrl := "http://www.myshadowsocks.me"
	resp, err := http.Get(ssUrl)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != 200 {
		return "", errors.New(string(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	ssWeb := string(body)

	return ssWeb, nil
}
