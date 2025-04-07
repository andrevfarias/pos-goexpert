package entity

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidZipCode = errors.New("invalid zipcode")
	ErrEmptyZipCode   = errors.New("empty zipcode")
)

type ZipCode struct {
	ZipCode string
}

func NewZipCode(zipcode string) (*ZipCode, error) {
	z := &ZipCode{
		ZipCode: zipcode,
	}

	if err := z.Validate(); err != nil {
		return nil, err
	}

	return z, nil
}

func (z *ZipCode) Validate() error {
	if z.ZipCode == "" {
		return ErrEmptyZipCode
	}

	re := regexp.MustCompile(`^\d{8}$`)
	if !re.MatchString(z.ZipCode) {
		return ErrInvalidZipCode
	}

	return nil
}
