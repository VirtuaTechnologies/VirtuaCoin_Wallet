package bsc

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/bsc/approve"
	approveall "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/bsc/approveAll"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/bsc/checkbalance"
	signmessage "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/bsc/signMessage"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/bsc/transfer"
	verifysignature "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/bsc/verifySignature"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/bsc/virtuacoinnft"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/bsc")
	{
		checkbalance.ApplyRoutes(v1)
		verifysignature.ApplyRoutes(v1)

		signmessage.ApplyRoutes(v1)
		// burn.ApplyRoutes(v1)
		transfer.ApplyRoutes(v1)
		approve.ApplyRoutes(v1)
		approveall.ApplyRoutes(v1)
		virtuacoinnft.ApplyRoutes(v1)
	}
}
