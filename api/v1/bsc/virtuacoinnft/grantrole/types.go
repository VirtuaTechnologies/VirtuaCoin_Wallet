package grantrole

type GrantRoleRequest struct {
	RoleId          string `json:"roleId" binding:"required"`
	WalletAddress   string `json:"walletAddress" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
}

type GrantRolePayload struct {
	TrasactionHash string
}
