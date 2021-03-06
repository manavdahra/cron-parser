package common

import "errors"

var ErrIntervalOutOfBounds = errors.New("interval out of bounds")

var ErrInvalidIntervalLimits = errors.New("invalid interval limits. start > end")
var ErrInvalidIntervalRanges = errors.New("multiple ranges specified for interval")
var ErrInvalidExpMultipleSteps = errors.New("multiple steps specified")
var ErrInvalidExpStepsNegative = errors.New("steps cannot be negative")
var ErrInvalidExpMismatchedSeg = errors.New("cron expression is invalid. Doesnt match required no of segments")

var ErrCronObjectNotDefined = errors.New("cron object is not defined")
