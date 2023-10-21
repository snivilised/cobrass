package i18n

import (
	"strings"
)

func LeadsWith(name, text string) string {
	if strings.HasPrefix(text, name) {
		return text
	}

	return name + " " + text
}
