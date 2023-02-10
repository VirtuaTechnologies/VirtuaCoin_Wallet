package transfer_erc721

import (
	"math/big"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/erc721")
	{
		g.POST("", transferWithSalt)
	}
}

func transferWithSalt(c *gin.Context) {
	network := "matic"
	var req TransferRequestSalt
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logo.Errorf("Invalid request body: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, " Invalid body").SendD(c)
		return
	}
	hash, err := polygon.TransferERC721(req.Mnemonic, common.HexToAddress(req.To), common.HexToAddress(req.ContractAddress), *big.NewInt(int64(req.TokenId)))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to tranfer").SendD(c)
		logo.Errorf("failed to tranfer to: %v from wallet: %v , network: %v, contractAddr: %v , error: %s", req.To, req.WalletAddress, network, req.ContractAddress, err)
		return
	}

	payload := TransferPayload{TrasactionHash: hash}
	httpo.NewSuccessResponse(http.StatusOK, "trasaction initiated", payload).SendD(c)
}
