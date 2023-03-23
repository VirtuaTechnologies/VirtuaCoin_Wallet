package checkbalance_erc20

import (
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
		g.GET("/:contractAddress/:walletAddress", erc20CheckBalanceSalt)
	}
}

func erc20CheckBalanceSalt(c *gin.Context) {
	paramContractAddress := c.Param("contractAddress")
	paramWalletAddress := c.Param("walletAddress")
	if len(paramWalletAddress) <= 0 {
		httpo.NewErrorResponse(http.StatusBadRequest, "valid wallet address is required").SendD(c)
		return
	}
	network := "matic"

	balance, err := polygon.GetERC20BalanceInDecimalsFromWalletAddress(common.HexToAddress(paramWalletAddress), common.HexToAddress(paramContractAddress))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get balance").SendD(c)
		logo.Errorf("failed to get ERC20 balance of wallet: %v , network: %v, contractAddr: %v , error: %s", paramWalletAddress, network, paramContractAddress, err)
		return
	}

	balanceInFloat, _ := balance.Float64()
	balanceInString := strconv.FormatFloat(balanceInFloat, 'f', -1, 64)

	payload := CheckErc20BalancePayload{
		Balance: balanceInString,
	}
	httpo.NewSuccessResponse(200, "balance successfully fetched", payload).SendD(c)
}
