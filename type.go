package dingbot

// Text 文本json
type dText struct {
	At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Msgtype string  `json:"msgtype"`
}
//Link Link型json
type dLink struct {
	Msgtype string `json:"msgtype"`
	Link    struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	} `json:"link"`
}
//MD Markdown型json
type dMD struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

// AActionCard 整体跳转ActionCard类型
type dAActionCard struct {
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		SingleTitle    string `json:"singleTitle"`
		SingleURL      string `json:"singleURL"`
	} `json:"actionCard"`
	Msgtype string `json:"msgtype"`
}

// DActionCard 独立跳转ActionCard类型
type dDActionCard struct {
	Msgtype    string `json:"msgtype"`
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		Btns           []struct {
			Title     string `json:"title"`
			ActionURL string `json:"actionURL"`
		} `json:"btns"`
	} `json:"actionCard"`
}

// FeedCard FeedCard类型
type dFeedCard struct {
	Msgtype  string `json:"msgtype"`
	FeedCard struct {
		Links []struct {
			Title      string `json:"title"`
			MessageURL string `json:"messageURL"`
			PicURL     string `json:"picURL"`
		} `json:"links"`
	} `json:"feedCard"`
}

// ErrorReport 返回的错误
type dErrorReport struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}



/*
目前Markdown支持

标题
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

引用
> A man who stands for nothing will fall for anything.

文字加粗、斜体
**bold**
*italic*

链接
[this is a link](http://name.com)

图片
![](http://name.com/pic.jpg)

无序列表
- item1
- item2

有序列表
1. item1
2. item2
*/