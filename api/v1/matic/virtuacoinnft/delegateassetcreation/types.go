package delegateassetcreation

type DelegateAssetCreationRequest struct {
	WalletAddress   string `json:"walletAddress" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	MetadataURI     string `json:"metadataURI" binding:"required,min=1"`
}

type DelegateErc721Payload struct {
	TrasactionHash string
}
