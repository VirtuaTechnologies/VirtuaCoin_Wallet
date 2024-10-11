package checkbalance_erc721

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
		g.POST("", erc721CheckBalanceWithSalt)
	}
}

func erc721CheckBalanceWithSalt(c *gin.Context) {
	var req CheckErc721BalanceWithSalt
	err := c.ShouldBindJSON(&req)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}
	network := "polygon"

	var balance *big.Int
	balance, err = polygon.GetERC721Balance(req.Mnemonic, common.HexToAddress(req.ContractAddress))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get ERC721 balance").SendD(c)
		logo.Errorf("failed to get ERC721 balance of wallet: %v , network: %v, contractAddress: %v , error: %s", req.WalletAddress, network, req.ContractAddress, err)
		return
	}

	payload := CheckErc721BalancePayload{
		Balance: balance.String(),
	}
	httpo.NewSuccessResponse(200, "balance successfully fetched", payload).SendD(c)
}
