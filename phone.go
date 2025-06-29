package irphone

import (
	"fmt"
	"regexp"
	"strings"
)

var InvalidPhoneError = fmt.Errorf("invalid phone number")

func Validate(input string) error {
	match, _ := regexp.MatchString("^(\\+98|0|98|0098)?9\\d{9}$", input)
	if !match {
		return InvalidPhoneError
	}
	return nil
}

func To09(input string) (string, error) {
	if err := Validate(input); err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	var normalized string
	switch {
	case strings.HasPrefix(input, "+98"):
		normalized = "0" + input[3:]
	case strings.HasPrefix(input, "0098"):
		normalized = "0" + input[4:]
	case strings.HasPrefix(input, "98"):
		normalized = "0" + input[2:]
	case strings.HasPrefix(input, "9"):
		normalized = "0" + input
	case strings.HasPrefix(input, "09"):
		normalized = input
	default:
		return "", InvalidPhoneError
	}
	return normalized, nil
}

func To98(input string) (string, error) {
	if err := Validate(input); err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	var normalized string
	switch {
	case strings.HasPrefix(input, "+98"):
		normalized = "0" + input[3:]
	case strings.HasPrefix(input, "0098"):
		normalized = "0" + input[4:]
	case strings.HasPrefix(input, "98"):
		normalized = "0" + input[2:]
	case strings.HasPrefix(input, "9"):
		normalized = "0" + input
	case strings.HasPrefix(input, "09"):
		normalized = input
	default:
		return "", InvalidPhoneError
	}
	return "98" + normalized[1:], nil
}

func ToPlus98(input string) (string, error) {
	if err := Validate(input); err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	var normalized string
	switch {
	case strings.HasPrefix(input, "+98"):
		normalized = "0" + input[3:]
	case strings.HasPrefix(input, "0098"):
		normalized = "0" + input[4:]
	case strings.HasPrefix(input, "98"):
		normalized = "0" + input[2:]
	case strings.HasPrefix(input, "9"):
		normalized = "0" + input
	case strings.HasPrefix(input, "09"):
		normalized = input
	default:
		return "", InvalidPhoneError
	}
	return "+98" + normalized[1:], nil
}

// normalizePhone ensures the phone number is in the 09 format for consistency.
func normalizePhone(input string) (string, error) {
	return To09(input)
}

func sMobile(input string) (bool, error) {
	normalized, err := normalizePhone(input)
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(normalized, "09"), nil
}

func MaskPhone(input string) (string, error) {
	normalized, err := normalizePhone(input)
	if err != nil {
		return "", err
	}
	return normalized[:4] + "****" + normalized[7:], nil
}

func FormatWithDash(input string) (string, error) {
	normalized, err := normalizePhone(input)
	if err != nil {
		return "", err
	}
	return normalized[:4] + "-" + normalized[4:7] + "-" + normalized[7:], nil
}

func ExtractPhoneNumbers(text string) []string {
	re := regexp.MustCompile(`(?i)(\+98|0098|98|0)?9\d{9}`)
	matches := re.FindAllString(text, -1)
	var validNumbers []string
	for _, match := range matches {
		// Accept numbers that start with +98, 0098, 98, 0, or 9
		if err := Validate(match); err == nil {
			// If the match starts with '98' but not '+98' or '0098', add '+98' prefix for consistency
			if strings.HasPrefix(match, "98") && !strings.HasPrefix(match, "+98") && !strings.HasPrefix(match, "0098") {
				validNumbers = append(validNumbers, "+"+match)
			} else {
				validNumbers = append(validNumbers, match)
			}
		}
	}
	return validNumbers
}
