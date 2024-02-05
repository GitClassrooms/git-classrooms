package mail

import (
	"bytes"
	"crypto/tls"
	mailConfig "gitlab.hs-flensburg.de/gitlab-classroom/config/mail"
	"html/template"
	"net/url"

	gomail "gopkg.in/gomail.v2"
)

type GoMailRepository struct {
	publicURL *url.URL
	dialer    *gomail.Dialer
}

func NewMailRepository(publicURL *url.URL, config mailConfig.Config) (*GoMailRepository, error) {
	dialer := gomail.NewDialer(config.GetHost(), config.GetPort(), config.GetUser(), config.GetPassword())
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &GoMailRepository{
		publicURL: publicURL,
		dialer:    dialer,
	}, nil
}

func (m *GoMailRepository) SendClassroomInvitation(to string, subject string, data ClassroomInvitationData) error {
	t, err := template.ParseFiles(
		"./templates/base.tmpl.html",
		"./templates/invitation.tmpl.html",
	)
	if err != nil {
		return err
	}
	publicURL, err := m.generateExternalURL(data.InvitationPath)
	if err != nil {
		return err
	}
	data.InvitationPath = publicURL.String()

	return m.sendMail(to, subject, t, data)
}

func (m *GoMailRepository) SendAssignmentNotification(to string, subject string, data AssignmentNotificationData) error {
	t, err := template.ParseFiles(
		"./templates/base.tmpl.html",
		"./templates/assignmentNotification.tmpl.html",
	)
	if err != nil {
		return err
	}
	publicURL, err := m.generateExternalURL(data.JoinPath)
	if err != nil {
		return err
	}
	data.JoinPath = publicURL.String()

	return m.sendMail(to, subject, t, data)
}

func (m *GoMailRepository) sendMail(to string, subject string, t *template.Template, data interface{}) error {
	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "base", data); err != nil {
		return err
	}

	result := tpl.String()

	mail := gomail.NewMessage()
	mail.SetHeader("From", m.dialer.Username)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", result)

	return m.dialer.DialAndSend(mail)
}

func (m *GoMailRepository) generateExternalURL(path string) (*url.URL, error) {
	newPath, err := url.JoinPath(m.publicURL.String(), path)
	if err != nil {
		return nil, err
	}
	return url.Parse(newPath)
}