package util

import (
	"deliveroo-cron/common"
	"fmt"
	"strings"
)

func GenerateOutputFmt() string {
	var lines []string
	for _, field := range common.Fields {
		line := fmt.Sprintf("%s%s%s", field, strings.Repeat(" ", 14-len(field)), "%s")
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
