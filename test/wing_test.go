package test

import (
	"encoding/hex"
	"fmt"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/oep4"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/big"
	"testing"
	"time"
)


type AssetInfo struct {
	AssetName string
	UnderlyingAddress string
	MarketAddress string
}

var asset_infos = []AssetInfo{
	{AssetName:"ETH",UnderlyingAddress:"7009a2f7c8a2e45fa386a6078c7bfeaf518be487",MarketAddress:"9be819ef7f45fd45fdc2680a7691317814f408a8"},
	{AssetName:"WING",UnderlyingAddress:"ff31ec74d01f7b7d45ed2add930f5d2239f7de33",MarketAddress:"36cf1ec1cae26e997bf79787b66a3fc7b9917b79"},
	{AssetName:"USDT",UnderlyingAddress:"c6f91c11d740d39943b99a6b1c6fd2b5f476e2a3",MarketAddress:"abccea5e651d3030a8199e9d76e5687b9df89479"},
	{AssetName:"WBTC",UnderlyingAddress:"aede525f05065306423a5522bfcd31b5847ffa52",MarketAddress:"55eb94463a3247687bf7e1da3ef82c06af6b0cdb"},
	{AssetName:"renBTC",UnderlyingAddress:"ec547bc4486dea97cb659f1fe73407922f9e63c8",MarketAddress:"92a0ba9d8597fc53512230e4c9dec60998ada9a5"},
	{AssetName:"USDC",UnderlyingAddress:"07a12c0a6bdce4df04ef4b2045d1b0fd63a56e25",MarketAddress:"e39e960c884f4d59b41b95387b14d0e4c57f3bd9"},
	{AssetName:"UNI",UnderlyingAddress:"b79d2064947f61070cb68ef26cbc12cbf3b98d9e",MarketAddress:"44ac2d9de9bab53767dcf514b4320df739bae517"},
	{AssetName:"SUSD",UnderlyingAddress:"37f4497b6f5f511e73843a0bda1042777666f7ec",MarketAddress:"b447d8946c51d8da28ec234a343488e407a35a38"},
	{AssetName:"DAI",UnderlyingAddress:"96cf88356123592835a2fa75068a242260be1791",MarketAddress:"3ba975c2ac17b3b4ddab88fa96db142e99f4ff25"},
}

func InitEnv() (*sdk.OntologySdk ,*sdk.Wallet,error){

	fmt.Println("==========================start============================")
	//testUrl := "http://127.0.0.1:20336"
	testUrl := "http://polaris2.ont.io:20336"
	//initialize ontsdk
	ontSdk := sdk.NewOntologySdk()
	//suppose you already start up a local wasm ontology node
	ontSdk.NewRpcClient().SetAddress(testUrl)
	//your wallet file
	wallet, err := ontSdk.OpenWallet("../wallet.dat")
	if err != nil {
		fmt.Printf("error in OpenWallet:%s\n", err)
		return nil,nil,err
	}
	return ontSdk,wallet,err
}

func Test_contractExist(t *testing.T){
	ontsdk,_,err := InitEnv()
	if err != nil {
		panic(err)
	}
	address := []string{"9be819ef7f45fd45fdc2680a7691317814f408a8",
		"36cf1ec1cae26e997bf79787b66a3fc7b9917b79",
	"cacfa48d3fbfc70ac6c055760fde73f44955abb5",
	"7009a2f7c8a2e45fa386a6078c7bfeaf518be487",
	"ff31ec74d01f7b7d45ed2add930f5d2239f7de33",
	"6165f58fe4210629c2788ac49f5db437b56122d6",
	"a1ac292510459c2583079f53c150e8aa885e6b4b",
	"22fc643aa439ec713e936a3848976f734a556046",
	"33c439c502cb4b6ac5a1e8057a65fe1fa7c300e2",
	"abccea5e651d3030a8199e9d76e5687b9df89479"}

	for _,addr:=range address{
		//contractAddr ,_:= common.AddressFromHexString(addr)
		cd,err := ontsdk.GetSmartContract(addr)
		if err != nil{
			fmt.Printf("addr:%s ,err:%s\n",addr,err.Error())
		}
		if cd == nil {
			fmt.Printf("addr:%s ,is nil\n",addr)

		}

	}

}

