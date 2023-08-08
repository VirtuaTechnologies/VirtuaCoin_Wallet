package transfer_erc20

type TransferPayload struct {
	TrasactionHash string
}

type TransferRequestSalt struct {
	WalletAddress   string  `json:"walletAddress" binding:"required"`
	Mnemonic        string  `json:"mnemonic" binding:"required"`
	To              string  `json:"to" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
	ContractAddress string  `json:"contractAddress" binding:"required"`
	Salt            string  `json:"salt" binding:"required"`
}
