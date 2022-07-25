package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func sendMsg(quote string) {

	apiurl := "https://api.twilio.com/2010-04-01/Accounts/AC2411e21ed7f6ae68ceadc9f94f71e821/Messages.json"
	method := "POST"
	data := url.Values{}
	data.Set("To", "+917751085627")
	data.Set("MessagingServiceSid", "MGa207c6bfb1b0c8bfc5d1d08544f19f95")
	data.Set("Body", quote)

	//payload := strings.NewReader("To=%2B917328800942&MessagingServiceSid=MGa207c6bfb1b0c8bfc5d1d08544f19f95&Body=hii")

	client := &http.Client{}
	req, err := http.NewRequest(method, apiurl, strings.NewReader(data.Encode()))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic QUMyNDExZTIxZWQ3ZjZhZTY4Y2VhZGM5Zjk0ZjcxZTgyMTo0NDJjNTcyZDg5NjU1Y2IzNDBjNTNlYjBkMjMzOTVlNg==")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