func InitEnvLocal() (*sdk.OntologySdk ,*sdk.Wallet,error){

	fmt.Println("==========================start============================")
	testUrl := "http://127.0.0.1:20336"
	//testUrl := "http://polaris2.ont.io:20336"
	//initialize ontsdk
	ontSdk := sdk.NewOntologySdk()
	//suppose you already start up a local wasm ontology node
	ontSdk.NewRpcClient().SetAddress(testUrl)
	//your wallet file
	wallet, err := ontSdk.OpenWallet("../wallet.dat")
	if err != nil {
		fmt.Printf("error in OpenWallet:%s\n", err)
		return nil,nil,err
	}
	return ontSdk,wallet,err
}


func Test_DepolyNeoContract(t *testing.T) {
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	err = deployNeo(ontsdk,signer ,"./lockproxy.avm",false)
	if err != nil {
		panic(err)
	}
}

func Test_DepolyNeoTestInvokeContract(t *testing.T) {
	ontsdk,wallet,err := InitEnvLocal()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	err = deployNeo(ontsdk,signer ,"./testinvoke.avm",true)
	if err != nil {
		panic(err)
	}
}

func TestInvokeNeoContract(t *testing.T){
	ontsdk,wallet,err := InitEnvLocal()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	addr,_ := common.AddressFromHexString("5b6d1d7419b76c33569e35842f6830ab44d4620b")
	param,_:= hex.DecodeString("0010040000000106000000737570706c79011400000026356cb66f8fd62c03f569ec3691b6f00173eb020287e48b51affe7b8c07a686a35fe4a2c8f7a20970040080c6a47e8d03000000000000000000")

	txHash,err := ontsdk.NeoVM.InvokeNeoVMContract(uint64(0),uint64(20000),signer,signer,addr,[]interface{}{"invokeWing",[]interface{}{"AHVQcVFuiCRU6rxkN6zkfvFU88rbpJ5oZt",param}})
	assert.Nil(t,err,"err is not nil")
	_, err = ontsdk.WaitForGenerateBlock(30*time.Second)
	if err != nil {
		fmt.Printf("error in WaitForGenerateBlock:%s\n", err)
		panic(err)
	}
	fmt.Printf("init txhash is :%s\n", txHash.ToHexString())
	//get smartcontract event by txhash
	events, err := ontsdk.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		fmt.Printf("error in GetSmartContractEvent:%s\n", err)

		panic(err)
	}
	fmt.Printf("event is %v\n", events)
	//State = 0 means transaction failed
	if events.State == 0 {
		fmt.Printf("error in events.State is 0 failed.\n")

		return
	}
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		fmt.Printf("%+v\n", notify)
	}
	return
}



func Test_DepolyNeoLockWrapperContract(t *testing.T) {
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	err = deployNeo(ontsdk,signer,"./lockwrapper.avm",true)
	if err != nil {
		panic(err)
	}
}




