package polygon

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"strings"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ERC20 ABI (same as in erc20.go)
const erc20ABI = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"inputs":[{"name":"_initialAmount","type":"uint256"},{"name":"_tokenName","type":"string"},{"name":"_decimalUnits","type":"uint8"},{"name":"_tokenSymbol","type":"string"}],"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"}]`

func TransferERC20(mnemonic string, toAddress common.Address, contractAddr common.Address, amount *big.Int) (string, error) {
	privateKeyHex, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		log.Printf("Failed to connect to Ethereum client: %v", err)
		return "", err
	}

	// privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	// if err != nil {
	// 	log.Printf("Invalid private key: %v", err)
	// 	return "", err
	// }

	publicKey := privateKeyHex.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Printf("Error casting public key to ECDSA")
		return "", errors.New("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("Failed to retrieve nonce: %v", err)
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Failed to suggest gas price: %v", err)
		return "", err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Printf("Failed to get chain ID: %v", err)
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyHex, chainID)
	if err != nil {
		log.Printf("Failed to create transactor: %v", err)
		return "", err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice

	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		log.Printf("Failed to parse ABI: %v", err)
		return "", err
	}
	instance := bind.NewBoundContract(contractAddr, parsedABI, client, client, client)

	tx, err := instance.Transact(auth, "transfer", toAddress, amount)
	if err != nil {
		log.Printf("Failed to send transaction: %v", err)
		return "", err
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}
