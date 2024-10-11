package apiv1

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/bsc"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/ethereum"
	polygon "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/wallet"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		// register.ApplyRoutes(v1)
		wallet.ApplyRoutes(v1)
		polygon.ApplyRoutes(v1)
		bsc.ApplyRoutes(v1)
		ethereum.ApplyRoutes(v1)
	}
}
