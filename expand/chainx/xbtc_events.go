package chainx

type XBtcV1 struct {
	XTransactionFee_FeePaid 			[]EventXTransactionFeeFeePaid
	XGatewayBitcoin
	XAssets
	XMiningAsset
	XStaking
}

type XBtcV2 struct {
	XTransactionFee_FeePaid 			[]EventXTransactionFeeFeePaid
	///TODO: XBtcV2
	XAssets
	XMiningAsset
	XStaking
}