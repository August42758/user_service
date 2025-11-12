package validator

import (
	"regexp"
	"unicode/utf8"
)

var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
var MinPasswordLen = 5

type IValidator interface {
	MatchEmail(s string, regex *regexp.Regexp) error
	CountMinAmountCharsInPassword(s string, n int) error
}

type Validator struct{}

func (v *Validator) MatchEmail(s string, regex *regexp.Regexp) error {
	if !regex.MatchString(s) {
		return ErrWrongEmailFormat
	}

	return nil
}

func (v *Validator) CountMinAmountCharsInPassword(s string, n int) error {
	if utf8.RuneCountInString(s) < n {
		return ErrWrongPasswordFormat
	}

	return nil
}
