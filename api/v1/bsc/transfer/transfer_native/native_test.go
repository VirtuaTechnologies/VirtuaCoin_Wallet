package transfer_native

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

func Test_Transfer(t *testing.T) {
	envconfig.InitEnvVars()

	appinit.Init()
	gin.SetMode(gin.TestMode)

	t.Run("Native token", func(t *testing.T) {
		rr := httptest.NewRecorder()

		req := TransferRequestSalt{
			WalletAddress: "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
			Mnemonic:      "test test test test test test",
			To:            "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
			Amount:        1,
		}
		d, e := json.Marshal(req)
		if e != nil {
			t.Fatal(e)
		}
		c, _ := gin.CreateTestContext(rr)
		httpReq, e := http.NewRequest("GET", "/", bytes.NewBuffer(d))
		if e != nil {
			t.Fatal(e)
		}
		c.Request = httpReq
		nativeTransferWithSalt(c)

		assert.Equal(t, 200, rr.Result().StatusCode)

	})

}
