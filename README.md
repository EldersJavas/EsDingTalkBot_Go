# EsDingTalkBot_Go

[![Go Reference](https://pkg.go.dev/badge/github.com/EldersJavas/EsDingTalkBot_Go.svg)](https://pkg.go.dev/github.com/EldersJavas/EsDingTalkBot_Go)
[![GitHub license](https://img.shields.io/github/license/EldersJavas/EsDingTalkBot_Go?color=red&logo=apache&logoColor=red&style=flat-square)](https://github.com/EldersJavas/EsDingTalkBot_Go/blob/main/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/EldersJavas/EsDingTalkBot_Go?color=green&style=flat-square)](https://github.com/EldersJavas/EsDingTalkBot_Go/stargazers)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/EldersJavas/EsDingTalkBot_Go?logo=go&style=flat-square)


Using DingTalk GroupBot with one line

So easy!

DingTalk Doc:https://developers.dingtalk.com/document/robots/custom-robot-access

# Use

## Text
![image](https://user-images.githubusercontent.com/55266266/133955143-1bf09b70-57d4-488f-bd04-bc16ab76180a.png)

```go
//func Text(token string, content string, AtMobiles []string, AtUserIds []string, IsAtAll bool) error
dingbot.Text("123","hello",[]string{},[]string{},false)
```
## Link
![image](https://user-images.githubusercontent.com/55266266/133955155-20e9ecf4-a950-43ad-96ef-02968c562aa7.png)

```go
//func Link(token string, content string, Title string, PicUrl string, MessageUrl string) error
dingbot.Link("123","Hello World!","Hello","https://github.com/fluidicon.png","https://github.com/EldersJavas/EsDingTalkBot_Go")
```
## MarkDown
![image](https://user-images.githubusercontent.com/55266266/133955161-11a264e0-2a1f-46e3-bcd6-3fefd98245ae.png)

```go
//func MD(token string, title string, content string, AtMobiles []string, AtUserIds []string, IsAtAll bool) error
dingbot.MD("123","# Hello",[]string{},[]string{},false)
```
## AActionCard
![image](https://user-images.githubusercontent.com/55266266/133955175-40d80988-3ecc-4281-8e84-d4ff37b232c4.png)

```go
//func AActionCard(token string, Title string, content string, BtnOrientation string, SingleTitle string, SingleURL string) error
dingbot.AActionCard("123","hello","Hello World!","1","EsDingTalkBot_Go","https://github.com/EldersJavas/EsDingTalkBot_Go")
```
## DActionCard
![image](https://user-images.githubusercontent.com/55266266/133955181-13e3166c-577e-45b3-9fff-0e97b971478f.png)

```go
/*
func DActionCard(token, Title, content, BtnOrientation string, Btns []struct {
Title     string `json:"title"`
ActionURL string `json:"actionURL"`
}) error
*/
var Btn []struct {
  Title     string `json:"title"`
  ActionURL string `json:"actionURL"`}
var Add struct{
  Title     string `json:"title"`
  ActionURL string `json:"actionURL"`}
Add.Title="EsDingTalkBot_Go"
Add.ActionURL="https://github.com/EldersJavas/EsDingTalkBot_Go"
Btn = append(Btn, Add)
err := dingbot.DActionCard("123", "hello", "", "1", Btn)
```
## FeedCard
![image](https://user-images.githubusercontent.com/55266266/133955184-d43e7aa0-cdc1-4ab6-9062-8f5c2dc69454.png)

```go
/*
func func FeedCard(token string, Links []struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}) error
*/
var Links []struct{
  Title      string `json:"title"`
  MessageURL string `json:"messageURL"`
  PicURL     string `json:"picURL"`}
var Add struct{
  Title      string `json:"title"`
  MessageURL string `json:"messageURL"`
  PicURL     string `json:"picURL"`}
Add.Title="EsDingTalkBot_Go"
Add.PicURL="https://github.com/fluidicon.png"
Add.MessageURL="https://github.com/EldersJavas/EsDingTalkBot_Go"
Links = append(Links, Add)
err := dingbot.FeedCard("123", Links)
```

# Error Back
```go	
err := FeedCard("123", Links)
if err != nil {
  return
}
```



