package checkbalance_erc20

type CheckErc20BalanceRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
	ContractAddr  string `json:"contractAddr" binding:"required"`
}

type CheckErc20BalancePayload struct {
	Balance string `json:"balance"`
}
