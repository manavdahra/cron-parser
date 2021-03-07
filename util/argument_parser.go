package util

import "strings"

func ParseCommandLineArgs(argument string) (string, string) {
	if argument == "" {
		return "", ""
	}
	segments := strings.Split(argument, " ")
	commandArgIndex := len(segments) - 1
	return strings.Join(segments[:commandArgIndex], " "), segments[commandArgIndex]
}
