package polygon

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/approve"
	approveall "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/approveAll"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/checkbalance"
	signmessage "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/signMessage"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/transfer"
	verifysignature "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/verifySignature"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/virtuacoinnft"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/polygon")
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
