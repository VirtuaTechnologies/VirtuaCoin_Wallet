package verifysignature

import (
	"fmt"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/bsc"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/verify-signature")
	{
		g.POST("", verifySignature)
	}
}

func verifySignature(c *gin.Context) {
	network := "bsc"
	var req VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		err := fmt.Errorf("body is invalid: %w", err)
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).SendD(c)
		return
	}

	res, err := bsc.VerifySignature(req.Mnemonic, req.Message, req.Signature)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to verify signature").SendD(c)
		logo.Errorf("failed to verify signature from wallet of userId: %v and network: %v, error: %s",
			req.WalletAddress, network, err)
		return
	}

	if !res {
		httpo.NewErrorResponse(httpo.SignatureDenied, "signature is invalid").Send(c, 403)
		return
	}
	httpo.NewSuccessResponse(200, "signature is valid", nil).SendD(c)
}
