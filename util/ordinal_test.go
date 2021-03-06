package util

import (
	"deliveroo-cron/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceWithOrdinal(t *testing.T) {
	ordinal := "3"
	assert.Equal(t, ordinal, ReplaceWithOrdinal("APR", common.Months))
	assert.Equal(t, ordinal, ReplaceWithOrdinal("apr", common.Months))

	ordinal = "9"
	assert.Equal(t, ordinal, ReplaceWithOrdinal("OCT", common.Months))
	assert.Equal(t, ordinal, ReplaceWithOrdinal("oct", common.Months))

	ordinal = "4"
	assert.Equal(t, ordinal, ReplaceWithOrdinal("THU", common.DaysOfWeek))
	assert.Equal(t, ordinal, ReplaceWithOrdinal("thu", common.DaysOfWeek))
}
