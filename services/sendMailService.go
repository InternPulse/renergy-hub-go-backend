package services

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

func SendEmail(fromAddr string, toAddr []string, subject string, body string) error {
	if fromAddr == "" {
		fromAddr = os.Getenv("EMAIL_FROM")
	}

	senderName := os.Getenv("EMAIL_SENDER_NAME")
	emailServer := os.Getenv("EMAIL_SERVER")
	emailPort := os.Getenv("EMAIL_PORT")
	emailPassword := os.Getenv("EMAIL_PASSWORD")

	if fromAddr == "" || len(toAddr) == 0 || emailServer == "" || emailPassword == "" {
		return fmt.Errorf("missing required email parameters")
	}

	from := mail.Address{"", fromAddr}

	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("%s %s", senderName, from.String())
	headers["Subject"] = subject

	var toAddrs []string
	for _, addr := range toAddr {
		to := mail.Address{"", addr}
		toAddrs = append(toAddrs, to.String())
	}
	headers["To"] = strings.Join(toAddrs, ", ")

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	servername := fmt.Sprintf("%s:%s", emailServer, emailPort)
	host, _, err := net.SplitHostPort(servername)
	if err != nil {
		return fmt.Errorf("invalid server address: %v", err)
	}

	auth := smtp.PlainAuth("", fromAddr, emailPassword, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %v", err)
	}
	defer c.Close()

	if err = c.Auth(auth); err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	if err = c.Mail(from.Address); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}

	for _, addr := range toAddr {
		if err = c.Rcpt(addr); err != nil {
			return fmt.Errorf("failed to set recipient %s: %v", addr, err)
		}
	}

	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("failed to create data writer: %v", err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write message: %v", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close data writer: %v", err)
	}

	if err = c.Quit(); err != nil {
		return fmt.Errorf("failed to quit SMTP connection: %v", err)
	}

	return nil
}
