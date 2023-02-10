package checkbalance_erc721

type CheckErc721BalanceRequest struct {
	UserId       string `json:"userId" binding:"required"`
	ContractAddr string `json:"contractAddr" binding:"required"`
}

type CheckErc721BalanceWithSalt struct {
	WalletAddress   string `json:"walletAddress" binding:"required"`
	Mnemonic        string `json:"mnemonic" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
}

type CheckErc721BalancePayload struct {
	Balance string `json:"balance"`
}
