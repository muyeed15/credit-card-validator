package validation

import (
	"strconv"
	"strings"
)

func CheckLuhn(number string) string {
	if len(number) != 15 && len(number) != 16 {
		return "invalid"
	}

	digits := strings.Split(number, "")

	var sum int
	for i := len(digits) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(digits[i])

		if err != nil {
			return "invaild"
		}

		if (len(digits)-i)%2 == 0 {
			digit *= 2

			if digit > 9 {
				digit -= 9 // equivalent of getting the last digit of the sum
			}
		}

		sum += digit
	}

	if sum%10 == 0 {
		return "valid"
	}

	return "invalid"

}
