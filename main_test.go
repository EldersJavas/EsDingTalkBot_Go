package dingbot

import (
	"reflect"
	"testing"
)

func TestAActionCard(t *testing.T) {
	type args struct {
		token          string
		Title          string
		content        string
		BtnOrientation string
		SingleTitle    string
		SingleURL      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AActionCard(tt.args.token, tt.args.Title, tt.args.content, tt.args.BtnOrientation, tt.args.SingleTitle, tt.args.SingleURL); (err != nil) != tt.wantErr {
				t.Errorf("AActionCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDActionCard(t *testing.T) {
	type args struct {
		token          string
		Title          string
		content        string
		BtnOrientation string
		Btns           []struct {
			Title     string `json:"title"`
			ActionURL string `json:"actionURL"`
		}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DActionCard(tt.args.token, tt.args.Title, tt.args.content, tt.args.BtnOrientation, tt.args.Btns); (err != nil) != tt.wantErr {
				t.Errorf("DActionCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFeedCard(t *testing.T) {
	type args struct {
		token string
		Links []struct {
			Title      string `json:"title"`
			MessageURL string `json:"messageURL"`
			PicURL     string `json:"picURL"`
		}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FeedCard(tt.args.token, tt.args.Links); (err != nil) != tt.wantErr {
				t.Errorf("FeedCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLink(t *testing.T) {
	type args struct {
		token      string
		content    string
		Title      string
		PicUrl     string
		MessageUrl string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Link(tt.args.token, tt.args.content, tt.args.Title, tt.args.PicUrl, tt.args.MessageUrl); (err != nil) != tt.wantErr {
				t.Errorf("Link() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMD(t *testing.T) {
	type args struct {
		token     string
		title     string
		content   string
		AtMobiles []string
		AtUserIds []string
		IsAtAll   bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MD(tt.args.token, tt.args.title, tt.args.content, tt.args.AtMobiles, tt.args.AtUserIds, tt.args.IsAtAll); (err != nil) != tt.wantErr {
				t.Errorf("MD() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSend(t *testing.T) {
	type args struct {
		message []byte
		pushURL string
	}
	tests := []struct {
		name    string
		args    args
		want    *dErrorReport
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := send(tt.args.message, tt.args.pushURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText(t *testing.T) {
	type args struct {
		token     string
		content   string
		AtMobiles []string
		AtUserIds []string
		IsAtAll   bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Text(tt.args.token, tt.args.content, tt.args.AtMobiles, tt.args.AtUserIds, tt.args.IsAtAll); (err != nil) != tt.wantErr {
				t.Errorf("Text() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
