package tools

import (
	"context"
	"time"
)

func Timeout(second ...int64) context.Context {
	var (
		t int64
	)
	if len(second) == 0 {
		t = 60 * 30
	} else {
		t = second[0]
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	return ctx
}
