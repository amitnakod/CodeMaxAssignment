package mailModels


type mailtemplate struct {
  to string
  cc string 
  Bcc string 
  subject string
  message string 
}

type config struct {
  mailservicepro string
  publicKey string
  secretKey string
}