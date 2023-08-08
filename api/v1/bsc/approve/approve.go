package approve

import (
	"math/big"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/bsc"
	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/approve")
	{

		g.POST("", approveWithSalt)
	}
}

func sendSuccessResponse(c *gin.Context, hash string, userId string) {
	payload := ApprovePayload{
		TrasactionHash: hash,
	}
	httpo.NewSuccessResponse(200, "trasaction initiated", payload).SendD(c)
}

func approveWithSalt(c *gin.Context) {
	network := "bsc"
	var req ApproveWithSalt
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("Invalid request body:%s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, " Invalid body").SendD(c)
		return
	}

	hash, err := bsc.ApproveERC721(req.Mnemonic, common.HexToAddress(req.ToAddress), common.HexToAddress(req.ContractAddress), *big.NewInt(req.TokenId))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to approve").SendD(c)
		logo.Errorf("failed to approve to: %v from wallet of userId: %v, network: %v, contractAddr: %v, tokenId: %v, error: %s", req.ToAddress,
			req.WalletAddress, network, req.ContractAddress, req.TokenId, err)
		return
	}
	sendSuccessResponse(c, hash, req.WalletAddress)

}
