package matic

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/approve"
	approveall "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/approveAll"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/checkbalance"
	signmessage "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/signMessage"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/transfer"
	verifysignature "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/verifySignature"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/virtuacoinnft"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/matic")
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
