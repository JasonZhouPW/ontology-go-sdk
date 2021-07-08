package main

import (
	"fmt"
	"time"

	"github.com/ontio/ontology/common"

	sdk "github.com/ontio/ontology-go-sdk"
)

func main(){
	fmt.Println("==============start===================")


	nodeUrl := "http://dappnode1.ont.io:20336"
	pwd := "ZZZZZ"
	contractHex := ""

	//testUrl := "http://polaris2.ont.io:20336"
	//initialize ontsdk
	ontSdk := sdk.NewOntologySdk()
	//suppose you already start up a local wasm ontology node
	ontSdk.NewRpcClient().SetAddress(nodeUrl)
	//your wallet file
	wallet, err := ontSdk.OpenWallet("../wallet.dat")
	if err != nil{
		panic(err)
	}


	signer, err := wallet.GetDefaultAccount( []byte(pwd))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}

	contractAddr ,_ := common.AddressFromHexString(contractHex)

	txHash,err:= ontSdk.NeoVM.InvokeNeoVMContract(uint64(2500),uint64(200000),signer,signer,contractAddr,[]interface{}{"lock",[]interface{}{
		signer.Address,
		uint64(2),
		"0x26356Cb66F8fd62c03F569EC3691B6F00173EB02",
		uint64(1000000000000000),
		uint64(100)}})

	fmt.Printf("txhash:%s\n",txHash.ToHexString())
	timeoutSec := 60 * time.Second

	_, err = ontSdk.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		fmt.Printf("error in WaitForGenerateBlock:%s\n", err)

		panic(err)
	}
	//get smartcontract event by txhash
	events, err := ontSdk.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		fmt.Printf("error in GetSmartContractEvent:%s\n", err)
		panic(err)
	}
	fmt.Printf("event is %v\n", events)
	//State = 0 means transaction failed
	if events.State == 0 {
		fmt.Printf("error in events.State is 0 failed.\n")

		panic(err)
	}
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		//you check the notify here
		fmt.Printf("%+v\n", notify)
	}
}