func deployNeo(ontSdk *sdk.OntologySdk,signer *sdk.Account,avmfile string,local bool) error {
	//avmfile:= "./lockproxy.avm"
	//set timeout
	timeoutSec := 60 * time.Second
	//address1 := "AX8opZCQBpEpYsFPKpZHNguWz2s3xpT7Wk"

	// read wasm file and get the Hex fmt string
	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		fmt.Printf("error in ReadFile:%s\n", err)

		return err
	}

	codeHash := common.ToHexString(code)

	//===========================================
	gasprice := uint64(2500)
	if local{
		gasprice = uint64(0)
	}
	//invokegaslimit := uint64(200000)
	deploygaslimit := uint64(200000000)
	// deploy the wasm contract
	fmt.Println("======DeployNeoVMSmartContract ==========")

	txHash, err := ontSdk.NeoVM.DeployNeoVMSmartContract(
		gasprice,
		deploygaslimit,
		signer,
		true,
		codeHash,
		"lockproxy",
		"1.0",
		"author",
		"email",
		"desc",
	)
	if err != nil {
		fmt.Printf("error in DeployNeoVMSmartContract:%s\n", err)

		return err
	}
	_, err = ontSdk.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		fmt.Printf("error in WaitForGenerateBlock:%s\n", err)

		return err
	}
	fmt.Printf("the deploy contract txhash is %s\n", txHash.ToHexString())

	//calculate the contract address from code
	contractAddr, err := utils.GetContractAddress(codeHash)
	if err != nil {
		fmt.Printf("error in GetContractAddress:%s\n", err)

		return err
	}
	fmt.Printf("the contractAddr is %s\n", contractAddr.ToBase58())
	fmt.Printf("the contractAddr is %s\n", contractAddr.ToHexString())
	fmt.Printf("the reversed contractAddr is %s\n",hex.EncodeToString(contractAddr[:]))

	fmt.Println("============Done===============")
	return nil
}

func Test_ontWrapper(t *testing.T){
	ontSdk,wallet,err := InitEnv()
	//ontsdk,wallet,err := InitEnvLocal()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}

	wrapperaddr,_:= common.AddressFromHexString("a1ac292510459c2583079f53c150e8aa885e6b4b")
	eth_asset_addr,_:= common.AddressFromHexString("7009a2f7c8a2e45fa386a6078c7bfeaf518be487")

	txHash,err:= ontSdk.NeoVM.InvokeNeoVMContract(uint64(2500),uint64(200000),signer,signer,wrapperaddr,[]interface{}{"lock",[]interface{}{eth_asset_addr,
																									signer.Address,
																									uint64(2),
																									"0x26356Cb66F8fd62c03F569EC3691B6F00173EB02",
																									uint64(1000000000000000),
		uint64(100),eth_asset_addr}})
	if err != nil{
		panic(err)
	}
	timeoutSec := 100 * time.Second

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

func Test_DeployWasmUserAgent(t *testing.T){
	now := time.Now().Nanosecond()
	ontsdk,wallet,err := InitEnv()
	//ontsdk,wallet,err := InitEnvLocal()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())

	err,contractAddr := deployWasm(ontsdk,signer,false)
	if err != nil {
		panic(err)
	}
	fmt.Println("====setLockProxy====")
	setLockProxy(ontsdk,signer,contractAddr)
	//setLockProxyTest(ontsdk,signer,contractAddr)

	fmt.Println("====setWingToken====")
	setWingToken(ontsdk,signer,contractAddr)

	fmt.Println("====setUnderlying====")
	setUnderlying(ontsdk,signer,contractAddr)

	fmt.Println("====setMarket====")
	setMarket(ontsdk,signer,contractAddr)

	fmt.Println("====setGovAddr====")
	setGovAddr(ontsdk,signer,contractAddr)

	fmt.Println("====setOracle====")
	setOracle(ontsdk,signer,contractAddr)

	fmt.Println("====setOriginLockProxy====")
	setOriginLockProxy(ontsdk,signer,contractAddr)

	//fmt.Printf("====setFeeCollector====")
	//setFeeCollector(ontsdk,signer,contractAddr)


	//for test
	//fmt.Println("====supply====")
	//supply(ontsdk,signer,contractAddr)
	////
	//fmt.Println("====querybalance====")
	//balance := querySupplyBalance(ontsdk,signer,contractAddr)
	//fmt.Printf("balance is %d\n",balance)
	////
	//fmt.Println("====withdraw====")
	//withdraw(ontsdk,signer,contractAddr)
	//
	//fmt.Println("====queryWingBalance====")
	//queryWingBalance(ontsdk,contractAddr)
	//
	//fmt.Println("====withdrawWing====")
	//withdrawWing(ontsdk,signer,contractAddr)



	fmt.Println("====set evn done====")
	fmt.Printf("time used:%d\n",time.Now().Nanosecond() - now)

}


//modify after deploy user agent
var userAgentAddr = "b27fb30b951bf22501a1adb89733e9198264389a"
func deployWasm(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool) (error,common.Address) {
	//get a compiled wasm file from ont_cpp
	wasmfile := "./user_agent_opt.wasm"

	//set timeout
	timeoutSec := 60 * time.Second
	//address1 := "AX8opZCQBpEpYsFPKpZHNguWz2s3xpT7Wk"

	// read wasm file and get the Hex fmt string
	code, err := ioutil.ReadFile(wasmfile)
	if err != nil {
		fmt.Printf("error in ReadFile:%s\n", err)

		return err,common.Address{}
	}

	codeHash := common.ToHexString(code)
	//calculate the contract address from code
	contractAddr, err := utils.GetContractAddress(codeHash)
	if err != nil {
		fmt.Printf("error in GetContractAddress:%s\n", err)

		return err,common.Address{}

	}
	fmt.Printf("the contractAddr is %s\n", contractAddr.ToBase58())
	fmt.Printf("the contractAddr is %s\n", contractAddr.ToHexString())
	fmt.Printf("the revert contractAddr is %s\n", hex.EncodeToString(contractAddr[:]))

	//===========================================
	gasprice := uint64(2500)
	if local{
		gasprice = uint64(0)
	}
	//invokegaslimit := uint64(200000)
	deploygaslimit := uint64(2000000000)
	// deploy the wasm contract
	fmt.Println("======DeployWasmVMSmartContract ==========")
	txHash, err := ontSdk.WasmVM.DeployWasmVMSmartContract(
		gasprice,
		deploygaslimit,
		signer,
		codeHash,
		"user agent wasm",
		"1.0",
		"author",
		"email",
		"desc",
	)
	if err != nil {
		fmt.Printf("error in DeployWasmVMSmartContract:%s\n", err)

		return err,common.Address{}

	}
	_, err = ontSdk.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		fmt.Printf("error in WaitForGenerateBlock:%s\n", err)

		return err,common.Address{}
	}
	fmt.Printf("the deploy contract txhash is %s\n", txHash.ToHexString())

	//calculate the contract address from code
	//contractAddr, err := utils.GetContractAddress(codeHash)
	//if err != nil {
	//	fmt.Printf("error in GetContractAddress:%s\n", err)
	//
	//	return err
	//}
	//fmt.Printf("the contractAddr is %s\n", contractAddr.ToBase58())
	//fmt.Printf("the contractAddr is %s\n", contractAddr.ToHexString())
	fmt.Println("============Done===============")

	return nil,contractAddr
}

