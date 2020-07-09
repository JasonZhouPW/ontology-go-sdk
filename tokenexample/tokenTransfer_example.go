package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/core/types"

	"github.com/ontio/ontology/common"
)
func main() {
	fmt.Println("test")


	//testUrl := "http://127.0.0.1:20336"
	testUrl := "http://polaris2.ont.io:20336"
	//initialize ontsdk
	ontSdk := sdk.NewOntologySdk()
	//suppose you already start up a local wasm ontology node
	ontSdk.NewRpcClient().SetAddress(testUrl)
	//your wallet file
	wallet, err := ontSdk.OpenWallet("./wallet.dat")
	if err != nil {
		fmt.Printf("error in OpenWallet:%s\n", err)
		return
	}

	//modify me
	walletpassword := "123456"

	signer, err := wallet.GetDefaultAccount([]byte(walletpassword))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	acct2 ,err := wallet.GetAccountByIndex(2,[]byte(walletpassword))
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	gasprice := uint64(2500)
	gaslimit := uint64(20000)

	tx,err:= ontSdk.Native.Ont.NewTransferTransaction(gasprice,gaslimit,signer.Address,acct2.Address,1)
	if err != nil {
		panic(err)
	}
	//tx.Payer = signer.Address
	tx.Payer = acct2.Address
	itx,err := tx.IntoImmutable()
	if err != nil {
		panic(err)
	}
	itxhex := hex.EncodeToString(common.SerializeToBytes(itx))
	fmt.Printf("unsigned tx:= %s\n",itxhex)
	txhash := tx.Hash()
	fmt.Printf("unsigned txhash := %s\n",hex.EncodeToString(txhash.ToArray()))


	err = ontSdk.SignToTransaction(tx,acct2)
	if err != nil {
		panic(err)
	}



	err = ontSdk.SignToTransaction(tx,signer)
	if err != nil {
		panic(err)
	}
	for _,sig := range tx.Sigs {
		sigdata := sig.SigData[0]
		fmt.Printf("sigdata := %s\n",hex.EncodeToString(sigdata))
		fmt.Printf("publickey := %s\n",hex.EncodeToString(keypair.SerializePublicKey(sig.PubKeys[0])))
	}


	//sigdata := tx.Sigs[0].SigData[0]
	txhash = tx.Hash()
	fmt.Printf("signed txhash := %s\n",hex.EncodeToString(txhash.ToArray()))
	//fmt.Printf("sigdata := %s\n",hex.EncodeToString(sigdata))
	//fmt.Printf("publickey := %s\n",tx.Sigs[0].PubKeys[0])
	itx,err = tx.IntoImmutable()
	if err != nil {
		panic(err)
	}
	itxhex = hex.EncodeToString(common.SerializeToBytes(itx))
	fmt.Printf("signed tx:= %s\n",itxhex)

	bs ,err := hex.DecodeString("00d12981cfeec409000000000000204e000000000000ffe723aefd01bac311d8b16ff8bfd594d77f31ee7100c66b14092118e0112274581b60dfb6fedcbfdcfc044be76a7cc814ffe723aefd01bac311d8b16ff8bfd594d77f31ee6a7cc8516a7cc86c51c1087472616e736665721400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65000241400d92482c51240b7b8497ec8c1efb41d1d94e434433fb0b3642e4483ecd2083c672b9e73ff4194d84cc84012db78c2d283a4bb1dbec5dc66015a42dac0042f8d7232102263e2e1eecf7a45f21e9e0f865510966d4e93551d95876ecb3c42acf2b68aaaeac4140637fdd1ac4b0cc5f36334b48b5fcd9494531ba88b5ec3e6ceb0fc49461e0e856f8a3aa4394be2c4a15a0358767a66d1acbf9a2ac9a12e4c628c9f8762db353db232103944e3ff777b14add03a76fd6767aaf4a65c227ec201375d9118d4e6b272494c7ac")

	dtx ,err:= types.TransactionFromRawBytes(bs)
	dhas := dtx.Hash()

	fmt.Printf(dhas.ToHexString())


}