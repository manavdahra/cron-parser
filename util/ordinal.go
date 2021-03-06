package util

import (
	"fmt"
	"github.com/thoas/go-funk"
	"strings"
)

func ReplaceWithOrdinal(value string, slice []string) string {
	index := funk.IndexOf(slice, strings.ToUpper(value))
	if index < 0 {
		return value
	}
	return fmt.Sprintf("%d", index)
}
