package main

import (
  "net/http"
  "net/url"
  "fmt"
  "strings"
  "os"
)

func main() {
  fmt.Printf("start\n")
  accountSid := "AC19dd69d986658313d0c871bdbf0c37de"
  authToken := os.Getenv("TWILIO")
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
 
  // Build out the data for our message
  v := url.Values{}
  v.Set("To","+18475626149")
  v.Set("From","+12247740141")
  v.Set("Body","Brooklyn's in the house!")
  rb := *strings.NewReader(v.Encode())
 
  // Create client
  client := &http.Client{}
 
  req, _ := http.NewRequest("POST", urlStr, &rb)
  req.SetBasicAuth(accountSid, authToken)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
 
  // Make request
  resp, err := client.Do(req)
  if err != nil {
    fmt.Printf("fail")
  } else {
    fmt.Printf(resp.Status)
  }
}
