package checkbalance_erc20

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/app/stage/appinit"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/config/envconfig"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_CheckBalance(t *testing.T) {
	envconfig.InitEnvVars()

	appinit.Init()
	gin.SetMode(gin.TestMode)

	t.Run("Fetch user balance for ERC20", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		req := CheckErc20BalanceRequest{
			WalletAddress: "0x2d7882bedcbfddce29ba99965dd3cdf7fcb10a1e",
		}
		body, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}

		httpReq, err := http.NewRequest("POST", "/?erc20address=0x2d7882bedcbfddce29ba99965dd3cdf7fcb10a1e", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		c.Request = httpReq
		erc20CheckBalanceSalt(c)
		assert.Equal(t, 200, rr.Result().StatusCode)
	})
}
