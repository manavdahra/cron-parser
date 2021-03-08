package util

import "strings"

func ParseCommandLineArgs(argument string) (string, string) {
	if argument == "" {
		return "", ""
	}
	segments := strings.Split(argument, " ")
	commandArgIndex := len(segments) - 1 // this should always be last segment in array

	// return a tuple {cron expression, command}
	return strings.Join(segments[:commandArgIndex], " "), segments[commandArgIndex]
}
