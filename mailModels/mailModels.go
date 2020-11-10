package mailModels


type mailtemplate struct {
  to string
  cc string 
  Bcc string 
  subject string
  message string 
}

type Config struct {
  mailservicepro string `yaml:"mailservicepro"`
  publicKey string `yaml:"publicKey"`
  secretKey string `yaml:"secretKey"`
}