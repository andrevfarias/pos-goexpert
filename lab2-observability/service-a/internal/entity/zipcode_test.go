package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewZipCode_ShouldReturnErrorWhenZipCodeIsInvalid(t *testing.T) {
	// Arrange (Invalid Zipcode)
	rawZipCode := "123456789"

	// Act
	zipCode, err := NewZipCode(rawZipCode)

	// Assert
	require.Error(t, err)
	require.Nil(t, zipCode)
}

func TestNewZipCode_ShouldReturnZipCodeWhenZipCodeIsValid(t *testing.T) {
	// Arrange (Valid Zipcode)
	rawZipCode := "89802000"

	// Act
	zipCode, err := NewZipCode(rawZipCode)

	// Assert
	require.NoError(t, err)
	require.NotNil(t, zipCode)
	require.Equal(t, rawZipCode, zipCode.ZipCode)
}
