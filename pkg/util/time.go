package util

import "google.golang.org/protobuf/types/known/timestamppb"

func TimestamppbToUnixMill(t *timestamppb.Timestamp) *int64 {
	if t == nil {
		return nil
	}
	unixMilli := t.AsTime().UnixMilli()
	return &unixMilli
}
