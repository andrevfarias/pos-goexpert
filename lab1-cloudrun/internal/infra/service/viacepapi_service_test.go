package service

import (
	"testing"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/config"
	"github.com/stretchr/testify/suite"
)

type ViaCepApiServiceTestSuite struct {
	suite.Suite
	viacepService *ViaCepApiService
}

func (s *ViaCepApiServiceTestSuite) SetupSuite() {
	cfg, err := config.LoadConfig("./../../../")
	s.Require().NoError(err)

	s.viacepService = NewViaCepApiService(cfg.ViacepAPIBaseURL)
}

func TestViaCepApiServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ViaCepApiServiceTestSuite))
}

func (s *ViaCepApiServiceTestSuite) TestGetAddressByZipCode_ShouldReturnAddress() {
	// Arrange (Valid Zipcode from Chapecó)
	zipcode := "89802000"

	// Act
	address, err := s.viacepService.GetAddressByZipcode(zipcode)

	// Assert
	s.Require().NoError(err)
	s.Equal("Chapecó", address.City)
}

func (s *ViaCepApiServiceTestSuite) TestGetAddressByZipCode_ShouldReturnAPIError() {
	// Arrange (Invalid Zipcode)
	zipcode := "123456789"

	// Act
	address, err := s.viacepService.GetAddressByZipcode(zipcode)

	// Assert
	s.Require().Error(err)
	s.Equal(err, ErrApiError)
	s.Nil(address)
}

func (s *ViaCepApiServiceTestSuite) TestGetAddressByZipCode_ShouldReturnZipCodeNotFoundError() {
	// Arrange (Inexistent Zipcode)
	zipcode := "12345678"

	// Act
	address, err := s.viacepService.GetAddressByZipcode(zipcode)

	// Assert
	s.Require().Error(err)
	s.Equal(err, ErrZipCodeNotFound)
	s.Nil(address)
}
