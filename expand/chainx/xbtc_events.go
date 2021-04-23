package chainx

import (
	"github.com/rjman-self/substrate-go/expand/chainx/xevents"
	"github.com/rjman-self/substrate-go/expand/chainx/xevents/xgateway"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

type XBtcV1 struct {
	XTransactionFee_FeePaid []EventXTransactionFeeFeePaid
	xevents.XAssets
	xevents.XStaking
	xevents.XMiningAsset
	xgateway.XGateWay
	xevents.XSystem
}

type XBtcV2 struct {
	XTransactionFee_FeePaid []EventXTransactionFeeFeePaid
	///TODO: XBtcV2
}

// EventXTransactionFeeFeePaid is emitted when some XTransactionFee was Paid
type EventXTransactionFeeFeePaid struct {
	Phase        types.Phase
	Author       types.AccountID
	AuthorFee    types.U128
	RewardPot    types.AccountID
	RewardPotFee types.U128
	Topics       []types.Hash
}
