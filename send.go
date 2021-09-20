package dingbot

// from https://github.com/CatchZeng/dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const httpTimoutSecond = time.Duration(30) * time.Second

func Send(message []byte,pushURL string) (*dErrorReport, error) {
res := &dErrorReport{}

reqBytes := message


req, err := http.NewRequest(http.MethodPost, pushURL, bytes.NewReader(reqBytes))
if err != nil {
return res, err
}
req.Header.Add("Accept-Charset", "utf8")
req.Header.Add("Content-Type", "application/json")

client := new(http.Client)
client.Timeout = httpTimoutSecond
resp, err := client.Do(req)
if err != nil {
return res, err
}
defer resp.Body.Close()

resultByte, err := ioutil.ReadAll(resp.Body)
if err != nil {
return res, err
}

err = json.Unmarshal(resultByte, &res)
if err != nil {
return res, fmt.Errorf("unmarshal http response body from json error = %w", err)
}

if res.Errcode != 0 {
return res, fmt.Errorf("send message to dingtalk error = %s", res.Errmsg)
}

return res, nil
}