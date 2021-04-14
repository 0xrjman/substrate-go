package chainx

import (
	"github.com/rjmand/go-substrate-rpc-client/v2/types"
)

type AssetId uint32

type XGatewayBitcoin struct {
	XGatewayBitcoin_HeaderInserted              []EventXGatewayBitcoinHeaderInserted
	XGatewayBitcoin_TxProcessed                 []EventXGatewayBitcoinTxProcessed
	XGatewayBitcoin_Deposited                   []EventXGatewayBitcoinDeposited
	XGatewayBitcoin_Withdrawn                   []EventXGatewayBitcoinWithdrawn
	XGatewayBitcoin_UnclaimedDeposit            []EventXGatewayBitcoinUnclaimedDeposit
	XGatewayBitcoin_PendingDepositRemoved       []EventXGatewayBitcoinPendingDepositRemoved
	XGatewayBitcoin_WithdrawalProposalCreated   []EventXGatewayBitcoinWithdrawalProposalCreated
	XGatewayBitcoin_WithdrawalProposalVoted     []EventXGatewayBitcoinWithdrawalProposalVoted
	XGatewayBitcoin_WithdrawalProposalDropped   []EventXGatewayBitcoinWithdrawalProposalDropped
	XGatewayBitcoin_WithdrawalProposalCompleted []EventXGatewayBitcoinWithdrawalProposalCompleted
	XGatewayBitcoin_WithdrawalFatalErr          []EventXGatewayBitcoinWithdrawalFatalErr
}

type BtcTxType uint

const (
	Withdrawal        BtcTxType = 0
	Deposit           BtcTxType = 1
	HotAndCold        BtcTxType = 2
	TrusteeTransition BtcTxType = 3
	Irrelevance       BtcTxType = 4
)

type BtcTxResult uint

const (
	Success BtcTxResult = 0
	Failure BtcTxResult = 1
)

type BtcTxState struct {
	tx_type BtcTxType
	result  BtcTxResult
}

type EventXGatewayBitcoinHeaderInserted struct {
	Phase      types.Phase
	HeaderHash types.H256
	Topics     []types.Hash
}

type EventXGatewayBitcoinTxProcessed struct {
	Phase     types.Phase
	TxHash    types.H256
	BlockHash types.H256
	TxState   BtcTxState
	Topics    []types.Hash
}

type EventXGatewayBitcoinDeposited struct {
	Phase  types.Phase
	TxHash types.H256
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventXGatewayBitcoinWithdrawn struct {
	Phase  types.Phase
	TxHash types.H256
	Ids    []uint32
	Total  types.U128
	Topics []types.Hash
}

type EventXGatewayBitcoinUnclaimedDeposit struct {
	Phase      types.Phase
	TxHash     types.H256
	BtcAddress []uint8
	Topics     []types.Hash
}

type EventXGatewayBitcoinPendingDepositRemoved struct {
	Phase      types.Phase
	Depositor  types.AccountID
	Amount     types.U128
	TxHash     types.H256
	BtcAddress []uint8
	Topics     []types.Hash
}

type EventXGatewayBitcoinWithdrawalProposalCreated struct {
	Phase    types.Phase
	Proposer types.AccountID
	Ids      []uint32
	Topics   []types.Hash
}

type EventXGatewayBitcoinWithdrawalProposalVoted struct {
	Phase      types.Phase
	Trustee    types.AccountID
	VoteStatus bool
	Topics     []types.Hash
}

type EventXGatewayBitcoinWithdrawalProposalDropped struct {
	Phase       types.Phase
	RejectCount uint32
	TotalCount  uint32
	Ids         []uint32
	Topics      []types.Hash
}

type EventXGatewayBitcoinWithdrawalProposalCompleted struct {
	Phase  types.Phase
	TxHash types.H256
	Topics []types.Hash
}

type EventXGatewayBitcoinWithdrawalFatalErr struct {
	Phase        types.Phase
	TxHash       types.H256
	ProposalHash types.H256
	Topics       []types.Hash
}
