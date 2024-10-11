package ethereum

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/ethereum/checkbalance"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/ethereum/transfer"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/ethereum")
	{
		checkbalance.ApplyRoutes(v1)

		transfer.ApplyRoutes(v1)

	}
}
