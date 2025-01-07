package util

import (
	"regexp"
	"strings"
)

func CleanTime(data string) string {
	re := regexp.MustCompile(`\[\d{2}:\d{2}\.\d{2}\]`)

	// Loại bỏ tất cả các đoạn thời gian
	cleanedLyrics := re.ReplaceAllString(data, "")

	// Loại bỏ khoảng trắng thừa
	cleanedLyrics = strings.TrimSpace(cleanedLyrics)

	return cleanedLyrics
}
