package main

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	send("hello there")
}

var cfg Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("failed to read viper config"))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(errors.New("failed to unmarshal viper config"))
	}

}

type (
	// Attachment defines a file.
	Attachment struct {
		Filename string
		Filepath string
	}

	// Attachments define a set of attachments.
	Attachments []Attachment

	// Config defines the configuration.
	Config struct {
		Gmail struct {
			Email    string `mapstructure:"email"`
			Password string `mapstructure:"password"`
		} `mapstructure:"gmail"`
		NTC struct {
			Email string `mapstructure:"email"`
		}
	}
)

// send
func send(body string) {
	from := cfg.Gmail.Email
	pass := cfg.Gmail.Password
	to := cfg.NTC.Email

	var (
		serverAddr = "smtp.gmail.com"
		password   = pass
		emailAddr  = from
		portNumber = 465
		tos        = []string{
			to,
		}
		cc        = []string{}
		delimeter = "**=myohmy689407924327"
	)

	attachments := Attachments{
		Attachment{
			Filename: "ntc-form.pdf",
			Filepath: "./ntc-form.pdf",
		},
		Attachment{
			Filename: "diff.png",
			Filepath: "./diff.png",
		},
		Attachment{
			Filename: "govt-id-one.png",
			Filepath: "./govt-id-one.png",
		},
		Attachment{
			Filename: "govt-id-two.png",
			Filepath: "./govt-id-two.png",
		},
	}

	log.Println("======= Test Gmail client (with attachment) =========")
	log.Println("NOTE: user need to turn on 'less secure apps' options")
	log.Println("URL:  https://myaccount.google.com/lesssecureapps\n\r")

	tlsConfig := tls.Config{
		ServerName:         serverAddr,
		InsecureSkipVerify: true,
	}

	log.Println("Establish TLS connection")
	conn, connErr := tls.Dial("tcp", fmt.Sprintf("%s:%d", serverAddr, portNumber), &tlsConfig)
	if connErr != nil {
		log.Panic(connErr)
	}
	defer conn.Close()

	log.Println("create new email client")
	client, clientErr := smtp.NewClient(conn, serverAddr)
	if clientErr != nil {
		log.Panic(clientErr)
	}
	defer client.Close()

	log.Println("setup authenticate credential")
	auth := smtp.PlainAuth("", emailAddr, password, serverAddr)

	if err := client.Auth(auth); err != nil {
		log.Panic(err)
	}

	log.Println("Start write mail content")
	log.Println("Set 'FROM'")
	if err := client.Mail(emailAddr); err != nil {
		log.Panic(err)
	}
	log.Println("Set 'TO(s)'")
	for _, to := range tos {
		if err := client.Rcpt(to); err != nil {
			log.Panic(err)
		}
	}

	writer, writerErr := client.Data()
	if writerErr != nil {
		log.Panic(writerErr)
	}

	//basic email headers
	sampleMsg := fmt.Sprintf("From: %s\r\n", emailAddr)
	sampleMsg += fmt.Sprintf("To: %s\r\n", strings.Join(tos, ";"))
	if len(cc) > 0 {
		sampleMsg += fmt.Sprintf("Cc: %s\r\n", strings.Join(cc, ";"))
	}
	sampleMsg += "Subject: Golang example send mail in HTML format with attachment\r\n"

	log.Println("Mark content to accept multiple contents")
	sampleMsg += "MIME-Version: 1.0\r\n"
	sampleMsg += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", delimeter)

	//place HTML message
	log.Println("Put HTML message")
	sampleMsg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	sampleMsg += "Content-Type: text/html; charset=\"utf-8\"\r\n"
	sampleMsg += "Content-Transfer-Encoding: 7bit\r\n"
	sampleMsg += fmt.Sprintf("\r\n%s", "<html><body><h1>Hi There</h1>"+
		"<p>this is sample email (with attachment) sent via golang program</p></body></html>\r\n")

	//place file
	log.Println("Put file attachment")

	for _, a := range attachments {
		sampleMsg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
		sampleMsg += "Content-Type: text/plain; charset=\"utf-8\"\r\n"
		sampleMsg += "Content-Transfer-Encoding: base64\r\n"
		sampleMsg += "Content-Disposition: attachment;filename=\"" + a.Filename + "\"\r\n"
		//read file
		rawFile, fileErr := ioutil.ReadFile(a.Filepath)
		if fileErr != nil {
			log.Panic(fileErr)
		}
		sampleMsg += "\r\n" + base64.StdEncoding.EncodeToString(rawFile)
	}

	//write into email client stream writter
	log.Println("Write content into client writter I/O")
	if _, err := writer.Write([]byte(sampleMsg)); err != nil {
		log.Panic(err)
	}

	if closeErr := writer.Close(); closeErr != nil {
		log.Panic(closeErr)
	}

	client.Quit()

	log.Print("done.")
}