func Test_InitUserAgent(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	initUserAgent(ontsdk,signer,contractAddr)
}

func initUserAgent(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	err := invokeWasm(ontsdk,signer,contractAddr,"init",[]interface{}{})
	if err != nil {
		panic(err)
	}
}

func Test_SetLockProxy(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	//lpaddr,_:= common.AddressFromHexString("cacfa48d3fbfc70ac6c055760fde73f44955abb5")
	setLockProxy(ontsdk,signer,contractAddr)
}

func setLockProxy(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	lpaddr,_:= common.AddressFromHexString("cacfa48d3fbfc70ac6c055760fde73f44955abb5")
	err := invokeWasm(ontsdk,signer,contractAddr,"setLockProxyAddr",[]interface{}{lpaddr})
	if err != nil {
		panic(err)
	}
}

func Test_setLockProxyTest(t *testing.T)  {
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	//lpaddr,_:= common.AddressFromHexString("cacfa48d3fbfc70ac6c055760fde73f44955abb5")
	setLockProxyTest(ontsdk,signer,contractAddr)
}

func setLockProxyTest(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	//lpaddr,_:= common.AddressFromHexString("cacfa48d3fbfc70ac6c055760fde73f44955abb5")
	err := invokeWasm(ontsdk,signer,contractAddr,"setLockProxyAddr",[]interface{}{signer.Address})
	if err != nil {
		panic(err)
	}
}

