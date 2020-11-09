package main
import (
  "fmt"
  "log"
  mailjet "github.com/mailjet/mailjet-apiv3-go"
  "basicauthmiddleware"
)
func main () {
  mailjetClient := mailjet.NewMailjetClient("cb0d02ce897bba90f89012a44783c23a", "490054fdc83be6067986d8741a570adf")
  messagesInfo := []mailjet.InfoMessagesV31 {
    mailjet.InfoMessagesV31{
      From: &mailjet.RecipientV31{
        Email: "nakod.amit@yahoo.com",
        Name: "amit",
      },
      To: &mailjet.RecipientsV31{
        mailjet.RecipientV31 {
          Email: "nakod.amit@yahoo.com",
          Name: "amit",
        },
      },
      Subject: "Greetings from Mailjet.",
      TextPart: "My first Mailjet email",
      HTMLPart: "<h3>Dear passenger 1, welcome to <a href='https://www.mailjet.com/'>Mailjet</a>!</h3><br />May the delivery force be with you!",
      CustomID: "AppGettingStartedTest",
    },
  }
  messages := mailjet.MessagesV31{Info: messagesInfo }
  res, err := mailjetClient.SendMailV31(&messages)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Data: %+v\n", res)
}
