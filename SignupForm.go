package main

//SignupForm is a signup form
type SignupForm struct {
}

//Signup signs someone up
func (form SignupForm) Signup(firstName string, lastName string, emailAddress string) bool {
	access := DataAccess{}
	success := access.SaveRecord(firstName, lastName, emailAddress)
	return success
}
