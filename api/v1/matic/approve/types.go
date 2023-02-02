package approve

type ApproveRequest struct {
	UserId          string `json:"userId" binding:"required"`
	ToAddress       string `json:"toAddress" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	TokenId         int64  `json:"tokenId" binding:"required"`
}

type TransferPayload struct {
	TrasactionHash string
}

type ApproveWithSalt struct {
	UserId          string `json:"userId" binding:"required"`
	Mnemonic        string `json:"mnemonic" binding:"required"`
	ToAddress       string `json:"toAddress" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	TokenId         int64  `json:"tokenId" binding:"required"`
	Salt            string `json:"salt" binding:"required"`
}