func setOriginLockProxy(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	lpaddr,_:= common.AddressFromHexString("33c439c502cb4b6ac5a1e8057a65fe1fa7c300e2")
	err := invokeWasm(ontsdk,signer,contractAddr,"setOriginLockProxy",[]interface{}{lpaddr})
	if err != nil {
		panic(err)
	}
}

func setFeeCollector(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	lpaddr,_:= common.AddressFromBase58("AXdmdzbyf3WZKQzRtrNQwAR91ZxMUfhXkt")
	err := invokeWasm(ontsdk,signer,contractAddr,"setOriginLockProxy",[]interface{}{lpaddr})
	if err != nil {
		panic(err)
	}
}


func Test_SetUnderlying(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	setUnderlying(ontsdk,signer,contractAddr)

}

func setUnderlying(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){

	for _,asset := range asset_infos{
		fmt.Printf("===setting underlying:%v\n",asset)

		asset_address,_ := common.AddressFromHexString(asset.UnderlyingAddress)
		err := invokeWasm(ontsdk,signer,contractAddr,"setUnderlyingName",[]interface{}{asset_address,asset.AssetName})
		if err != nil {
			panic(err)
		}
	}

	//
	//eth_asset_addr,_:= common.AddressFromHexString("7009a2f7c8a2e45fa386a6078c7bfeaf518be487")
	//err := invokeWasm(ontsdk,signer,contractAddr,"setUnderlyingName",[]interface{}{eth_asset_addr,"ETH"})
	//if err != nil {
	//	panic(err)
	//}
	//wing_asset_addr,_:= common.AddressFromHexString("ff31ec74d01f7b7d45ed2add930f5d2239f7de33")
	//err = invokeWasm(ontsdk,signer,contractAddr,"setUnderlyingName",[]interface{}{wing_asset_addr,"WING"})
	//if err != nil {
	//	panic(err)
	//}
	//usdt_asset_addr,_:= common.AddressFromHexString("c6f91c11d740d39943b99a6b1c6fd2b5f476e2a3")
	//err = invokeWasm(ontsdk,signer,contractAddr,"setUnderlyingName",[]interface{}{usdt_asset_addr,"USDT"})
	//if err != nil {
	//	panic(err)
	//}
}


func Test_SetMarket(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	setMarket(ontsdk,signer,contractAddr)

}

func setMarket(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){

	type MarketInfo struct{
		Address string
		Name string
	}

	//markets := []MarketInfo{{Address:"9be819ef7f45fd45fdc2680a7691317814f408a8",Name:"ETH"},
	//	{Address:"36cf1ec1cae26e997bf79787b66a3fc7b9917b79",Name:"WING"},{},
	//	{Address:"abccea5e651d3030a8199e9d76e5687b9df89479",Name:"USDT"}}

	for _,m:=range asset_infos{
		fmt.Printf("===setting market:%v\n",m)
		market,_:= common.AddressFromHexString(m.MarketAddress)
		err := invokeWasm(ontsdk,signer,contractAddr,"setMarket",[]interface{}{m.AssetName,market})
		if err != nil {
			panic(err)
		}
	}
}

func Test_SetGovAddr(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	//contractAddr,_ := common.AddressFromHexString("ce9e78b9ce27dd5a5a9f614bf08c8c6c2ce3993d")
	setGovAddr(ontsdk,signer,contractAddr)
	//gov,_:= common.AddressFromHexString("6165f58fe4210629c2788ac49f5db437b56122d6")
	//err = invokeWasm(ontsdk,signer,contractAddr,"setGovAddress",[]interface{}{gov})
	//if err != nil {
	//	panic(err)
	//}

}

func Test_SetLockproxyWrapper(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	setOntLockproxyWrapper(ontsdk,signer,contractAddr)
}

