package dingbot

import (
	"encoding/json"
)

const URL = "https://oapi.dingtalk.com/robot/send?access_token="

// Text 推送文本
func Text(token string,content string,AtMobiles []string, AtUserIds []string, IsAtAll bool) error {
	t:=new(dText)
	t.Text = struct {
		Content string `json:"content"`
	}(struct{ Content string }{Content:content})
	t.At= struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	}(struct {
		AtMobiles []string
		AtUserIds []string
		IsAtAll   bool
	}{AtMobiles: AtMobiles, AtUserIds: AtUserIds, IsAtAll: IsAtAll})
	t.Msgtype="text"
	if jsonByte, err := json.Marshal(t);err != nil {
		return err
	} else if _, err := Send(jsonByte, URL+token);err != nil {
		return err
	}
	return nil
} // Link 推送Link
func Link(token string,content string,Title string,PicUrl string,MessageUrl string) error {
	t:=new(dLink)
	t.Link= struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	}(struct {
		Text       string
		Title      string
		PicUrl     string
		MessageUrl string
	}{Text: content, Title: Title, PicUrl: PicUrl, MessageUrl: MessageUrl})
	t.Msgtype="link"
	if jsonByte, err := json.Marshal(t);err != nil {
		return err
	} else if _, err := Send(jsonByte, URL+token);err != nil {
		return err
	}
	return nil
}

// MD 推送Markdown
func MD(token string,title string,content string,AtMobiles []string, AtUserIds []string, IsAtAll bool) error {
	t:=new(dMD)
	t.At= struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	}(struct {
		AtMobiles []string
		AtUserIds []string
		IsAtAll   bool
	}{AtMobiles:AtMobiles, AtUserIds:AtUserIds, IsAtAll:IsAtAll})
	t.Markdown= struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}(struct {
		Title string
		Text  string
	}{Title:title, Text:content})
	t.Msgtype="markdown"
	if jsonByte, err := json.Marshal(t);err != nil {
		return err
	} else if _, err := Send(jsonByte, URL+token);err != nil {
		return err
	}
	return nil
}

// AActionCard 推送 整体跳转ActionCard
func AActionCard(token string,Title string,content string,BtnOrientation string, SingleTitle string, SingleURL string) error {
	t:=new(dAActionCard)
	t.ActionCard= struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		SingleTitle    string `json:"singleTitle"`
		SingleURL      string `json:"singleURL"`
	}(struct {
		Title          string
		Text           string
		BtnOrientation string
		SingleTitle    string
		SingleURL      string
	}{Title:Title, Text:content, BtnOrientation:BtnOrientation, SingleTitle:SingleTitle, SingleURL:SingleURL})
	t.Msgtype="ActionCard"
	if jsonByte, err := json.Marshal(t);err != nil {
		return err
	} else if _, err := Send(jsonByte, URL+token);err != nil {
		return err
	}
	return nil
}

// DActionCard 推送 独立跳转ActionCard
func DActionCard(token, Title, content, BtnOrientation string, Btns []struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}) error {
	t:=new(dDActionCard)
	t.ActionCard= struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		Btns           []struct {
			Title     string `json:"title"`
			ActionURL string `json:"actionURL"`
		} `json:"btns"`
	}(struct {
		Title          string
		Text           string
		BtnOrientation string
		Btns           []struct {
			Title     string `json:"title"`
			ActionURL string `json:"actionURL"`
		}
	}{Title:Title, Text:content, BtnOrientation:BtnOrientation, Btns:Btns})

	t.Msgtype="ActionCard"
	if jsonByte, err := json.Marshal(t);err != nil {
		return err
	} else if _, err := Send(jsonByte, URL+token);err != nil {
		return err
	}
	return nil
}

// FeedCard 推送FeedCard
func FeedCard(token string,Links []struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}) error {
	t:=new(dFeedCard)
	t.FeedCard= struct {
		Links []struct {
			Title      string `json:"title"`
			MessageURL string `json:"messageURL"`
			PicURL     string `json:"picURL"`
		} `json:"links"`
	}(struct {
		Links []struct {
			Title      string `json:"title"`
			MessageURL string `json:"messageURL"`
			PicURL     string `json:"picURL"`
		}
	}{Links:Links})
	t.Msgtype="FeedCard"
	if jsonByte, err := json.Marshal(t);err != nil {
		return err
	} else if _, err := Send(jsonByte, URL+token);err != nil {
		return err
	}
	return nil
}