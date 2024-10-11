package virtuacoinnft

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/virtuacoinnft/burn"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/virtuacoinnft/delegateassetcreation"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/virtuacoinnft/grantrole"
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
