package transfer_native

import (
	"errors"
	"math"
	"math/big"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/models/user"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/native")
	{

		g.POST("", nativeTransferWithSalt)
	}
}

func nativeTransfer(c *gin.Context) {
	network := "matic"
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}
	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(httpo.UserNotFound, "user not found").Send(c, 404)

			return
		}

		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch user").SendD(c)
		logo.Errorf("failed to fetch user mnemonic for userId: %v, error: %s",
			req.UserId, err)
		return
	}

	var hash string
	hash, err = polygon.Transfer(mnemonic, common.HexToAddress(req.To), *big.NewInt(req.Amount))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to tranfer").SendD(c)
		logo.Errorf("failed to tranfer to: %v from wallet of userId: %v and network: %v, error: %s", req.To,
			req.UserId, network, err)
		return
	}
	sendSuccessResponse(c, hash, req.UserId)
}

func sendSuccessResponse(c *gin.Context, hash string, userId string) {
	payload := TransferPayload{
		TrasactionHash: hash,
	}
	if err := user.AddTrasactionHash(userId, hash); err != nil {
		logo.Errorf("failed to add transaction hash: %v to user id: %v, error: %s", hash, userId, err)
	}
	httpo.NewSuccessResponse(200, "trasaction initiated", payload).SendD(c)
}

func nativeTransferWithSalt(c *gin.Context) {
	network := "matic"
	var req TransferRequestSalt
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}

	// Convert float64 to bigInt
	var amountInBigInt = big.NewInt(int64(req.Amount * math.Pow(10, 18)))

	var hash string
	hash, err := polygon.Transfer(req.Mnemonic, common.HexToAddress(req.To), *amountInBigInt)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to tranfer").SendD(c)
		logo.Errorf("failed to tranfer to: %v from wallet: %v and network: %v, error: %s", req.To, req.WalletAddress, network, err)
		return
	}
	payload := TransferPayload{
		TrasactionHash: hash,
	}
	httpo.NewSuccessResponse(200, "trasaction initiated", payload).SendD(c)
}
