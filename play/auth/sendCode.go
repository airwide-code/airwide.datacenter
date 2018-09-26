
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
)

func main() {
	endpoint := "https://api.twilio.com/2010-04-01/Accounts/AC25c34e873eb0348a2a7b9510f9282319/Messages.json"
	v := url.Values{}
	v.Set("To", "+27729745087")
	v.Add("From", "+14342774779")
	v.Add("Body", "Your Airwide Code is 24680")
	payload := strings.NewReader(v.Encode())

	var username string = "AC25c34e873eb0348a2a7b9510f9282319"
	var passwd string = "c052c7e3068c0f1e64ba5067836b10d4"

	req, _ := http.NewRequest("POST", endpoint, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control","no-cache")
	req.SetBasicAuth(username, passwd)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Fatal error occured")
	}

	defer res.Body.Close();
	body, _ := ioutil.ReadAll(res.Body);

	fmt.Println(string(body));

	// TODO cater for http redirects. https://stackoverflow.com/questions/16673766/basic-http-auth-in-go
}
