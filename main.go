package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	handler := http.NewServeMux()
	///we create a new router to expose our api
	//to our users
	handler.HandleFunc("/v1/phone-numbers", PhoneInfo)
	//Every time a  request is sent to the endpoint ("/v1/phone-numbers")
	//the function PhoneInfo will be invoked
	http.ListenAndServe("0.0.0.0:8080", handler)
	//we tell our api to listen to all request to port 8080.
}

// this is a handler for the enpoint
func PhoneInfo(w http.ResponseWriter, r *http.Request) {
	//reading the Query parameter from the URL path
	originalPhoneNumber := r.URL.Query().Get("phoneNumber")
	if strings.Count(originalPhoneNumber, " ") > 2 {
		fmt.Fprintf(w, "Invalid")
		return
	}
	originalCc := r.URL.Query().Get("countryCode")
	cc := originalCc
	// first check if country code length is greater than 2
	if len(cc) > 2 {
		fmt.Fprintf(w, "Invalid")
		return
	}
	result, err := validatePhone(originalPhoneNumber, originalCc)
	if err != nil {
		fmt.Fprintf(w, "Invalid")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
