package creditcard

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrNotNumber = errors.New("<NOT-A-NUMBER>")
	ErrInvalid   = errors.New("<INVALID>")
)

type CreditCard struct {
	// num is the actual credit number with following structure
	// according to https://en.wikipedia.org/wiki/Payment_card_number
	//
	// 1. first six-digit                  - Issuer Identification Number
	// 2. variable length (upto 12 digits) - Individual Accound Identifier
	// 3. last digit                       - check digit
	num string

	// typ denotes the Issuer Identification(MII)
	// Identified by first digit of `num`
	typ string
}

func (c *CreditCard) Type() string {
	return c.typ
}

func (c *CreditCard) Number() string {
	return c.num
}

// getType currently checks only from major card type.
func getType(n string) string {
	for issuerName, issuerFunc := range issuers {
		if issuerFunc(n) {
			return issuerName
		}
	}
	return "<UNKNOWN>"
}

// Normalize removes fancy userfriendly characters
// and returns `raw` number.
func Normalize(num string) string {
	// remove spaces
	num = strings.Join(strings.Fields(num), "")

	// remove '-'
	return strings.Join(
		strings.FieldsFunc(num, func(r rune) bool {
			return r == '-'
		}), "",
	)
}

// Parse takes the creditcard number,  validates it and
// return the instance of `CreditCard` if all good.
// Returns non-nil error if given number is not a valid
// creditcard number.
//
// It uses `Luhn Algorithm` for validation
// https://en.wikipedia.org/wiki/Luhn_algorithm
func Parse(card string) (*CreditCard, error) {
	card = Normalize(card)

	sum := 0
	pairity := len(card) % 2
	for i, c := range card {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, ErrNotNumber
		}
		if i%2 == pairity {
			digit *= 2
		}
		digit = digit/10 + digit%10
		sum += digit
	}
	if sum%10 != 0 {
		return nil, ErrInvalid
	}
	return &CreditCard{num: card, typ: getType(card)}, nil
}

// Pretty return credit card number in more friendly format.
// e.g: "5500000000000004" becomes "5500-0000-0000-0004"
func Pretty(num string) string {
	num = Normalize(num)
	return strings.Join(splitEvery(num, 4), "-")
}

func splitEvery(s string, n int) []string {
	res := make([]string, 0)

	start := 0
	for i, _ := range s {
		if i == 0 {
			continue
		}
		if i%n == 0 {
			res = append(res, s[start:i])
			start = i
		}
	}
	res = append(res, s[start:])
	return res
}
