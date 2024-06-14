package alert

import (
    "fmt"
    "net/smtp"
    "sysmon/internal/config"
)

var cfg *config.Config

func InitConfig(config *config.Config) {
    cfg = config
}

func SendAlert(message string) {
    fmt.Println("ALERT:", message)
    if cfg.Alerts.Email.Enabled {
        sendEmailAlert(message)
    }
}

func sendEmailAlert(message string) {
    auth := smtp.PlainAuth("", cfg.Alerts.Email.SenderEmail, cfg.Alerts.Email.SenderPassword, cfg.Alerts.Email.SMTPServer)
    to := []string{cfg.Alerts.Email.RecipientEmail}
    msg := []byte("To: " + cfg.Alerts.Email.RecipientEmail + "\r\n" +
        "Subject: System Alert\r\n" +
        "\r\n" + message + "\r\n")
    err := smtp.SendMail(fmt.Sprintf("%s:%d", cfg.Alerts.Email.SMTPServer, cfg.Alerts.Email.SMTPPort), auth, cfg.Alerts.Email.SenderEmail, to, msg)
    if err != nil {
        fmt.Println("Error sending email:", err)
    }
}
