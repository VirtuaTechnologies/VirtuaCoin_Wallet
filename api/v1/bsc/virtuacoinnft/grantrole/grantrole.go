package grantrole

import (
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	virtuacoin "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/bsc/virtuacoinnft"
	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/grantrole")
	{
		g.POST("", grantRole)
	}
}

func grantRole(c *gin.Context) {
	network := "bsc"
	var req GrantRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}

	erc721ContractAddress := common.HexToAddress(req.ContractAddress)

	var hash string
	hash, err := virtuacoin.GrantRole(req.RoleId, req.WalletAddress, erc721ContractAddress)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to grantRole").SendD(c)
		logo.Errorf("failed to grantRole of %v in VirtuaCoinNFT to wallet Address: %v , network: %v, contractAddr: %v, error: %s", req.RoleId, req.WalletAddress, network, req.ContractAddress, err)
		return
	}
	payload := GrantRolePayload{
		TrasactionHash: hash,
	}

	httpo.NewSuccessResponse(200, "grantRole trasaction initiated", payload).SendD(c)
}
