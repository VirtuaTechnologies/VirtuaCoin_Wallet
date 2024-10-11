package approve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/app/stage/appinit"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/config/envconfig"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Approve(t *testing.T) {
	envconfig.InitEnvVars()

	appinit.Init()
	gin.SetMode(gin.TestMode)

	type network struct {
		name          string
		chainId       int64
		erc20Address  string
		erc721Address string
	}
	networks := []network{
		{
			name:          "polygon",
			chainId:       80001,
			erc721Address: "0x2d7882beDcbfDDce29Ba99965dd3cdF7fcB10A1e",
		}, {
			name:          "ethereum",
			chainId:       1,
			erc721Address: "0x514910771af9ca656af840dff83e8264ecf986ca",
		},
	}

	t.Run("Native token", func(t *testing.T) {
		for _, n := range networks {
			rr := httptest.NewRecorder()

			req := ApproveWithSalt{
				WalletAddress:   "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
				ToAddress:       "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
				ContractAddress: n.erc721Address,
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
			approveWithSalt(c)

			assert.Equal(t, 200, rr.Result().StatusCode)
		}

	})

	t.Run("ERC20 Token", func(t *testing.T) {
		for _, n := range networks {
			rr := httptest.NewRecorder()

			req := ApproveWithSalt{
				WalletAddress:   "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
				ToAddress:       "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
				ContractAddress: n.erc721Address,
			}
			d, e := json.Marshal(req)
			if e != nil {
				t.Fatal(e)
			}
			c, _ := gin.CreateTestContext(rr)
			reqUrl := fmt.Sprintf("/?erc20Address=%v", n.erc20Address)
			httpReq, e := http.NewRequest("GET", reqUrl, bytes.NewBuffer(d))
			if e != nil {
				t.Fatal(e)
			}
			c.Request = httpReq
			approveWithSalt(c)
			assert.Equal(t, 200, rr.Result().StatusCode)
		}

	})
	t.Run("ERC721 Token", func(t *testing.T) {
		for _, n := range networks {
			rr := httptest.NewRecorder()

			req := ApproveWithSalt{
				WalletAddress:   "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
				ToAddress:       "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
				ContractAddress: n.erc721Address,
			}
			d, e := json.Marshal(req)
			if e != nil {
				t.Fatal(e)
			}
			c, _ := gin.CreateTestContext(rr)
			reqUrl := fmt.Sprintf("/?erc721Address=%v", n.erc721Address)
			httpReq, e := http.
				NewRequest("GET", reqUrl,
					bytes.NewBuffer(d))
			if e != nil {
				t.Fatal(e)
			}
			c.Request = httpReq
			approveWithSalt(c)
			assert.Equal(t, 200, rr.Result().StatusCode)
		}

	})

}
