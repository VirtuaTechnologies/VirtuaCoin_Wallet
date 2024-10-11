package transfer_erc20

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/erc20")
	{

		g.POST("", transferWithSalt)
	}
}

func transferWithSalt(c *gin.Context) {
	network := "polygon"
	var req TransferRequestSalt
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}
	var hash string
	amount := new(big.Int)
	amountStr := strconv.FormatFloat(req.Amount, 'f', -1, 64)
	amount.SetString(amountStr, 10)

	hash, err := polygon.TransferERC20(req.Mnemonic, common.HexToAddress(req.To), common.HexToAddress(req.ContractAddress), amount)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to transfer").SendD(c)
		logo.Errorf("failed to transfer to: %v from wallet: %v , network: %v, contractAddr: %v , error: %s", req.To, req.WalletAddress, network, req.ContractAddress, err)
		return
	}
	payload := TransferPayload{
		TrasactionHash: hash,
	}
	httpo.NewSuccessResponse(200, "trasaction initiated", payload).SendD(c)
}
