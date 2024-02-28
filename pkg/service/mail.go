package service

import (
	"log"
	"strconv"
	"strings"

	cnfg "github.com/Shakezidin/pkg/config"
	"gopkg.in/gomail.v2"
)

type Messages struct {
	Username string
	Email    string
	Amount   int
	Messages string
	Subject  string
}


func SendConfirmationEmail(cnfg *cnfg.Conf, bookingDetails Messages) error {
	sender := cnfg.EMAIL
	password := cnfg.PASSWORD

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	recipient := strings.TrimSpace(strings.Trim(bookingDetails.Email, `"`))
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", bookingDetails.Subject)
	m.SetBody("text/plain", bookingDetails.Messages+strconv.Itoa(bookingDetails.Amount))

	d := gomail.NewDialer("smtp.gmail.com", 587, sender, password)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Could not send mail %v", err)
		return err
	} else {
		log.Printf("Email Sent Successfully")
	}
	return nil
}
