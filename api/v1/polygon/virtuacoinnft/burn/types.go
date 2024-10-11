package burn

type BurnWithSalt struct {
	WalletAddress   string `json:"walletAddress" binding:"required"`
	Mnemonic        string `json:"mnemonic" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	TokenId         int64  `json:"tokenId" binding:"required"`
	Salt            string `json:"salt" binding:"required"`
}

type BurnPayload struct {
	TrasactionHash string
}
