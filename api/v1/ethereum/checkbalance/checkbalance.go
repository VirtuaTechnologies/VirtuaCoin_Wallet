package checkbalance

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/ethereum/checkbalance/checkbalance_erc20"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/checkbalance")
	{
		checkbalance_erc20.ApplyRoutes(g)

	}
}
