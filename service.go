package main
//This is go file which will serve as service to the handler.

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"strings"
)


type PhoneInformation struct {
	PhoneNumber      string      `json:"phoneNumber,omitempty"`
	CountryCode      string      `json:"countryCode,omitempty"`
	AreaCode         string      `json:"areaCode,omitempty"`
	LocalPhoneNumber string      `json:"localPhoneNumber,omitempty"`
	Error            interface{} `json:"error,omitempty"`
}

type ErrorBody struct {
	CountryCode string `json:"countryCode,omitempty"`
}

func getPhoneInfo(originalPhoneNumber string, num *phonenumbers.PhoneNumber) PhoneInformation {
	// make output body for valid phone numbers
	region := phonenumbers.GetRegionCodeForNumber(num)
	localNum := fmt.Sprint(num.GetNationalNumber())
	areaCode := localNum[0:3]
	localPhoneNumber := localNum[3:]

	result := PhoneInformation{
		PhoneNumber:      originalPhoneNumber,
		CountryCode:      region,
		AreaCode:         areaCode,
		LocalPhoneNumber: localPhoneNumber,
	}
	return result
}

func getPhoneErrorInfo(originalPhoneNumber string, originalCc string) PhoneInformation {
	var text string
	if originalCc != "" {
		text = "invalid country code"
	} else {
		text = "required value is missing"
	}
	// make custom body for error msg
	err1 := PhoneInformation{
		PhoneNumber: originalPhoneNumber,
		Error: ErrorBody{
			CountryCode: text,
		},
	}
	return err1
}

func validatePhone(originalPhoneNumber string, originalCc string) (PhoneInformation, error) {
	phoneNumber := originalPhoneNumber
	// adding + (if not added to the phone numbers) before parsing to check the numbers
	if !strings.Contains(phoneNumber, "+") {
		phoneNumber = "+" + phoneNumber
	}
	var num *phonenumbers.PhoneNumber
	// used lib to check for validation of the phone number
	num, err := phonenumbers.Parse(phoneNumber, strings.ToUpper(originalCc))
	if err != nil {
		fmt.Println(err.Error())
		var p PhoneInformation
		return p, err
	}
	countryCode := originalCc
	if originalCc == "" {
		countryCode = phonenumbers.GetRegionCodeForNumber(num)
	}
	// used lib to check if appropriate country code is given with the phone number
	if !phonenumbers.IsValidNumberForRegion(num, countryCode) {
		return getPhoneErrorInfo(originalPhoneNumber, originalCc), nil
	} else {
		return getPhoneInfo(originalPhoneNumber, num), nil
	}
}
