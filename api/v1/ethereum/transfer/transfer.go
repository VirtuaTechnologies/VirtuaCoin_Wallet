package transfer

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/ethereum/transfer/transfer_erc20"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/transfer")
	{
		transfer_erc20.ApplyRoutes(g)

	}
}
