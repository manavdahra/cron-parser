package model

import (
	"deliveroo-cron/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCron_String(t *testing.T) {
	cron := NewCron("")
	assert.Equal(t, fmt.Sprintf(util.GenerateOutputFmt(), "", "", "", "", "", ""), cron.String())
}
