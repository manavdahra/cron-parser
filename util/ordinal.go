package util

import (
	"fmt"
	"strings"
)

func ReplaceWithOrdinal(value string, slice []string) string {
	for index, day := range slice {
		value = strings.ReplaceAll(value, day, fmt.Sprintf("%d", index))
	}
	return value
}