func setGovAddr(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	gov,_:= common.AddressFromHexString("6165f58fe4210629c2788ac49f5db437b56122d6")
	err := invokeWasm(ontsdk,signer,contractAddr,"setGovAddress",[]interface{}{gov})
	if err != nil {
		panic(err)
	}
}
func setOntLockproxyWrapper(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	wrapper,_:= common.AddressFromHexString("a1ac292510459c2583079f53c150e8aa885e6b4b")
	err := invokeWasm(ontsdk,signer,contractAddr,"setLockWrapperAddress",[]interface{}{wrapper})
	if err != nil {
		panic(err)
	}
}

func setOracle(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	oralce ,_:= common.AddressFromHexString("22fc643aa439ec713e936a3848976f734a556046")
	err := invokeWasm(ontsdk,signer,contractAddr,"setOracleAddress",[]interface{}{oralce})
	if err != nil {
		panic(err)
	}
}


func Test_Supply(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	//contractAddr,_ := common.AddressFromHexString("ce9e78b9ce27dd5a5a9f614bf08c8c6c2ce3993d")
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	//asset_addr,_:= common.AddressFromHexString("7009a2f7c8a2e45fa386a6078c7bfeaf518be487")
	supply(ontsdk,signer,contractAddr)

}

func supply(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	fmt.Println("====approve asset to controller====")
	//eth
	//asset_addr,_:= common.AddressFromHexString("7009a2f7c8a2e45fa386a6078c7bfeaf518be487")
	//usdt
	asset_addr,_:= common.AddressFromHexString("c6f91c11d740d39943b99a6b1c6fd2b5f476e2a3")
	asset_amount := uint64(100000000)
	//1. transfer asset to user agent contract
	txHash,err :=ontsdk.NeoVM.InvokeNeoVMContract(uint64(2500),uint64(20000),signer,signer,asset_addr,[]interface{}{"transfer",[]interface{}{signer.Address,contractAddr,asset_amount}})
	if err != nil {
		panic(err)
	}
	timeoutSec := 60 * time.Second

	_, err = ontsdk.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		fmt.Printf("error in WaitForGenerateBlock:%s\n", err)

		panic(err)
	}
	//get smartcontract event by txhash
	events, err := ontsdk.GetSmartContractEvent(txHash.ToHexString())
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

	fmt.Println("==== supply ====")
	err = invokeWasm(ontsdk,signer,contractAddr,"supply",[]interface{}{"0x26356Cb66F8fd62c03F569EC3691B6F00173EB02",asset_addr,asset_amount})
	if err != nil {
		panic(err)
	}
}


func Test_QuerySupplyBalance(t *testing.T){
	ontsdk, _, err := InitEnv()
	if err != nil {
		panic(err)
	}

	contractAddr,_ := common.AddressFromHexString(userAgentAddr)


	for _,asset := range asset_infos {
		asset_addr,_:= common.AddressFromHexString(asset.UnderlyingAddress)
		res,err := ontsdk.WasmVM.PreExecInvokeWasmVMContract(contractAddr, "getUserSupplyBalance", []interface{}{asset_addr,"0x26356Cb66F8fd62c03F569EC3691B6F00173EB02"})
		if err != nil {
			panic(err)
		}
		bts,_ := res.Result.ToInteger()
		fmt.Printf("%s : %d\n",asset.AssetName,bts)
	}

}

func Test_QueryContractWingBalacne(t *testing.T){
	ontsdk, _, err := InitEnv()
	if err != nil {
		panic(err)
	}

	asset_addr,_:= common.AddressFromHexString("ff31ec74d01f7b7d45ed2add930f5d2239f7de33")
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)

	wing := oep4.NewOep4(asset_addr,ontsdk)
	b,_:=wing.BalanceOf(contractAddr)

	fmt.Printf("%d\n",b)

}


