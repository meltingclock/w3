package eth_test

import (
	"math/big"
	"testing"

	"github.com/meltingclock/w3"
	"github.com/meltingclock/w3/module/eth"
	"github.com/meltingclock/w3/rpctest"
)

func TestGasTipCap(t *testing.T) {
	tests := []rpctest.TestCase[big.Int]{
		{
			Golden:  "gas_tip_cap",
			Call:    eth.GasTipCap(),
			WantRet: w3.I("0xc0fe"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
