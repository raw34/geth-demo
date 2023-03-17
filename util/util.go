package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/raw34/eth-demo/contracts/store"
	"log"
	"math/big"
)

func Deploy() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	auth := getAccountAuth(client, "b7bb8468fd606f0d0ea6802eb505c814e39dc8732f452aa729eb3d7acf8bdb51")

	input := "1.0"
	address, tx, instance, err := store.DeployContracts(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x7C4028C06796F82Cf58Fd9c4BA7d369C2bD3B465
	fmt.Println(tx.Hash().Hex()) // 0xa1a50c789ae7f53d63275f3e87b8c919e1ab4aa94f4e77ef3a7c897a0f1b76e6

	_ = instance
}

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(3000000) // in units
	//auth.GasPrice = big.NewInt(1000000)

	return auth
}