func Test_QueryWingBalance(t *testing.T){
	ontsdk, wallet, err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount([]byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)

	queryWingBalance(ontsdk,contractAddr)
	//
	//asset_addr,_:= common.AddressFromHexString("7009a2f7c8a2e45fa386a6078c7bfeaf518be487")
	//
	//res,err := ontsdk.WasmVM.PreExecInvokeWasmVMContract(contractAddr, "getUserWingBalance", []interface{}{asset_addr,"0x26356Cb66F8fd62c03F569EC3691B6F00173EB02"})
	//if err != nil {
	//	panic(err)
	//}
	//bts,_ := res.Result.ToInteger()
	//fmt.Println(bts)
}

func queryWingBalance(ontologySdk *sdk.OntologySdk,contractAddr common.Address){
	asset_addr,_:= common.AddressFromHexString("7009a2f7c8a2e45fa386a6078c7bfeaf518be487")
	res,err := ontologySdk.WasmVM.PreExecInvokeWasmVMContract(contractAddr, "getUserWingBalance", []interface{}{asset_addr,"0x26356Cb66F8fd62c03F569EC3691B6F00173EB02"})
	if err != nil {
		panic(err)
	}
	bts,_ := res.Result.ToInteger()
	fmt.Println(bts)
	usdt_addr,_:= common.AddressFromHexString("c6f91c11d740d39943b99a6b1c6fd2b5f476e2a3")
	res,err = ontologySdk.WasmVM.PreExecInvokeWasmVMContract(contractAddr, "getUserWingBalance", []interface{}{usdt_addr,"0x26356Cb66F8fd62c03F569EC3691B6F00173EB02"})
	if err != nil {
		panic(err)
	}
	bts,_ = res.Result.ToInteger()
	fmt.Println(bts)
}


func Test_SetOriginLockProxy(t *testing.T){
	ontsdk, wallet, err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount([]byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)

	setOriginLockProxy(ontsdk,signer,contractAddr)

	res,err := ontsdk.WasmVM.PreExecInvokeWasmVMContract(contractAddr, "getOriginLockProxy", []interface{}{})
	if err != nil {
		panic(err)
	}
	bts,_ := res.Result.ToByteArray()
	fmt.Printf("%v\n",bts)
	addr,err:=common.AddressParseFromBytes(bts)
	if err != nil{
		fmt.Printf("err:%s\n",err.Error())
	}
	fmt.Println(addr.ToHexString())
}




func Test_SetOracle(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	setOracle(ontsdk,signer,contractAddr)
}


func Test_SetWingTokenAddr(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	setWingToken(ontsdk,signer,contractAddr)
}

func setWingToken(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	oralce ,_:= common.AddressFromHexString("ff31ec74d01f7b7d45ed2add930f5d2239f7de33")
	err := invokeWasm(ontsdk,signer,contractAddr,"setWingTokenAddress",[]interface{}{oralce})
	if err != nil {
		panic(err)
	}
}



func Test_Withdraw(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)

	withdraw(ontsdk,signer,contractAddr)
}

func withdraw(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	//eth
	//asset_addr,_:= common.AddressFromHexString("7009a2f7c8a2e45fa386a6078c7bfeaf518be487")
	//usdt
	asset_addr,_:= common.AddressFromHexString("c6f91c11d740d39943b99a6b1c6fd2b5f476e2a3")
	asset_amount := uint64(90000000)

	err := invokeWasm(ontsdk,signer,contractAddr,"withdraw",[]interface{}{"0x26356Cb66F8fd62c03F569EC3691B6F00173EB02",asset_addr,asset_amount})
	if err != nil {
		panic(err)
	}
}
func Test_withdrawWing(t *testing.T){
	ontsdk,wallet,err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount( []byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr,_ := common.AddressFromHexString(userAgentAddr)
	withdrawWing(ontsdk,signer,contractAddr)

}

func withdrawWing(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address){
	err := invokeWasm(ontsdk,signer,contractAddr,"withdrawWing",[]interface{}{"0x26356Cb66F8fd62c03F569EC3691B6F00173EB02",signer.Address})
	if err != nil {
		panic(err)
	}
}



func invokeWasm(ontSdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address,method string,param []interface{})error{

	txHash, err := ontSdk.WasmVM.InvokeWasmVMSmartContract(
		uint64(2500), uint64(2000000), nil, signer, contractAddr, method, param)
	if err != nil {
		fmt.Printf("error in InvokeWasmVMSmartContract:%s\n", err)
		return err
	}
	timeoutSec := 70 * time.Second

	_, err = ontSdk.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		fmt.Printf("error in WaitForGenerateBlock:%s\n", err)

		return err
	}
	//get smartcontract event by txhash
	events, err := ontSdk.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		fmt.Printf("error in GetSmartContractEvent:%s\n", err)
		return err
	}
	fmt.Printf("event is %v\n", events)
	//State = 0 means transaction failed
	if events.State == 0 {
		fmt.Printf("error in events.State is 0 failed.\n")

		return err
	}
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		//you check the notify here
		fmt.Printf("%+v\n", notify)
	}
	return nil
}

