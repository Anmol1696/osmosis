package stableswap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Anmol1696/osmosis/v11/x/gamm/types"
)

func (params PoolParams) Validate() error {
	if params.ExitFee.IsNegative() {
		return types.ErrNegativeExitFee
	}

	if params.ExitFee.GTE(sdk.OneDec()) {
		return types.ErrTooMuchExitFee
	}

	if params.SwapFee.IsNegative() {
		return types.ErrNegativeSwapFee
	}

	if params.SwapFee.GTE(sdk.OneDec()) {
		return types.ErrTooMuchSwapFee
	}
	return nil
}
