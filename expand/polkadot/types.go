package polkadot

import (
	"github.com/rjman-self/go-polkadot-rpc-client/expand/base"
	"github.com/rjmand/go-substrate-rpc-client/v2/types"
)

type PolkadotEventRecords struct {
	types.EventRecords
	Claims_Claimed                    []EventClaimsClaimed
	ElectionsPhragmen_VoterReported   []EventElectionsPhragmenVoterReported
	ElectionsPhragmen_MemberRenounced []EventElectionsPhragmenMemberRenounced
	ElectionsPhragmen_MemberKicked    []EventElectionsPhragmenMemberKicked
	ElectionsPhragmen_ElectionError   []EventElectionsPhragmenElectionError
	ElectionsPhragmen_EmptyTerm       []EventElectionsPhragmenEmptyTerm
	//ElectionsPhragmen_NewTerm		[]EventElectionsPhragmenNewTerm		暂不支持解析
	Democracy_Blacklisted []EventDemocracyBlacklisted
}

func (p PolkadotEventRecords) GetMultisigNewMultisig() []types.EventMultisigNewMultisig {
	return p.Multisig_NewMultisig
}

func (p PolkadotEventRecords) GetMultisigApproval() []types.EventMultisigApproval {
	return p.Multisig_MultisigApproval
}

func (p PolkadotEventRecords) GetMultisigExecuted() []types.EventMultisigExecuted {
	return p.Multisig_MultisigExecuted
}

func (p PolkadotEventRecords) GetMultisigCancelled() []types.EventMultisigCancelled {
	return p.Multisig_MultisigCancelled
}

func (p PolkadotEventRecords) GetUtilityBatchCompleted() []types.EventUtilityBatchCompleted {
	return p.Utility_BatchCompleted
}

func (p PolkadotEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return p.Balances_Transfer
}

func (p PolkadotEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return p.System_ExtrinsicSuccess
}

func (p PolkadotEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
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
