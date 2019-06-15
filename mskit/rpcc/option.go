package rpcc

import (
	"time"

	"github.com/micro/go-micro/client"
)

func Timeout(d time.Duration) client.CallOption {
	return func(options *client.CallOptions) {
		options.RequestTimeout = d
	}
}
