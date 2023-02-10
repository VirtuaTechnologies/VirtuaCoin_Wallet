package virtuacoinnft

import (
	"fmt"
	"math/big"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/generated/virtuacoinnft"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/polygon"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet/rawtransaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Burn(mnemonic string, contractAddr common.Address, tokenId big.Int) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, polygon.GetPath())
	if err != nil {
		return "", err
	}

	client, err := ethclient.Dial(polygon.GetRpcUrl())
	if err != nil {
		return "", err
	}
	chainId, err := polygon.GetChainId()
	if err != nil {
		return "", err
	}

	tx, err := rawtransaction.SendRawTransaction(privKey, *client, int64(chainId), 310000, contractAddr, virtuacoinnft.VirtuacoinnftABI, "burn", &tokenId)
	if err != nil {
		err = fmt.Errorf("failed to send raw transaction: %w", err)
		return "", err
	}
	return tx.Hash().Hex(), nil

}
