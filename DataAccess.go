package main

import (
	"regexp"
	"strings"
)

//DataAccess type
type DataAccess struct {
}

//SaveRecord function
func (access DataAccess) SaveRecord(firstName string, lastName string, emailAddress string) bool {

	connection := DatabaseConnection{"(local)", 30, "Contacts"}
	if !connection.IsConnected() {
		return false
	}

	if len(firstName) <= 10 {
		matched, _ := regexp.MatchString(`^[a-zA-Z]+$`, firstName)
		if !matched {
			return false

		}
	} else {
		return false
	}

	// From the RFC822 Spec
	pattern := `^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$`

	match, _ := regexp.MatchString(pattern, strings.ToUpper(emailAddress))

	if !match {
		return false
	}

	isspam := SpamService{}.IsKnownSpam(emailAddress)
	if isspam {
		return false
	}

	return true
}

//DatabaseConnection type
type DatabaseConnection struct {
	_servername  string
	_timeout     int
	_datasetname string
}

//IsConnected function
func (connection DatabaseConnection) IsConnected() bool {
	return !(len(connection._servername) == 0)
}
