package luhn

import (
	"strings"
	"unicode"
)

func Valid(cardNumber string) bool {


	cardNumber = strings.ReplaceAll(cardNumber, " ", "")

	if len(cardNumber) <= 1 { return false }


	sumFromRight := 0
	sumFromLeft := 0
	for i := 0; i < len(cardNumber); i++ {
		if unicode.IsDigit(rune(cardNumber[i])) == false {
			return false
		}
		if i % 2 != 0 {

			num := int(cardNumber[i] - '0') * 2
			if num > 9  {
				num = num - 9
			}
			sumFromLeft += num

			sumFromRight += int(cardNumber[i] - '0')
		} else {
			num := int(cardNumber[i] - '0') * 2
			if num > 9  {
				num = num - 9
			}
			sumFromRight += num

			sumFromLeft += int(cardNumber[i] - '0')
		}
	}

	if sumFromLeft % 10 == 0 || sumFromRight % 10 == 0 { return true }


	return false
}
