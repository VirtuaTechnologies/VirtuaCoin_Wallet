package approveall

type ApproveAllRequest struct {
	UserId          string `json:"userId" binding:"required"`
	OperatorAddress string `json:"operatorAddress" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	Approved        bool   `json:"approved" binding:"required"`
}

type ApproveAllWithSalt struct {
	WalletAddress   string `json:"walletAddress" binding:"required"`
	OperatorAddress string `json:"operatorAddress" binding:"required"`
	Mnemonic        string `json:"mnemonic" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	Approved        bool   `json:"approved" binding:"required"`
}

type TransferPayload struct {
	TrasactionHash string
}
