package burn_erc721

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
	g := r.Group("/erc721")
	{
		g.POST("", burnWithSalt)
	}
}

func sendSuccessResponse(c *gin.Context, hash string, userId string) {
	payload := BurnPayload{
		TrasactionHash: hash,
	}
	httpo.NewSuccessResponse(200, "trasaction initiated", payload).SendD(c)
}

func burnWithSalt(c *gin.Context) {
	network := "bsc"
	var req BurnWithSalt
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("Invalid request body:%s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, " Invalid body").SendD(c)
		return
	}

	hash, err := bsc.BurnERC721(req.Mnemonic, common.HexToAddress(req.ContractAddress), *big.NewInt(req.TokenId))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to approve").SendD(c)
		logo.Errorf("failed to burn tokenId: %v from wallet of: %v, network: %v, contractAddr: %v, error: %s", req.TokenId, req.WalletAddress, network, req.ContractAddress, err)
		return
	}
	sendSuccessResponse(c, hash, req.WalletAddress)
}
