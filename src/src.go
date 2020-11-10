package src

import (
  "os"
  "gopkg.in/yaml.v2"
  "github.com/gorilla/mux"
  "encoding/json"
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
)

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

myRouter := mux.NewRouter().StrictSlash(true)
myRouter.HandleFunc("/mail", basicauthmiddleware.BasicAuthMiddleware(callMail)).Methods("POST")
log.Fatal(http.ListenAndServe(":10000", myRouter))

}

func callMail(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var message mailtemplate 
    json.Unmarshal(reqBody, &message)
    provider:=mailprovider(cfg)
    provider.send(message)
}