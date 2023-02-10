package verifysignature

type VerifyRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
	Mnemonic      string `json:"mnemonic" binding:"required"`
	Message       string `json:"message" binding:"required"`
	Signature     string `json:"signature" binding:"required,hexadecimal,len=132"`
}
