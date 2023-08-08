package virtuacoinnft

import (
	"fmt"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/config/envconfig"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/generated/virtuacoinnft"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/bsc"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet/rawtransaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GrantRole(roleId string, walletAddress string, contractAddr common.Address) (string, error) {
	operatorPrivKey, err := wallet.GetWallet(envconfig.EnvVars.OPERATOR_MNEMONIC, bsc.GetPath())
	if err != nil {
		err = fmt.Errorf("failed to get operator wallet from mnemonic: %w", err)
		return "", err
	}

	client, err := ethclient.Dial(bsc.GetRpcUrl())
	if err != nil {
		err = fmt.Errorf("failed to dial rpc url: %w", err)
		return "", err
	}
	chainId, err := bsc.GetChainId()
	if err != nil {
		return "", err
	}

	roleIdBytesSlice, _ := hexutil.Decode(roleId)
	var roleIdBytes [32]byte
	copy(roleIdBytes[:], roleIdBytesSlice)
	creatorAddress := common.HexToAddress(walletAddress)

	tx, err := rawtransaction.SendRawTransaction(operatorPrivKey, *client, int64(chainId), 310000, contractAddr, virtuacoinnft.VirtuacoinnftABI, "grantRole", roleIdBytes, creatorAddress)
	if err != nil {
		err = fmt.Errorf("failed to send raw transaction: %w", err)
		return "", err
	}
	return tx.Hash().Hex(), nil
}
