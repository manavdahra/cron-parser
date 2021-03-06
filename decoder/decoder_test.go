package decoder

import (
	"deliveroo-cron/model"
	"deliveroo-cron/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type TestCase struct {
	Input  []string
	Output []string
}

func TestDecode(t *testing.T) {
	testCases := []TestCase{
		{
			Input:  []string{"*/15", "0", "1,15", "*", "1-5", "command"},
			Output: []string{"0,15,30,45", "0", "1,15", "1,2,3,4,5,6,7,8,9,10,11,12", "1,2,3,4,5", "command"},
		},
		{
			Input:  []string{"*/15", "0", "1,15", "*", "*", "command"},
			Output: []string{"0,15,30,45", "0", "1,15", "1,2,3,4,5,6,7,8,9,10,11,12", "0,1,2,3,4,5,6", "command"},
		},
		{
			Input:  []string{"*/15", "0", "*/10", "*", "1-5", "command"},
			Output: []string{"0,15,30,45", "0", "10,20,30", "1,2,3,4,5,6,7,8,9,10,11,12", "1,2,3,4,5", "command"},
		},
		{
			Input:  []string{"30,40,59", "0", "1,15", "*", "1-5", "command"},
			Output: []string{"30,40,59", "0", "1,15", "1,2,3,4,5,6,7,8,9,10,11,12", "1,2,3,4,5", "command"},
		},
		{
			Input:  []string{"*/15", "0", "1,15", "*/3", "1-5", "command"},
			Output: []string{"0,15,30,45", "0", "1,15", "3,6,9,12", "1,2,3,4,5", "command"},
		},
	}

	for _, testCase := range testCases {
		validate(t, testCase)
	}

}

func validate(t *testing.T, testCase TestCase) {
	input := testCase.Input
	output := testCase.Output

	cron := model.NewCron(input[5])
	err := Decode(strings.Join(input[0:5], " "), cron)
	assert.NoError(t, err)

	assert.Equal(t, fmt.Sprintf(util.GenerateOutputFmt(),
		output[0], // minute
		output[1], // hour
		output[2], // day of month
		output[3], // month
		output[4], // day of week
		output[5], // command
	), cron.String())
}
