package mailModels


type mailtemplate struct {
  to string
  cc string 
  Bcc string 
  message string 
}

type config struct {
  mailservicepro string
  publicKey string
  secretKey string
}