package main

import (
	"net/smtp"

	"github.com/AllenDang/giu"
)

var (
	smtp_server   string
	smtp_client   string
	smtp_password string
	email_subject string
	email_body    string
	email_rcp     string
	email_rcps    []string
	email_status  string = "Waiting..."
)

func connecttosmtp() {
	email_status = "Sending..."
	auth := smtp.PlainAuth("", smtp_client, smtp_password, smtp_server)
	email_rcps = append(email_rcps, email_rcp)

	msg := []byte("To: " + email_rcps[0] + "\r\n" +
		"Subject: " + email_subject + "\r\n" +
		"\r\n" +
		email_body + "\r\n")
	err := smtp.SendMail(smtp_server+":25", auth, smtp_client, email_rcps, msg)
	if err != nil {
		email_status = "Error: " + err.Error()
	} else {
		email_status = "Sent!"
	}
}

func loop() {
	giu.SingleWindow().Layout(
		giu.Align(giu.AlignCenter).To(
			giu.Label("SMTP Server"),
			giu.InputText(&smtp_server),
			giu.Label("SMTP User (your email address)"),
			giu.InputText(&smtp_client),
			giu.Label("SMTP Password"),
			giu.InputText(&smtp_password).Flags(giu.InputTextFlagsPassword),
			giu.Label("E-Mail recipiener"),
			giu.InputText(&email_rcp),
			giu.Label("Subject"),
			giu.InputText(&email_subject),
			giu.Label("E-Mail body"),
			giu.InputTextMultiline(&email_body),
			giu.Button("Send! :)").OnClick(connecttosmtp),
			giu.Label(email_status),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Voltage's SMTP Client", 600, 420, giu.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
