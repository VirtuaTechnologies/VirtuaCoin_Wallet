package approve

type ApproveWithSalt struct {
	WalletAddress   string `json:"walletAddress" binding:"required"`
	Mnemonic        string `json:"mnemonic" binding:"required"`
	ToAddress       string `json:"toAddress" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	TokenId         int64  `json:"tokenId" binding:"required"`
	Salt            string `json:"salt" binding:"required"`
}

type ApprovePayload struct {
	TrasactionHash string
}
