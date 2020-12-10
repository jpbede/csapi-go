package validation

import (
	"errors"
	"regexp"
)

// ValidateDate checks if given date is in expected api format
func ValidateDate(date string) error {
	if date != "" {
		if matched, _ := regexp.MatchString(`\d{4}-\d{2}-\d{2}`, date); !matched {
			return errors.New("invalid date format. expected format is yyyy-mm-dd")
		}
	}
	return nil
}
