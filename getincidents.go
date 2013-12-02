package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	. "github.com/cconger/pd/lib"
)

const subdomain = "webdemo"

//TODO: To make extensible kick this request into the incidents.go file
//  Make it in a function that takes params to modify the things you might
//  change like sortby, fields and limit and then return a pagerDutyResponse to
//  fulfill the request.
//
//  Secondly: this can then be responsible for pulling the token from
//  commandline flag or from a dot-file.
func main() {
	client := &http.Client{}

	var token string
	flag.StringVar(&token, "token", "", "The token used to contact pager duty")
	flag.Parse()
	if token == "" {
		log.Fatal("Token is a required flag. Got:", token)
	}

	urlString := fmt.Sprintf("https://%s.pagerduty.com/api/v1/incidents", subdomain)
	url, _ := url.Parse(urlString)
	vals := url.Query()

	vals.Add("fields", "created_on,assigned_to_user,status")
	vals.Add("sort_by", "created_on:desc")
	vals.Add("limit", "10")
	url.RawQuery = vals.Encode()

	req, _ := http.NewRequest("GET", url.String(), nil)
	tokenHeaderVal := fmt.Sprintf("Token token=%s", token)
	req.Header.Add("Authorization", tokenHeaderVal)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to reach pagerduty: %v", err)
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	var response PagerDutyResponse
	err = dec.Decode(&response)
	if err != nil {
		log.Fatal("Failed to parse response from pagerduty: %v", err)
	}
	fmt.Print(response.PrettyPrint())
}
