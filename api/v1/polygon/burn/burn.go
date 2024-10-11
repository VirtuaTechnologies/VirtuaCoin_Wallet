package burn

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/polygon/burn/burn_erc721"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/burn")
	{
		burn_erc721.ApplyRoutes(g)
	}
}
