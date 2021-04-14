package chainx

import (
	"github.com/rjmand/go-substrate-rpc-client/v2/types"
)

/// XAssets Type
type AssetId types.U32
type AssetType uint

///XMining Type
type SessionIndex uint32

const(
	Usable				AssetType = 0
	Locked				AssetType = 1
	Reserved			AssetType = 2
	ReservedWithdrawal  AssetType = 3
	ReservedDexSpot		AssetType = 4
)

/// Some balances of an asset was moved from one to another. [asset_id, from, from_type, to, to_type, amount]
type EventXAssetsMoved struct {
	Phase    		types.Phase
	AssetId  		AssetId
	From     		types.AccountID
	FromType 		AssetType
	To       		types.AccountID
	ToType   		AssetType
	Balance  		types.U128
	Topics   		[]types.Hash
}

/// New balances of an asset were issued. [asset_id, receiver, amount]
type EventXAssetsIssued struct {
	Phase    		types.Phase
	AssetId  		AssetId
	Receiver 		types.AccountID
	Amount   		types.U128
	Topics   		[]types.Hash
}

/// Some balances of an asset were destoryed. [asset_id, who, amount]
type EventXAssetsDestroyed struct {
	Phase    		types.Phase
	AssetId  		AssetId
	Who 			types.AccountID
	Amount   		types.U128
	Topics   		[]types.Hash
}

/// Set asset balance of an account by root. [asset_id, who, asset_type, amount]
type EventXAssetsBalanceSet struct {
	Phase    		types.Phase
	AssetId  		AssetId
	Who 			types.AccountID
	AssetType       AssetType
	Balance   		types.U128
	Topics   		[]types.Hash
}

/// An asset miner claimed the mining reward. [claimer, asset_id, amount]
type EventXMiningAssetClaimed struct {
	Phase    		types.Phase
	Claimer 		types.AccountID
	AssetId  		AssetId
	Amount   		types.U128
	Topics   		[]types.Hash
}

/// Issue new balance to the reward pot. [reward_pot_account, amount]
type EventXMiningAssetMinted struct {
	Phase    				types.Phase
	RewardPotAccount 		types.AccountID
	Amount   				types.U128
	Topics   				[]types.Hash
}

/// Issue new balance to this account. [account, reward_amount]
type EventXStakingMinted struct {
	Phase    				types.Phase
	Account 				types.AccountID
	RewardAmount   			types.U128
	Topics   				[]types.Hash
}

/// A validator (and its reward pot) was slashed. [validator, slashed_amount]
type EventXStakingSlashed struct {
	Phase    				types.Phase
	Validator 				types.AccountID
	SlashedAmount   		types.U128
	Topics   				[]types.Hash
}

/// A nominator bonded to the validator this amount. [nominator, validator, amount]
type EventXStakingBonded struct {
	Phase    				types.Phase
	Nominator 				types.AccountID
	Validator 				types.AccountID
	Amount   		types.U128
	Topics   				[]types.Hash
}

/// A nominator switched the vote from one validator to another. [nominator, from, to, amount]
type EventXStakingRebonded struct {
	Phase    				types.Phase
	Nominator 				types.AccountID
	From 					types.AccountID
	To 						types.AccountID
	Amount   				types.U128
	Topics   				[]types.Hash
}

/// A nominator unbonded this amount. [nominator, validator, amount]
type EventXStakingUnbonded struct {
	Phase    				types.Phase
	Nominator 				types.AccountID
	Validator 				types.AccountID
	Amount   				types.U128
	Topics   				[]types.Hash
}

/// A nominator claimed the staking dividend. [nominator, validator, dividend]
type EventXStakingClaimed struct {
	Phase    				types.Phase
	Nominator 				types.AccountID
	Validator 				types.AccountID
	Dividend   				types.U128
	Topics   				[]types.Hash
}

/// The nominator withdrew the locked balance from the unlocking queue. [nominator, amount]
type EventXStakingWithdrawn struct {
	Phase    				types.Phase
	Nominator 				types.AccountID
	Amount   				types.U128
	Topics   				[]types.Hash
}

/// Offenders were forcibly to be chilled due to insufficient reward pot balance. [session_index, chilled_validators]
type EventXStakingForceChilled struct {
	Phase    				types.Phase
	SessionIndex 			SessionIndex
	ChilledValidators       []types.AccountID
	Topics   				[]types.Hash
}

/// Unlock the unbonded withdrawal by force. [account]
type EventXStakingForceAllWithdrawn struct {
	Phase    				types.Phase
	Account 				types.AccountID
	Topics   				[]types.Hash
}