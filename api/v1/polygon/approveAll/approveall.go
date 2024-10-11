package approveall

import (
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/approve-all")
	{
		g.POST("", approveAllWithSalt)
	}
}

func sendSuccessResponse(c *gin.Context, hash string, userId string) {
	payload := TransferPayload{
		TrasactionHash: hash,
	}
	httpo.NewSuccessResponse(200, "trasaction initiated", payload).SendD(c)
}

func approveAllWithSalt(c *gin.Context) {
	network := "pollygon"
	var req ApproveAllWithSalt
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}

	hash, err := polygon.SetAprovalForAllErc721(req.Mnemonic, common.HexToAddress(req.OperatorAddress), common.HexToAddress(req.ContractAddress), req.Approved)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to approve all").SendD(c)
		logo.Errorf("failed to approve all to operator: %v from wallet of userId: %v, network: %v, contractAddr: %v, error: %s", req.OperatorAddress,
			req.WalletAddress, network, req.ContractAddress, err)
		return
	}
	sendSuccessResponse(c, hash, req.WalletAddress)
}