func Test_QueryOracle(t *testing.T) {
	ontsdk, wallet, err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount([]byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr, _ := common.AddressFromHexString("b9c72f455281958c3262a0f6f43d7af8d2b60ba8")

	res,err := ontsdk.WasmVM.PreExecInvokeWasmVMContract(contractAddr, "getOracleAddress", []interface{}{})
	if err != nil {
		panic(err)
	}
	bts,_ := res.Result.ToByteArray()
	fmt.Println(hex.EncodeToString(common.ToArrayReverse(bts)))
}

func querySupplyBalance(ontsdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address)*big.Int{
	//asset_addr,_:= common.AddressFromHexString("7009a2f7c8a2e45fa386a6078c7bfeaf518be487")
	usdt_asset_addr,_:= common.AddressFromHexString("c6f91c11d740d39943b99a6b1c6fd2b5f476e2a3")

	res,err := ontsdk.WasmVM.PreExecInvokeWasmVMContract(contractAddr, "getUserSupplyBalance", []interface{}{usdt_asset_addr,"0x26356Cb66F8fd62c03F569EC3691B6F00173EB02"})
	if err != nil {
		panic(err)
	}
	bts,_ := res.Result.ToInteger()
	return bts
}

func Test_Mint(t *testing.T){
	ontsdk, wallet, err := InitEnv()
	if err != nil {
		panic(err)
	}
	//we get the first account of the wallet by your password
	signer, err := wallet.GetDefaultAccount([]byte("123456"))
	if err != nil {
		fmt.Printf("error in GetDefaultAccount:%s\n", err)
		return
	}
	fmt.Printf("===signer address is %s\n", signer.Address.ToBase58())
	contractAddr, _ := common.AddressFromHexString("abccea5e651d3030a8199e9d76e5687b9df89479")
	usdt_ul,_:= common.AddressFromHexString("c6f91c11d740d39943b99a6b1c6fd2b5f476e2a3")
	//eth_ul,_:= common.AddressFromHexString("7009a2f7c8a2e45fa386a6078c7bfeaf518be487")

	wing_gov ,_:= common.AddressFromHexString("6165f58fe4210629c2788ac49f5db437b56122d6")
	usdt := oep4.NewOep4(usdt_ul,ontsdk)
	txhash,err := usdt.Approve(signer,wing_gov,big.NewInt(1000000),signer,uint64(2500),uint64(200000))
	timeoutSec := 60 * time.Second

	_, err = ontsdk.WaitForGenerateBlock(timeoutSec)
	assert.Nil(t,err)
	//get smartcontract event by txhash
	events, err := ontsdk.GetSmartContractEvent(txhash.ToHexString())
	assert.Nil(t,err)

	fmt.Printf("event is %v\n", events)
	//State = 0 means transaction failed
	//assert.NotEqual(t,stats)

	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		//you check the notify here
		fmt.Printf("%+v\n", notify)
	}

	err = invokeWasm(ontsdk,signer,contractAddr,"mint",[]interface{}{signer.Address,uint64(1000000)})
	assert.Nil(t,err)
}