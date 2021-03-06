package common

const (
	// TODO: make this more readable and descriptive
	HelpFmt = `
Command line arguments supplied are not valid:
Example format: "%s /usr/bin/find"
`
	CronExp = "*/15 0 1,15 * 1-5"
)

var Fields = []string{"minute", "hour", "day of month", "month", "day of week", "command"}
var Months = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
var DaysOfWeek = []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}