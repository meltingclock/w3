package txpool_test

import (
	"testing"

	"github.com/meltingclock/w3/rpctest"
	"github.com/meltingclock/w3/module/txpool"
)

func TestStatus(t *testing.T) {
	tests := []rpctest.TestCase[txpool.StatusResponse]{
		{
			Golden:  "status",
			Call:    txpool.Status(),
			WantRet: &txpool.StatusResponse{Pending: 10, Queued: 7},
		},
	}

	rpctest.RunTestCases(t, tests)
}
