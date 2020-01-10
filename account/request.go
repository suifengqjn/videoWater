package account

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"myTool/sys"
	"net/http"
	"sync"
)

const remoteHost  = "http://106.12.220.252:8001"

type Account struct {
	AccType int `json:"acc_type"`
	Count   int `json:"count"`
	Time    string	`json:"time"`
	Msg string	`json:"msg"`
	AppId string `json:"-"`
	lock sync.Mutex
}

func GetAccountInfo(appId string) *Account {
	url := remoteHost + "/vd/account_info"
	method := "POST"


	param := make(map[string]string)
	param["host"] = sys.GetSysInfo().IP
	param["app_id"] = appId

	buf, _ :=json.Marshal(param)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(buf))

	if err != nil {
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var acc  = Account{}

	err = json.Unmarshal(body,&acc)
	if err != nil {
		return nil
	}
	acc.AppId = appId
	acc.lock = sync.Mutex{}
	return &acc
}

func (a *Account)Add() error  {
	url := remoteHost + "/vd/count"
	method := "POST"

	param := make(map[string]string)
	param["app_id"] = a.AppId

	buf, _ :=json.Marshal(param)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(buf))

	if err != nil {
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if string(body) == "success" {
		return nil
	}
	return err
}