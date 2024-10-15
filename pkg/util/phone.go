package util

import "strings"

func FormatPhone(phone string) string {
	if !strings.HasPrefix(phone, "+") {
		phone = "+" + phone
	}
	return phone
}

