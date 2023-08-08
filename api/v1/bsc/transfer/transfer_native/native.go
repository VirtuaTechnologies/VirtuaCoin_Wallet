package transfer_native

import (
	"math"
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
	g := r.Group("/native")
	{

		g.POST("", nativeTransferWithSalt)
	}
}

func nativeTransferWithSalt(c *gin.Context) {
	network := "bsc"
	var req TransferRequestSalt
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}

	// Convert float64 to bigInt
	var amountInBigInt = big.NewInt(int64(req.Amount * math.Pow(10, 18)))

	var hash string
	hash, err := bsc.Transfer(req.Mnemonic, common.HexToAddress(req.To), *amountInBigInt)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to tranfer").SendD(c)
		logo.Errorf("failed to tranfer to: %v from wallet: %v and network: %v, error: %s", req.To, req.WalletAddress, network, err)
		return
	}
	payload := TransferPayload{
		TrasactionHash: hash,
	}
	httpo.NewSuccessResponse(200, "trasaction initiated", payload).SendD(c)
}
