package signmessage

import (
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/polygon"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/sign-message")
	{

		g.POST("", signMessage)
	}
}

func signMessage(c *gin.Context) {
	var req SignMessageRequestWithSalt
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)

		return
	}

	signature, err := polygon.SignMessage(req.Mnemonic, req.Message)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to sign").SendD(c)
		logo.Errorf("failed to sign with walletAddress: %s", err)
		return
	}

	sendSuccessResponse(c, signature, req.WalletAddress)
}

func sendSuccessResponse(c *gin.Context, signature string, userId string) {
	payload := SignMessagePayload{
		Signature: signature,
	}
	httpo.NewSuccessResponse(200, "signature generated", payload).SendD(c)
}
