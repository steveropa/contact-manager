package main

//SpamService comment
type SpamService struct {
}

//IsKnownSpam comment
func (service SpamService) IsKnownSpam(emailAddress string) bool {
	if emailAddress == "knownspam@spam.com" {
		return true
	}
	return false
}
