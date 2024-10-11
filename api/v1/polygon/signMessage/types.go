package signmessage

type SignMessageRequestWithSalt struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
	Mnemonic      string `json:"mnemonic" binding:"required"`
	Message       string `json:"message" binding:"required"`
}

type SignMessagePayload struct {
	Signature string `json:"signature"`
}
