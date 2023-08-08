package delegateassetcreation

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
	g := r.Group("/delegateassetcreation")
	{
		g.POST("", delegateAssetCreation)
	}
}

func delegateAssetCreation(c *gin.Context) {
	network := "bsc"
	var req DelegateAssetCreationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}

	erc721ContractAddr := common.HexToAddress(req.ContractAddress)
	var hash string
	hash, err := virtuacoin.DelegateAssetCreation(req.WalletAddress, erc721ContractAddr, req.MetadataURI)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to delegateAssetCreation").SendD(c)
		logo.Errorf("failed to delegateAssetCreation of erc721 to wallet Address: %v , network: %v, contractAddr: %v, error: %s", req.WalletAddress, network, req.ContractAddress, err)
		return
	}
	payload := DelegateErc721Payload{
		TrasactionHash: hash,
	}

	httpo.NewSuccessResponse(200, "delegateAssetCreation trasaction initiated", payload).SendD(c)
}
