package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//ContactModel annoying comment to get the lame compiler to leave me alone
type ContactModel struct {
	FirstName        string
	LastName         string
	EmailAddress     string
	WorkEmailAddress string
	FavoriteColor    string
	DateOfBirth      time.Time
	Mvp              bool
	Salutation       string
	Suffix           string
	JobTitle         string

	HomeAddressLine1      string
	HomeAddressLine2      string
	HomeAddressCity       string
	HomeAddressState      string
	HomeAddressPostalCode int

	WorkAddressLine1      string
	WorkAddressLine2      string
	WorkAddressCity       string
	WorkAddressState      string
	WorkAddressPostalCode int

	HomePhoneNumber   string
	WorkPhoneNumber   string
	MobilePhoneNumber string
}

//IsValid is a function to verify if there really is a santa claus
func (model ContactModel) IsValid(validationerror *string) bool {

	if len(model.FirstName) == 0 {

		*validationerror = "First name empty"
		return false
	}
	if len(model.LastName) == 0 {
		*validationerror = "Last name empty"
		return false
	}
	if len(model.EmailAddress) == 0 {
		*validationerror = "email empty"
		return false
	}

	matched, _ := regexp.MatchString(`^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$`, strings.ToUpper(model.EmailAddress))
	if matched == false {
		*validationerror = "Email not valid"

		return false
	}

	if len(model.HomeAddressLine1) == 0 {
		*validationerror = "Home address line 1 empty"
		return false
	}
	if len(model.HomeAddressCity) == 0 {
		*validationerror = "Home address line 2 empty"
		return false
	}
	if len(model.HomeAddressState) == 0 {
		*validationerror = "Home address state empty"
		return false
	}
	if len(strconv.Itoa(model.HomeAddressPostalCode)) == 0 {
		*validationerror = "Home address postal empty"
		return false
	}
	if len(model.WorkAddressLine1) == 0 {
		*validationerror = "Work address line 1 empty"
		return false
	}
	if len(model.WorkAddressCity) == 0 {
		*validationerror = "Work address city empty"
		return false
	}
	if len(model.WorkAddressState) == 0 {
		*validationerror = "Work address state empty"
		return false
	}
	if len(strconv.Itoa(model.HomeAddressPostalCode)) == 0 {
		*validationerror = "Work address postal empty"
		return false
	}

	*validationerror = ""

	return true

}

func main() {

	fmt.Println("Hi everybody! Here we go!")

}
