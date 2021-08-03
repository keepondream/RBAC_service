package smtp

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type GmailAuth struct {
	User     string
	Password string
}

type GmailSTMP struct {
	Auths        []GmailAuth
	CurrentIndex int
	Host         string
}

// 谷歌账号进行授权  重点,使用个人账号一定要授权
// https://myaccount.google.com/lesssecureapps?pli=1&rapt=AEjHL4NypvcXbFr4vtNMrHaKtWFLcng2b5Pj_p4HpCC4b_wnSRXXaL7sQJfczB0HMoDJ6EtuhZlSmWVhQFM_coEnKA8V4WasLA
// GMAIL_ACCOUNTS 环境变量, 格式 账号::密码;账号::密码;账号::密码 多账号用;分隔,会进行轮询发送
// NewGmailSTMP 构造邮件发送
func NewGmailSTMP() *GmailSTMP {
	account := os.Getenv("GMAIL_ACCOUNTS")

	accountArr := strings.Split(account, ";")

	auths := []GmailAuth{}

	for _, v := range accountArr {
		if v == "" {
			continue
		}
		a := strings.Split(v, "::")
		if len(a) == 2 {
			auths = append(auths, GmailAuth{
				User:     a[0],
				Password: a[1],
			})
		}
	}

	if len(auths) <= 0 {
		logrus.Panic("new NewGmailSTMP failed no account")
	}

	return &GmailSTMP{
		Auths:        auths,
		CurrentIndex: 0,
		Host:         "smtp.gmail.com:587",
	}
}

// SendByGOMail 发送邮件
func (s *GmailSTMP) SendByGOMail(to, subject, body string) error {
	index := s.CurrentIndex % len(s.Auths)
	form := s.Auths[index]
	s.CurrentIndex += 1
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", form.User)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", body)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, form.User, form.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("send email failed , current index: %d, form: %s , err: %v", s.CurrentIndex-1, form.User, err)
	}

	return nil
}
