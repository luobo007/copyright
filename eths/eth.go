package eths

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"tidy/wallet/hdwallet"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var PXA_ADDR = "0xe3dbB1d7A9e869467F55478d88552D2240ac7AC4"
var PXC_ADDR = "0xCCAEE0d151a7269D22d66f3E123292081C36DB22"

// Geth客户端全局连接句柄
var ethcli *ethclient.Client

// PXA合约全局实例
var instancePXA *Pxa721

// PXC合约全局实例
var instancePXC *Pxc20

func init() {
	cli, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Panic("Failed to ethclient.Dial ", err)
	}
	ethcli = ethcli
	instance, err := NewPxa721(common.HexToAddress(PXA_ADDR), cli)
	if err != nil {
		log.Panic("Failed to NewPxa721", err)
	}
	instancePXA = instance
	ins, err := NewPxc20(common.HexToAddress(PXC_ADDR), cli)
	if err != nil {
		log.Panic("Failed to NewPxa721", err)
	}
	instancePXC = ins
}

// 上传图片调用
func UploadPic(from, pass, to string, tokenid *big.Int) error {
	//3. 设置签名 -- 需要owner的keystore文件
	w, err := hdwallet.LoadWalletByPass(from, "./data", pass)
	if err != nil {
		fmt.Println("failed to LoadWalletByPass", err)
		return err
	}
	auth := w.HdKeyStore.NewTransactOpts()
	//4. 调用
	_, err = instancePXA.UploadMint(auth, common.HexToAddress(to), tokenid)
	if err != nil {
		fmt.Println("failed to UploadMint  ", err)
		return err
	}
	return nil
}

const adminkey = `{"address":"be9ab3661f83d3385b896513ead7ec3ef78d9bbc","crypto":{"cipher":"aes-128-ctr","ciphertext":"2ac6475e6f319f330e93f315849178dffda7b05636b9104a0db808cbee256930","cipherparams":{"iv":"ebbc07a5f3af629b45824e56d2003f40"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"fc6fb2cb01716c9c7c1107ecf00752c119cd6311c6f56abbf7275fc6871893dc"},"mac":"425509f8ec535f11b14f8bb4febbe251669d4edbb93737003f68db60c0aff3c2"},"id":"8a2b1e3a-7bba-42c6-b9d6-deef3bfa4dc4","version":3}`
const adminAddr = "be9ab3661f83d3385b896513ead7ec3ef78d9bbc"

// 授权
func SetApprove(from, pass string, tokenid *big.Int) error {

	//3. 设置签名 -- 需要owner的keystore文件
	w, err := hdwallet.LoadWalletByPass(from, "./data", pass)
	if err != nil {
		fmt.Println("failed to LoadWalletByPass", err)
		return err
	}
	auth := w.HdKeyStore.NewTransactOpts()
	//4. 调用
	_, err = instancePXA.Approve(auth, common.HexToAddress(adminAddr), tokenid)
	if err != nil {
		fmt.Println("failed to Approve  ", err)
		return err
	}
	return nil
}

// 转移erc20
func TransferPXC(from, pass, to string, value *big.Int) error {
	//3. 设置签名 -- 需要owner的keystore文件
	w, err := hdwallet.LoadWalletByPass(from, "./data", pass)
	if err != nil {
		fmt.Println("failed to LoadWalletByPass", err)
		return err
	}
	auth := w.HdKeyStore.NewTransactOpts()
	//4. 调用
	_, err = instancePXC.Transfer(auth, common.HexToAddress(to), value)
	if err != nil {
		fmt.Println("failed to Transfer  ", err)
		return err
	}
	return nil
}

// 转移erc721
func PartTransferPXA(from, to string, tokenid, weight, price *big.Int) error {
	//3. 设置签名 -- 需要owner的keystore文件
	keyin := strings.NewReader(adminkey)
	auth, err := bind.NewTransactor(keyin, "123")
	//4. 调用
	_, err = instancePXA.PartTransferFrom(auth, common.HexToAddress(from), common.HexToAddress(to), tokenid, weight, price)
	if err != nil {
		fmt.Println("failed to TransferPXA  ", err)
		return err
	}
	return nil
}
