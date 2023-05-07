package health

import (
	"fmt"
	"regexp"
)

func defineStatus(code int, status string) string {
	var emoji string

	if code < successCodesThreshold {
		emoji = "✔️"
	} else if code < redirectCodesThreshold {
		emoji = "❔"
	} else {
		emoji = "❌"
	}

	return fmt.Sprintf("%s %s", emoji, status)
}

func defineSpeed(ms int64) string {
	var emoji string

	if ms < fastSpeedThreshold {
		emoji = "🚀"
	} else if ms < mediumSpeedThreshold {
		emoji = "✈️"
	} else {
		emoji = "🐢"
	}

	return fmt.Sprintf("%s %d ms", emoji, ms)
}

func retriveServerName(body []byte) string {
	rg := regexp.MustCompile(titleRgPattern)
	match := rg.FindStringSubmatch(string(body))
	if len(match) < 2 {
		return "–"
	}
	return match[1]
}
