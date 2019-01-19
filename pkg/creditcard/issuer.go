package creditcard

import (
	"strconv"
	"strings"
)

type issuer func(num string) bool

var (
	issuers = map[string]issuer{
		"American-Express": AmericanExpress,
		"Visa":             Visa,
		"MasterCard":       MasterCard,
	}
)

func AmericanExpress(num string) bool {
	return strings.HasPrefix(num, "34") || strings.HasPrefix(num, "37")
}

func ChinaTUnion(num string) bool {
	return strings.HasPrefix(num, "31")
}

func ChinaUnionPay(num string) bool {
	return strings.HasPrefix(num, "62")
}

func CISS(num string) bool {
	return strings.HasPrefix(num, "686566")
}

func DinersClubInternationl(num string) bool {
	if strings.HasPrefix(num, "36") || strings.HasPrefix(num, "3095") || strings.HasPrefix(num, "38") || strings.HasPrefix(num, "39") {
		return true
	}
	n, err := strconv.Atoi(num[0:3])
	if err != nil {
		return false
	}
	if n >= 300 && n <= 305 {
		return true
	}
	return false
}

func DinersClubUSorCanada(num string) bool {
	return strings.HasPrefix(num, "54") || strings.HasPrefix(num, "55")
}

func Visa(num string) bool {
	return strings.HasPrefix(num, "4")
}

func MasterCard(num string) bool {
	n, err := strconv.Atoi(num[0:2])
	if err != nil {
		return false
	}
	if n >= 51 && n <= 55 {
		return true
	}

	n, err = strconv.Atoi(num[0:6])
	if err != nil {
		return false
	}
	if n >= 222100 && n <= 272099 {
		return true
	}
	return false
}
