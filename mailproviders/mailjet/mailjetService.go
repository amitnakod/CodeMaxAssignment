Package mailjet

import (
  "fmt"
  "log"
  mailjet "github.com/mailjet/mailjet-apiv3-go"
  "basicauthmiddleware"
)

type mailj struct{
  publickey string
  secretkey string
}

func (j mailj)send(msg mailtemplate){
  mailjetClient := mailjet.NewMailjetClient(j.publickey, j.secretkey)
  messagesInfo := []mailjet.InfoMessagesV31 {
    mailjet.InfoMessagesV31{
      From: &mailjet.RecipientV31{
        Email: "nakod.amit@yahoo.com",
        Name: "amit",
      },
      To: &mailjet.RecipientsV31{
        mailjet.RecipientV31 {
          Email: msg.To,
          Name: "amit",
        },
      },
      Subject: msg.Subject,
      TextPart: msg.messages,
    },
  }
  messages := mailjet.MessagesV31{Info: messagesInfo }
  res, err := mailjetClient.SendMailV31(&messages)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Data: %+v\n", res)
}