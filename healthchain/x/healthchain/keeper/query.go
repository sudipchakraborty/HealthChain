package keeper

import (
	"healthchain/x/healthchain/types"
)

var _ types.QueryServer = Keeper{}
