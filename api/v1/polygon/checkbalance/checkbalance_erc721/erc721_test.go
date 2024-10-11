package checkbalance_erc721

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

	t.Run("Fetch user balance for ERC721", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		req := CheckErc721BalanceWithSalt{
			WalletAddress:   "0x975362c36b6842d48d02DBD3A077745Fc1C64175",
			ContractAddress: "0x975362c36b6842d48d02DBD3A077745Fc1C64175",
		}
		body, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}

		httpReq, err := http.NewRequest("POST", "", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		c.Request = httpReq
		erc721CheckBalanceWithSalt(c)
		assert.Equal(t, 200, rr.Result().StatusCode)
	})
}
