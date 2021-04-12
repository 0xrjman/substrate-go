package chainx

import (
	"github.com/rjman-self/go-polkadot-rpc-client/expand/base"
	"github.com/rjmand/go-substrate-rpc-client/v2/types"
)

/// Polkadot MultiSignExtrinsic Type
var AsMultiNew = "as_multi_new"
var AsMultiApprove = "as_multi_approve"
var AsMultiExecuted = "as_multi_executed"
var AsMultiCancelled = "as_multi_cancelled"
var UtilityBatch = "multi_sign_batch"

type ChainXEventRecords struct {
	types.EventRecords
	Claims_Claimed                    []EventClaimsClaimed
	ElectionsPhragmen_VoterReported   []EventElectionsPhragmenVoterReported
	ElectionsPhragmen_MemberRenounced []EventElectionsPhragmenMemberRenounced
	ElectionsPhragmen_MemberKicked    []EventElectionsPhragmenMemberKicked
	ElectionsPhragmen_ElectionError   []EventElectionsPhragmenElectionError
	ElectionsPhragmen_EmptyTerm       []EventElectionsPhragmenEmptyTerm
	//ElectionsPhragmen_NewTerm		[]EventElectionsPhragmenNewTerm		暂不支持解析
	Democracy_Blacklisted []EventDemocracyBlacklisted

	XTransactionFee_FeePaid []EventXTransactionFeeFeePaid
	XAssets_Moved           []EventXAssetsMoved
}

func (p ChainXEventRecords) GetMultisigNewMultisig() []types.EventMultisigNewMultisig {
	return p.Multisig_NewMultisig
}

func (p ChainXEventRecords) GetMultisigApproval() []types.EventMultisigApproval {
	return p.Multisig_MultisigApproval
}

func (p ChainXEventRecords) GetMultisigExecuted() []types.EventMultisigExecuted {
	return p.Multisig_MultisigExecuted
}

func (p ChainXEventRecords) GetMultisigCancelled() []types.EventMultisigCancelled {
	return p.Multisig_MultisigCancelled
}

func (p ChainXEventRecords) GetUtilityBatchCompleted() []types.EventUtilityBatchCompleted {
	return p.Utility_BatchCompleted
}

func (p ChainXEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return p.Balances_Transfer
}

func (p ChainXEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return p.System_ExtrinsicSuccess
}

func (p ChainXEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return p.System_ExtrinsicFailed
}

type EventDemocracyBlacklisted struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

//type EventElectionsPhragmenNewTerm struct {
//	Phase    types.Phase
//	Vec
//	Topics []types.Hash
//}
type EventElectionsPhragmenEmptyTerm struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventElectionsPhragmenElectionError struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventElectionsPhragmenMemberKicked struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenMemberRenounced struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenVoterReported struct {
	Phase  types.Phase
	Who1   types.AccountID
	Who2   types.AccountID
	Bool   types.Bool
	Topics []types.Hash
}
type EventClaimsClaimed struct {
	Phase           types.Phase
	AccountId       types.AccountID
	EthereumAddress base.VecU8L20
	Balance         types.U128
	Topics          []types.Hash
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

// EventBalancesTransfer is emitted when a transfer succeeded (from, to, value)
type EventXAssetsMoved struct {
	Phase    types.Phase
	AssetId  AssetId
	From     types.AccountID
	FromType uint32
	To       types.AccountID
	ToType   uint32
	Balance  types.U128
	Topics   []types.Hash
}
