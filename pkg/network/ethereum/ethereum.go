package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"strings"

	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/config/envconfig"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/generated/generc20"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/generated/generc721"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet"
	rawtrasaction "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet/rawtransaction"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func GetChainId() (int, error) {
	return envconfig.EnvVars.CHAIN_ID_ETHEREUM, nil
}

func GetPath() string {
	return "m/44H/60H/0H/0/0"
}

func GetRpcUrl() string {
	return envconfig.EnvVars.NETWORK_RPC_URL_ETHEREUM
}

func GetBalance(mnemonic string) (*big.Int, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return nil, err
	}
	publicKey := privKey.PublicKey
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, err
	}
	bal, err := client.BalanceAt(context.Background(), crypto.PubkeyToAddress(publicKey), nil)
	if err != nil {
		return nil, err
	}
	return bal, nil
}

func GetERC20Balance(mnemonic string, contractAddr common.Address) (*big.Int, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return nil, err
	}
	publicKey := privKey.PublicKey

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, err
	}
	ins, err := generc20.NewErc20(contractAddr, client)
	if err != nil {
		return nil, err
	}
	bal, err := ins.BalanceOf(nil, crypto.PubkeyToAddress(publicKey))
	if err != nil {
		return nil, err
	} else {
		return bal, nil
	}
}

func Transfer(mnemonic string, to common.Address, value big.Int) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}
	publicKey := privKey.PublicKey

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return "", err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(publicKey))
	if err != nil {
		return "", err
	}

	maxPriorityFeePerGas, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		logo.Errorf("failed to suggestGasTipCap, error %s", err)
		return "", err
	}
	chainId, err := GetChainId()
	if err != nil {
		return "", err
	}
	config := &params.ChainConfig{
		ChainID: big.NewInt(int64(chainId)),
	}
	bn, _ := client.BlockNumber(context.Background())

	bignumBn := big.NewInt(0).SetUint64(bn)
	blk, _ := client.BlockByNumber(context.Background(), bignumBn)
	baseFee := misc.CalcBaseFee(config, blk.Header())
	big2 := big.NewInt(2)
	mulRes := big.NewInt(0).Mul(baseFee, big2)
	maxFeePerGas := big.NewInt(0).Add(mulRes, maxPriorityFeePerGas)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainId)),
		Nonce:     nonce,
		GasFeeCap: maxFeePerGas,
		GasTipCap: maxPriorityFeePerGas,
		Gas:       21000,
		To:        &to,
		Value:     &value,
	})
	types.SignTx(tx, types.NewLondonSigner(big.NewInt(int64(chainId))), privKey)
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(int64(chainId))), privKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

// func TransferERC20(mnemonic string, toAddress common.Address, contractAddr common.Address, amount big.Int) (string, error) {
// 	// TODO: Ammount, check in goethbook guide
// 	privKey, err := wallet.GetWallet(mnemonic, GetPath())
// 	if err != nil {
// 		return "", err
// 	}

// 	client, err := ethclient.Dial(GetRpcUrl())
// 	if err != nil {
// 		return "", err
// 	}
// 	ins, err := generc20.NewErc20(contractAddr, client)
// 	if err != nil {
// 		return "", err
// 	}
// 	decimals, err := ins.Decimals(nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	decimalsCal := big.NewInt(0)
// 	decimalsCal.Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
// 	amount.Mul(&amount, decimalsCal)
// 	chainId, err := GetChainId()
// 	if err != nil {
// 		return "", err
// 	}
// 	tx, err := rawtrasaction.SendRawTransaction(privKey, *client, int64(chainId), 310000, contractAddr, generc20.Erc20MetaData.ABI, "transfer", toAddress, &amount)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tx.Hash().Hex(), nil
// }

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

func TransferERC721(mnemonic string, toAddress common.Address, contractAddr common.Address, tokenId big.Int) (string, error) {
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
	publicKey := privKey.PublicKey
	fromAddr := crypto.PubkeyToAddress(publicKey)
	tx, err := rawtrasaction.SendRawTransaction(privKey, *client, int64(chainId), 310000, contractAddr, generc721.Erc721MetaData.ABI, "safeTransferFrom", fromAddr, toAddress, &tokenId)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil

}

func GetNetworkInfo() (*networkInfo, error) {
	chainId, err := GetChainId()
	if err != nil {
		return nil, err
	}
	return &networkInfo{
		Name:    "ethereum",
		ChainId: big.NewInt(int64(chainId)),
	}, nil
}

func GetWalletAddres(mnemonic string) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}
	walletAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	return walletAddr.String(), nil
}

func GetERC20BalanceInDecimalsFromWalletAddress(walletAddress common.Address, contractAddr common.Address) (*big.Float, error) {
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, err
	}
	ins, err := generc20.NewErc20(contractAddr, client)
	if err != nil {
		return nil, err
	}
	decimals, err := ins.Decimals(nil)
	if err != nil {
		return nil, err
	}
	decimalsCal := big.NewInt(0)
	decimalsCal.Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	// amount.Mul(&amount, decimalsCal)
	bal, err := ins.BalanceOf(nil, walletAddress)
	if err != nil {
		return nil, err
	}
	token := big.NewFloat(0).SetInt(big.NewInt(decimalsCal.Int64()))
	balanceInDecimals := big.NewFloat(0).SetInt(bal)
	balanceInDecimals.Quo(balanceInDecimals, token)
	return balanceInDecimals, nil
}
