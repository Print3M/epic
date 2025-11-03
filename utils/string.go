package utils

import "strings"

// Helper function to parse list of items
func StringToSlice(str string, sep string) []string {
	items := strings.Split(str, sep)
	result := make([]string, 0, len(items))

	for _, item := range items {
		trimmed := strings.TrimSpace(item)

		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}
