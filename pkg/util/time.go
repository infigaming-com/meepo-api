package util

import (
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimestamppbToUnixMill(t *timestamppb.Timestamp) *int64 {
	if t == nil {
		return nil
	}
	unixMilli := t.AsTime().UnixMilli()
	return &unixMilli
}

// IsValidTimestampTimeRange validates if the given time range is valid.
// A valid time range means:
// 1. If either timestamp is nil, the range is considered valid
// 2. If both timestamps are present, start time must be before or equal to end time
func IsValidTimestampTimeRange(start, end *timestamppb.Timestamp) (bool, error) {
	// If either timestamp is nil, consider it valid
	if start == nil || end == nil {
		return true, nil
	}

	// Check if start time is before or equal to end time
	if start.AsTime().After(end.AsTime()) {
		return false, errors.New("start time must be before or equal to end time")
	}

	return true, nil
}

// IsValidInt64TimeRange validates if the given time range is valid.
// A valid time range means:
// 1. If either timestamp is nil, the range is considered valid
// 2. If both timestamps are present:
//   - They must be of the same type (both Unix or both UnixMilli)
//   - Start time must be before or equal to end time
func IsValidInt64TimeRange(start, end *int64) (bool, error) {
	// If either timestamp is nil, consider it valid
	if start == nil || end == nil {
		return true, nil
	}

	// Check if both timestamps are of the same type
	// Unix timestamps are typically in billions (e.g., 1700000000)
	// UnixMilli timestamps are typically in trillions (e.g., 1700000000000)
	const unixMilliThreshold int64 = 10000000000 // 10 billion
	startIsUnix := *start < unixMilliThreshold
	endIsUnix := *end < unixMilliThreshold
	if startIsUnix != endIsUnix {
		return false, errors.New("timestamps must be of the same type (both Unix or both UnixMilli)")
	}

	if *start > *end {
		return false, errors.New("start time must be before or equal to end time")
	}

	return true, nil
}
