package src

import (
  "os"
  "gopkg.in/yaml.v2"
)
type Config struct {
  mailservicepro string `yaml:"mailservicepro"`
  publicKey string `yaml:"publicKey"`
  secretKey string `yaml:"secretKey"`
}
func start(){
  f, err := os.Open("config.yml")
if err != nil {
    
}
defer f.Close()

var cfg Config
decoder := yaml.NewDecoder(f)
err = decoder.Decode(&cfg)
if err != nil {
    
}


}