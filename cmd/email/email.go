package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"gopkg.in/gomail.v2"
)

type rawMessage struct {
	from    string
	to      string
	subject string
	name    string
	email   string
	phone   string
	message string
}

// SendMail sends emails
func SendMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	d := gomail.NewPlainDialer(
		viper.GetString("smtpHost"),
		viper.GetInt("port"),
		viper.GetString("username"),
		viper.GetString("password"),
	)

	var bodyJSON map[string]string
	decoder.Decode(&bodyJSON)

	m := buildMessage(&rawMessage{
		from:    viper.GetString("username"),
		to:      viper.GetString("targetAddress"),
		subject: bodyJSON["name"] + "的反馈",
		name:    bodyJSON["name"],
		email:   bodyJSON["email"],
		phone:   bodyJSON["phone"],
		message: bodyJSON["message"],
	})

	err := d.DialAndSend(m)
	if err != nil {
		w.WriteHeader(400)
		encoder.Encode(map[string]string{
			"code":    "400",
			"message": err.Error(),
		})
	}

	w.WriteHeader(200)
	encoder.Encode(map[string]string{
		"code":    "200",
		"message": "success",
	})
}

func buildMessage(rawMessage *rawMessage) (m *gomail.Message) {
	m = gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {rawMessage.from},
		"To":      {rawMessage.to},
		"Subject": {rawMessage.subject},
	})

	message := fmt.Sprintf(
		"<h1>用户反馈</h1><p>名字：</p>%s<p>邮箱：</p>%s<p>电话：</p>%s<p>信息：</p>%s",
		rawMessage.name, rawMessage.email, rawMessage.phone, rawMessage.message)
	m.SetBody("text/html", message)
	return
}
