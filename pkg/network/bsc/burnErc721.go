package bsc

import (
	"math/big"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/generated/generc721"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet/rawtransaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BurnERC721(mnemonic string, contractAddr common.Address, tokenId big.Int) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return "", err
	}
	chainId, err := GetChainId()
	if err != nil {
		return "", err
	}
	tx, err := rawtransaction.SendRawTransaction(privKey, *client, int64(chainId), 310000, contractAddr, generc721.Erc721MetaData.ABI, "burn", &tokenId)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}
