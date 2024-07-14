package eth_test

import (
	"math/big"
	"testing"

	"github.com/meltingclock/w3"
	"github.com/meltingclock/w3/module/eth"
	"github.com/meltingclock/w3/rpctest"
)

func TestBlockNumber(t *testing.T) {
	tests := []rpctest.TestCase[big.Int]{
		{
			Golden:  "block_number",
			Call:    eth.BlockNumber(),
			WantRet: w3.I("0xc0fe"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
