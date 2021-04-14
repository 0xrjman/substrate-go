package chainx

import "github.com/rjmand/go-substrate-rpc-client/v2/types"

///Chainx Type
type AssetId types.U32

//type AssetId uint32

const (
	Usable = iota
	Locked
	Reserved
	ReservedWithdrawal
	ReservedDexSpot
)