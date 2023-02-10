package virtuacoinnft

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/virtuacoinnft/burn"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/virtuacoinnft/delegateassetcreation"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/virtuacoinnft/grantrole"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/virtuacoinnft")
	{
		delegateassetcreation.ApplyRoutes(g)
		grantrole.ApplyRoutes(g)
		burn.ApplyRoutes(g)
	}
}
