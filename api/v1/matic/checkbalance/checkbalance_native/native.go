package checkbalance_native

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/polygon"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/native")
	{
		g.GET("/:walletAddress", nativeCheckBalanceWithSalt)
	}
}

func nativeCheckBalanceWithSalt(c *gin.Context) {
	paramWalletAddress := c.Param("walletAddress")
	if len(paramWalletAddress) <= 0 {
		httpo.NewErrorResponse(http.StatusBadRequest, "valid wallet address is required").SendD(c)
		return
	}
	network := "matic"

	var balance *big.Float
	balance, err := polygon.GetBalanceInDecimalsFromWalletAddress(paramWalletAddress)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get balance").SendD(c)
		logo.Errorf("failed to get balance from wallet of userId: %v and network: %v, error: %s", paramWalletAddress, network, err)
		return
	}

	balanceInFloat, _ := balance.Float64()
	balanceInString := strconv.FormatFloat(balanceInFloat, 'f', -1, 64)

	payload := CheckNativeBalancePayload{
		Balance: balanceInString,
	}
	httpo.NewSuccessResponse(200, "balance successfully fetched", payload).SendD(c)
}
