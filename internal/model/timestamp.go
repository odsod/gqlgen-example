package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func MarshalGoogleProtobufTimestamp(v timestamp.Timestamp) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		t, _ := ptypes.Timestamp(&v)
		_, _ = io.WriteString(w, strconv.Quote(t.Format(time.RFC3339Nano)))
	})
}

func UnmarshalGoogleProtobufTimestamp(v interface{}) (timestamp.Timestamp, error) {
	str, ok := v.(string)
	if !ok {
		return timestamp.Timestamp{}, fmt.Errorf("unmarshal google.protobuf.Timestamp: not a string: %v", v)
	}
	t, err := time.Parse(time.RFC3339Nano, str)
	if err != nil {
		return timestamp.Timestamp{}, fmt.Errorf("unmarshal google.protobuf.Timestamp: %w", err)
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return timestamp.Timestamp{}, fmt.Errorf("unmarshal google.protobuf.Timestamp: %w", err)
	}
	return *ts, nil
}
