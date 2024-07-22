package eths

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"

	"tidy/wallet/hdwallet"
)

func NewAcc(pass string) (string, error) {
	w, err := hdwallet.NewWallet("./data/keystore")
	if err != nil {
		fmt.Println("failed to NewWallet", err)
	}
	err = w.StoreKey(pass)
	if err != nil {
		fmt.Println("failed to StoreKey", err)
	}
	return w.Address.Hex(), nil
}

func KeccakHash(data []byte) []byte {
	return crypto.Keccak256(data)
}
