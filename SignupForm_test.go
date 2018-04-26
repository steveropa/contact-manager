package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidEmail_IsSaved(t *testing.T) {

	saved := SignupForm{}.Signup("J", "W", "wheelie33@gmail.com")
	assert.True(t, saved, "Signup failed to save")

}

func TestEmailMissingTld_IsNotSaved(t *testing.T) {

	saved := SignupForm{}.Signup("J", "W", "wheelie33@gmail")
	assert.False(t, saved, "Saved it anyway")
}

func TestForSanity(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestEmailMissingAtSign_IsNotSaved(t *testing.T) {
	saved := SignupForm{}.Signup("J", "W", "wheelie33.gmail.com")
	assert.False(t, saved)
}

func TestGenericData_IsSaved(t *testing.T) {
	saved := SignupForm{}.Signup("Jason", "Whelehon", "wheelie@gmail.com")
	assert.True(t, saved)
}

func TestLongFirstName_IsNotSaved(t *testing.T) {
	saved := SignupForm{}.Signup("JasonJasonJasonJasonJasonJason", "Whelehon", "wheelie33@gmail.com")
	assert.False(t, saved)
}

func TestHyphenatedLastName_IsSaved(t *testing.T) {
	saved := SignupForm{}.Signup("Jason", "Smith-Johnson", "wheelie33@gmail.com")
	assert.True(t, saved)
}

func TestInvalidEmail_IsNotSaved(t *testing.T) {
	saved := SignupForm{}.Signup("J", "W", "wheelie33@gmailcom")
	assert.False(t, saved)
}

func TestValidEmailKnownSpam_IsNotSaved(t *testing.T) {
	saved := SignupForm{}.Signup("J", "W", "knownspam@spam.com")
	assert.False(t, saved)
}
