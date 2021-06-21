package parachain

import gsrpcTypes "github.com/centrifuge/go-substrate-rpc-client/v3/types"

type ParaEvents struct {
	// Fix some events
	Balances_ReserveRepatriated []gsrpcTypes.EventBalancesReserveRepatriated
	System_Remarked             []Remarked
	Scheduler_Dispatched        []SchedulerDispatched

	// slot auction
	Crowdloan_Created           []Created
	Crowdloan_Contributed       []Contributed
	Crowdloan_Withdrew          []Withdrew
	Crowdloan_PartiallyRefunded []PartiallyRefunded
	Crowdloan_AllRefunded       []PartiallyRefunded
	Crowdloan_Dissolved         []Dissolved
	Crowdloan_HandleBidResult   []HandleBidResult
	Crowdloan_Edited            []Edited
	Crowdloan_MemoUpdated       []MemoUpdated
	Crowdloan_AddedToNewRaise   []AddedToNewRaise

	Auctions_AuctionStarted     []AuctionStarted
	Auctions_AuctionClosed      []AuctionClosed
	Auctions_Reserved           []Reserved
	Auctions_Unreserved         []Unreserved
	Auctions_ReserveConfiscated []ReserveConfiscated
	Auctions_BidAccepted        []BidAccepted
	Auctions_WinningOffset      []WinningOffset

	Slots_NewLeasePeriod []NewLeasePeriod
	Slots_Leased         []Leased

	// rococo
	Inclusion_CandidateIncluded []CandidateIncluded
	Inclusion_CandidateBacked   []CandidateBacked

	// kusama
	ParasInclusion_CandidateIncluded []CandidateIncluded
	ParasInclusion_CandidateBacked   []CandidateBacked

	RandomnessCollectiveFlip_Proposed       []Proposed
	RandomnessCollectiveFlip_Voted          []Voted
	RandomnessCollectiveFlip_Approved       []Approved
	RandomnessCollectiveFlip_Disapproved    []Disapproved
	RandomnessCollectiveFlip_Executed       []Executed
	RandomnessCollectiveFlip_MemberExecuted []MemberExecuted
	RandomnessCollectiveFlip_Closed         []Closed

	PhragmenElection_NewTerm           []gsrpcTypes.EventElectionsNewTerm
	PhragmenElection_EmptyTerm         []gsrpcTypes.EventElectionsEmptyTerm
	PhragmenElection_ElectionError     []gsrpcTypes.EventElectionsElectionError
	PhragmenElection_MemberKicked      []gsrpcTypes.EventElectionsMemberKicked
	PhragmenElection_Renounced         []EventElectionsRenounced
	PhragmenElection_CandidateSlashed  []EventElectionsCandidateSlashed
	PhragmenElection_SeatHolderSlashed []EventElectionsSeatHolderSlashed

	Gilt_BidPlaced    []BidPlaced
	Gilt_BidRetracted []BidRetracted
	Gilt_GiltIssued   []GiltIssued
	Gilt_GiltThawed   []GiltThawed

	XcmPallet_Attempted []Attempted
	XcmPallet_Sent      []Sent

	Paras_CurrentCodeUpdated   []CurrentCodeUpdated
	Paras_CurrentHeadUpdated   []CurrentHeadUpdated
	Paras_CodeUpgradeScheduled []CodeUpgradeScheduled
	Paras_NewHeadNoted         []NewHeadNoted
	Paras_ActionQueued         []ActionQueued

	ParasUmp_InvalidFormat          []InvalidFormat
	ParasUmp_UnsupportedVersion     []UnsupportedVersion
	ParasUmp_ExecutedUpward         []ExecutedUpward
	ParasUmp_WeightExhausted        []WeightExhausted
	ParasUmp_UpwardMessagesReceived []UpwardMessagesReceived

	ParasHrmp_OpenChannelRequested []OpenChannelRequested
	ParasHrmp_OpenChannelAccepted  []OpenChannelAccepted
	ParasHrmp_ChannelClosed        []ChannelClosed

	Registrar_Registered   []Registered
	Registrar_Deregistered []Deregistered
	Registrar_Reserved     []ParaReserved
}
