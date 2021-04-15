package xevents

import (
	"github.com/rjmand/go-substrate-rpc-client/v2/types"
)

/// XSystem Type
type XSystem struct {
	XSystem_Blacklisted   []EventXSystemBlacklisted
	XSystem_Unblacklisted []EventXSystemUnblacklisted
}

/// An account was added to the blacklist. [who]
type EventXSystemBlacklisted struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

/// An account was removed from the blacklist. [who]
type EventXSystemUnblacklisted struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}
