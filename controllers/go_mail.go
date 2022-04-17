package controllers

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT. Bioskop Fox <bioskopfoxofficial@gmail.com>"
const CONFIG_AUTH_EMAIL = "bioskopfoxofficial@gmail.com"
const CONFIG_AUTH_PASSWORD = "bioskopfox123"

func sendMail(to []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)
	//auth = menampung credentials untuk keperluan otentikasi ke mail server
	//CONFIG_AUTH_EMAIL, adalah alamat email yang digunakan untuk mengirim email.
	//CONFIG_AUTH_PASSWORD, adalah password alamat email yang digunakan untuk mengirim email.
	//smtpAddr = untuk kombinasi host dan port mail server

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, to, []byte(body))
	//sendMail() digunakan untuk mengirim email. Empat data yang disisipkan pada fungsi tersebut dijadikan satu dalam format tertentu, lalu disimpan ke variabel body.
	if err != nil {
		fmt.Print("ini error di dalem send Mail")
		return err
	}

	return nil
}

func GoMail(userEmail string) {
	to := []string{userEmail}
	subject := "Register Success Message!"
	message := "Congratulations, your email has successfully registered at Bioskop Fox !!"

	err := sendMail(to, subject, message)
	if err != nil {
		fmt.Print("ini setelah send Mail")
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
