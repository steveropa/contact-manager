package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnowValidContactModel_IsValid(t *testing.T) {

	contact := ContactModel{

		FirstName:             "Jason",
		LastName:              "Whelehon",
		EmailAddress:          "wheelie33@gmail.com",
		HomeAddressLine1:      "3922 Berview",
		HomeAddressCity:       "St Louis",
		HomeAddressState:      "MO",
		HomeAddressPostalCode: 63125,

		WorkAddressLine1:      "930 Kehrs Mill Rd",
		WorkAddressLine2:      "Suite 302",
		WorkAddressCity:       "Ballwin",
		WorkAddressState:      "Missouri",
		WorkAddressPostalCode: 63011,
	}

	var validationError string
	var isValid = contact.IsValid(&validationError)

	assert.True(t, len(validationError) == 0)
	assert.True(t, isValid)

}
