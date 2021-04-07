package bill_pool_test

import (
	"encoding/hex"
	"fmt"
	sdk "github.com/ontio/ontology-go-sdk"
	bill_pool "github.com/ontio/ontology-go-sdk/bill-pool-test"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"math/big"
	"strconv"
	"testing"
	"time"
)

var gasprice uint64 = 0

func InitEnv(islocal bool) (*sdk.OntologySdk ,*sdk.Wallet,error){

	fmt.Println("==========================start============================")
	testUrl := "http://127.0.0.1:20336"

	if !islocal {
		testUrl = "http://polaris2.ont.io:20336"
	}

	if islocal{
		gasprice = 0
	}else{
		gasprice = 2500
	}

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




func Test_Billpool(t *testing.T) {
	fmt.Println("==================start testing....========================")
	ontsdk,wallet,err := InitEnv(true)
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

	adminAddr,err := deployAdmin(ontsdk,signer,true)
	if err != nil {
		panic(err)
	}
	fmt.Println("==================init admin========================")
	err = invokeWasm(ontsdk,signer,adminAddr,"init",[]interface{}{signer.Address})
	if err!= nil {
		//panic(err)
		fmt.Println("admin already inited")
	}

	cvcAddr,err := deployCollateralVault(ontsdk,signer,true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cvcAddr address:%s\n",cvcAddr.ToHexString())

	qcvcAddr,err := deployQuoteCurrencylVault(ontsdk,signer,true)
	if err != nil {
		panic(err)
	}


	wingOra ,err := deployWingOracle(ontsdk,signer,true)
	if err != nil {
		panic(err)
	}

	oraAddr,err := deployOracle(ontsdk,signer,true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("oraAddr address:%s\n",oraAddr.ToHexString())
	fmt.Println("==================init oracle========================")
	err = invokeWasm(ontsdk,signer,oraAddr,"init",[]interface{}{signer.Address})
	if err!= nil {
		//panic(err)
		fmt.Println("oracle already inited")
	}
	fmt.Println("==================set oracle updateWingOracle========================")
	err = invokeWasm(ontsdk,signer,oraAddr,"updateWingOracle",[]interface{}{wingOra})
	if err!= nil {
		panic(err)
	}

	//deploy USDT
	usdtaddr,err := deployUSDT(ontsdk,signer,true)
	if err!= nil {
		panic(err)
	}

	ethaddr,err := deployETH(ontsdk,signer,true)
	if err!= nil {
		panic(err)
	}
	fmt.Println("==================init usdt========================")
	err = invokeNeo(ontsdk,signer,usdtaddr,"init",[]interface{}{})
	//_,err = ontsdk.NeoVM.InvokeNeoVMContract(uint64(0),uint64(200000),signer,signer,usdtaddr,[]interface{}{"init",[]interface{}{}})
	if err!= nil {
		panic(err)
	}
	fmt.Println("==================init eth========================")
	err = invokeNeo(ontsdk,signer,ethaddr,"init",[]interface{}{})
	if err!= nil {
		panic(err)
	}
	fmt.Println("==================updateAssetDetail oracle usdt========================")

	err = invokeWasm(ontsdk,signer,oraAddr,"updateAssetDetail",[]interface{}{usdtaddr,"USDT",6})
	if err!= nil {
		panic(err)
	}
	fmt.Println("==================updateAssetDetail oracle eth========================")
	err = invokeWasm(ontsdk,signer,oraAddr,"updateAssetDetail",[]interface{}{ethaddr,"ETH",18})
	if err!= nil {
		panic(err)
	}

	exchangeAddr,err := deployExchange(ontsdk,signer,true)
	if err != nil {
		panic(err)
	}

	contractAddr,err := deployBond(ontsdk,signer,true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bond address:%s\n",contractAddr.ToHexString())

	fmt.Println("==================init collateral vault========================")
	err = invokeWasm(ontsdk,signer,cvcAddr,"init",[]interface{}{signer.Address,contractAddr,adminAddr})
	if err != nil {
		fmt.Println("cvcAddr already initialized")
	}
	fmt.Println("==================update collateral adjustQuoteCurrency usdt========================")

	err = invokeWasm(ontsdk,signer,adminAddr,"adjustQuoteCurrency",[]interface{}{usdtaddr,true})
	if err!= nil {
		panic(err)
	}
	fmt.Println("==================update collateral adjustQuoteCurrency eth========================")

	err = invokeWasm(ontsdk,signer,adminAddr,"adjustQuoteCurrency",[]interface{}{ethaddr,true})
	if err!= nil {
		panic(err)
	}
	fmt.Println("==================update collateral adjustCollateral usdt========================")

	err = invokeWasm(ontsdk,signer,adminAddr,"adjustCollateral",[]interface{}{usdtaddr,true})
	if err!= nil {
		panic(err)
	}
	fmt.Println("==================update collateral adjustCollateral eth========================")

	err = invokeWasm(ontsdk,signer,adminAddr,"adjustCollateral",[]interface{}{ethaddr,true})
	if err!= nil {
		panic(err)
	}
	fmt.Println("==================update collateral updateCollateralFactor eth========================")

	err = invokeWasm(ontsdk,signer,adminAddr,"updateCollateralFactor",[]interface{}{ethaddr,1000000000000000000})  //100%
	if err!= nil {
		panic(err)
	}
	fmt.Println("==================update collateral updateCollateralFactor usdt========================")

	err = invokeWasm(ontsdk,signer,adminAddr,"updateCollateralFactor",[]interface{}{usdtaddr,1000000000000000000})  //100%
	if err!= nil {
		panic(err)
	}

	r,err := ontsdk.WasmVM.PreExecInvokeWasmVMContract(adminAddr,"isSupportedCollateral",[]interface{}{usdtaddr})
	if err!= nil {
		panic(err)
	}
	f,err := r.Result.ToBool()

	fmt.Printf("isSupportedCollateral:%s,%v\n",usdtaddr.ToHexString(),f)

	r,err = ontsdk.WasmVM.PreExecInvokeWasmVMContract(adminAddr,"isSupportedCollateral",[]interface{}{ethaddr})
	if err!= nil {
		panic(err)
	}
	f,err = r.Result.ToBool()

	fmt.Printf("isSupportedCollateral:%s,%v\n",ethaddr.ToHexString(),f)

	r,err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(usdtaddr,[]interface{}{"balanceOf",[]interface{}{signer.Address}})
	if err!= nil {
		panic(err)
	}
	b,_:=r.Result.ToInteger()

	fmt.Printf("balance of usdt:%d\n",b)

	err = invokeNeo(ontsdk,signer,usdtaddr,"approve",[]interface{}{signer.Address,cvcAddr,100000})

	if err!= nil {
		panic(err)
	}
	r,err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(usdtaddr,[]interface{}{"allowance",[]interface{}{signer.Address,cvcAddr}})
	if err!= nil {
		panic(err)
	}
	i,_:=r.Result.ToInteger()
	fmt.Printf("allowance of usdt:%d\n",i)

	err = invokeNeo(ontsdk,signer,ethaddr,"approve",[]interface{}{signer.Address,cvcAddr,1000})
	if err!= nil {
		panic(err)
	}
	r,err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(ethaddr,[]interface{}{"allowance",[]interface{}{signer.Address,cvcAddr}})
	if err!= nil {
		panic(err)
	}
	i,_=r.Result.ToInteger()
	fmt.Printf("allowance of eth:%d\n",i)

	err = invokeWasm(ontsdk,signer,cvcAddr,"depositAsset",[]interface{}{signer.Address,usdtaddr,0,0,100000})
	if err!= nil {
		panic(err)
	}
	err = invokeWasm(ontsdk,signer,cvcAddr,"depositAsset",[]interface{}{signer.Address,ethaddr,0,0,1000})
	if err!= nil {
		panic(err)
	}

	fmt.Println("========setCollareralVault===========")
	err = invokeWasm(ontsdk,signer,contractAddr,"setCollareralVault",[]interface{}{cvcAddr})
	if err != nil {
		panic(err)
	}

	fmt.Println("========init qcvc===========")
	err = invokeWasm(ontsdk,signer,qcvcAddr,"init",[]interface{}{signer.Address,contractAddr,adminAddr})
	if err != nil {
		fmt.Println("qcvcAddr already initialized")
	}

	fmt.Println("========setQuoteCurrencyValut===========")
	err = invokeWasm(ontsdk,signer,contractAddr,"setQuoteCurrencyValut",[]interface{}{qcvcAddr})
	if err != nil {
		panic(err)
	}

	fmt.Println("========setOracle===========")
	err = invokeWasm(ontsdk,signer,contractAddr,"setOracle",[]interface{}{oraAddr})
	if err != nil {
		panic(err)
	}
	fmt.Println("========setExchange===========")
	err = invokeWasm(ontsdk,signer,contractAddr,"setExchange",[]interface{}{exchangeAddr})
	if err != nil {
		panic(err)
	}

	fmt.Println("========init exchange===========")
	err = invokeWasm(ontsdk,signer,exchangeAddr,"init",[]interface{}{signer.Address,contractAddr})
	if err != nil {
		fmt.Printf("already initialized")
	}

	fmt.Println("========setAdmin===========")
	err = invokeWasm(ontsdk,signer,contractAddr,"setAdmin",[]interface{}{adminAddr})
	if err != nil {
		panic(err)
	}

	fmt.Println("========mint===========")
	tmp := uint64(time.Now().Unix())
	fmt.Printf("%d\n",tmp)
	fmt.Printf("usdt addr:%s\n",usdtaddr.ToHexString())
	bondid,totalsupply,err := mint(ontsdk,signer,true,contractAddr,usdtaddr,1)

	if err != nil {
		panic(err)
	}
	fmt.Printf("========mint Done bondid:%d===========\n",bondid)

	//todo
	fmt.Println("======== public sale ==========")
	fmt.Println("======== 1. approve bond to exchange ==========")
	err = invokeWasm(ontsdk,signer,contractAddr,"approve",[]interface{}{signer.Address,exchangeAddr,bondid,totalsupply})
	if err != nil {
		panic(err)
	}


	fmt.Printf("======== 2. query bond with id:%d ==========\n",bondid)
	r,err = ontsdk.WasmVM.PreExecInvokeWasmVMContract(contractAddr,"getBond",[]interface{}{bondid})
	if err != nil {
		panic(err)
	}
	ba,_:=r.Result.ToByteArray()
	fmt.Printf("%v\n",ba)

	fmt.Println("======== 3. make order to exchange ==========")
	orderid,err := makeOrder(ontsdk,signer,true,exchangeAddr,contractAddr,bondid,totalsupply,usdtaddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("======== 4. take order on exchange ==========")
	acct2,_ := wallet.GetAccountByAddress("AHoTem8EKxJhsCSMwR5k977vN7t2UWgtyh",[]byte("123456"))

	fmt.Println("======== 4.1  transfer some assets from signer ==========")
	usdtamt := 1000000*1000000
	err = invokeNeo(ontsdk,signer,usdtaddr,"transfer",[]interface{}{signer.Address,acct2.Address,10*usdtamt})
	if err != nil {
		panic(err)
	}
	ethamt :=new(big.Int).Mul(big.NewInt(1000),big.NewInt(1000000000000000000))
	err = invokeNeo(ontsdk,signer,ethaddr,"transfer",[]interface{}{signer.Address,acct2.Address,new(big.Int).Mul(big.NewInt(10),ethamt)})
	if err != nil {
		panic(err)
	}
	r, err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(usdtaddr,[]interface{}{"balanceOf",[]interface{}{acct2.Address}})
	if err != nil {
		panic(err)
	}
	ubalance,_ := r.Result.ToInteger()
	fmt.Printf("usdt balance:%d\n",ubalance)

	r, err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(ethaddr,[]interface{}{"balanceOf",[]interface{}{acct2.Address}})
	if err != nil {
		panic(err)
	}
	ebalance,_ := r.Result.ToInteger()
	fmt.Printf("eth balance:%d\n",ebalance)

	fmt.Println("======== 4.2  approve some assets to exchange ==========")
	err = invokeNeo(ontsdk,acct2,usdtaddr,"approve",[]interface{}{acct2.Address,exchangeAddr,usdtamt})
	if err != nil {
		panic(err)
	}
	err = invokeNeo(ontsdk,acct2,ethaddr,"approve",[]interface{}{acct2.Address,exchangeAddr,ethamt})
	if err != nil {
		panic(err)
	}

	r, err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(usdtaddr,[]interface{}{"allowance",[]interface{}{acct2.Address,exchangeAddr}})
	if err != nil {
		panic(err)
	}
	uallowance,_ := r.Result.ToInteger()
	fmt.Printf("usdt allowance:%d\n",uallowance)

	r, err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(ethaddr,[]interface{}{"allowance",[]interface{}{acct2.Address,exchangeAddr}})
	if err != nil {
		panic(err)
	}
	eallowance,_ := r.Result.ToInteger()
	fmt.Printf("eth allowance:%d\n",eallowance)


	fmt.Println("======== 4.3  take order ==========")
	err = invokeWasm(ontsdk,acct2,exchangeAddr,"takeOrder",[]interface{}{acct2.Address,bondid,orderid,totalsupply})
	if err != nil {
		panic(err)
	}

	fmt.Println("======== 5. mint another bond ==========")
	bondid2,totalsupply2,err := mint(ontsdk,signer,true,contractAddr,usdtaddr,0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("========mint Done bondid:%d===========\n",bondid2)
	fmt.Println("========6.1 add white list of acc2 ==========")
	err = invokeWasm(ontsdk,signer,contractAddr,"addSingleWhiteList",[]interface{}{bondid2,acct2.Address,1000000,totalsupply2})
	if err != nil {
		panic(err)
	}
	fmt.Println("========6.2 buy bond by acct2 ==========")
	err = invokeNeo(ontsdk,acct2,usdtaddr,"approve",[]interface{}{acct2.Address,qcvcAddr,usdtamt})
	if err != nil {
		panic(err)
	}
	r, err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(usdtaddr,[]interface{}{"allowance",[]interface{}{acct2.Address,qcvcAddr}})
	if err != nil {
		panic(err)
	}
	uallowance,_ = r.Result.ToInteger()
	fmt.Printf("acct2:%s,bond:%s\n",acct2.Address.ToHexString(),qcvcAddr.ToHexString())
	fmt.Printf("usdt allowance:%d\n",uallowance)



	err = invokeWasm(ontsdk,acct2,contractAddr,"transferByDealWhiteList",[]interface{}{bondid2,acct2.Address,totalsupply2})
	if err != nil {
		panic(err)
	}

	fmt.Println("======== 7. pay interest ==========")
	fmt.Println("======== 7.1 owner approve to qcvc ==========")
	err = invokeNeo(ontsdk,signer,usdtaddr,"approve",[]interface{}{signer.Address,qcvcAddr,usdtamt})
	if err != nil {
		panic(err)
	}
	r, err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(usdtaddr,[]interface{}{"allowance",[]interface{}{signer.Address,qcvcAddr}})
	if err != nil {
		panic(err)
	}
	uallowance,_ = r.Result.ToInteger()
	fmt.Printf("signer:%s,qcvc:%s\n",acct2.Address.ToHexString(),qcvcAddr.ToHexString())
	fmt.Printf("usdt allowance:%d\n",uallowance)

	fmt.Println("======== 7.2 owner pay interest bond1 ==========")
	err = invokeWasm(ontsdk,signer,contractAddr,"payInterest",[]interface{}{signer.Address,bondid})
	if err != nil {
		panic(err)
	}
	fmt.Println("======== 7.2 owner pay interest bond2==========")
	err = invokeWasm(ontsdk,signer,contractAddr,"payInterest",[]interface{}{signer.Address,bondid2})
	if err != nil {
		panic(err)
	}

	fmt.Println("======== 8 holder execute ==========")
	err = invokeWasm(ontsdk,acct2,contractAddr,"holderExecute",[]interface{}{acct2.Address,bondid})
	if err != nil {
		panic(err)
	}

	fmt.Printf("=======9.issuer execute ")
	time.Sleep(300*time.Second)

	err = invokeWasm(ontsdk,signer,contractAddr,"issuerExecute",[]interface{}{bondid2})
	if err != nil {
		panic(err)
	}



	fmt.Println("======== All Done ==========")
}

func Test_Bond2(t *testing.T){

	fmt.Println("==================start testing....========================")
	ontsdk,wallet,err := InitEnv(true)
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

	acct2,_ := wallet.GetAccountByAddress("AHoTem8EKxJhsCSMwR5k977vN7t2UWgtyh",[]byte("123456"))
	bondAddr,_ := common.AddressFromHexString("045cb4ff30f477e630b41451fdfbd0f25d55240d")
	bondid := 1

	fmt.Printf("======== 2. query bond with id:%d ==========\n",bondid)
	r,err := ontsdk.WasmVM.PreExecInvokeWasmVMContract(bondAddr,"getBond",[]interface{}{bondid})
	if err != nil {
		panic(err)
	}
	ba,_:=r.Result.ToByteArray()
	fmt.Printf("%v\n",ba)

	fmt.Printf("=======9.issuer execute ")
	err = invokeWasm(ontsdk,acct2,bondAddr,"issuerExecute",[]interface{}{bondid})
	if err != nil {
		panic(err)
	}
}



func makeOrder(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool,exchangeAddr common.Address,bondAddr common.Address,bondid uint64,amount uint64,assetaddr common.Address)(uint64, error){
	txHash,err := ontSdk.WasmVM.InvokeWasmVMSmartContract(gasprice, uint64(200000),signer,signer,exchangeAddr,"makeOrder",[]interface{}{signer.Address,
		bondid,
		bondAddr,
		bondid,
		amount,
		10000000,
		20000000,
		3600,
		0,
		assetaddr,
	})
	if err != nil {
		panic(err)
	}

	timeoutSec := 30 * time.Second

	_, err = ontSdk.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		fmt.Printf("error in WaitForGenerateBlock:%s\n", err)

		return 0,err
	}
	//get smartcontract event by txhash
	events, err := ontSdk.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		fmt.Printf("error in GetSmartContractEvent:%s\n", err)
		return 0,err
	}
	fmt.Printf("event is %v\n", events)
	//State = 0 means transaction failed
	if events.State == 0 {
		fmt.Printf("error in events.State is 0 failed.\n")

		return 0,err
	}
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		//you check the notify here
		fmt.Printf("%+v\n", notify)
		if len(notify.States.([]interface{})) > 0 {
			m := notify.States.([]interface{})[0].(string)
			if m == "MakeOrder" {
				id,err := strconv.Atoi( notify.States.([]interface{})[3].(string))
				if err != nil {
					return 0,err
				}
				return  uint64(id),nil
			}
		}
	}
	return 0,fmt.Errorf("no bondmint event")


}


func mint(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool,bondaddr common.Address,assetAddr common.Address,btype uint64)(uint64,uint64,error){
	tmp := uint64(time.Now().Unix())

	txHash, err := ontSdk.WasmVM.InvokeWasmVMSmartContract(
		gasprice, uint64(200000), nil, signer, bondaddr, "mint", []interface{}{
			signer.Address, //user_addr
			"test1", //name
			"test1", //symbol
			btype,       //bond_type
			tmp,     //created_time
			tmp+60*2,     //raised_time
			tmp+60*2,     //interest_start_time
			tmp+60*3,     // subscriber_redemption_time
			tmp+60*4,     //issuer_redemption_time
			tmp+60*4,     //liquidation_time
			2,      //interest_paymentnum
			tmp,     //last_interest_payment_time
			0,      //p_asset_type
			assetAddr, //p_address
			1000000,     //p_amount
			10,      //interest_rate
			100})    //total_supply)
	if err != nil {
		fmt.Printf("error in InvokeWasmVMSmartContract:%s\n", err)
		return 0,0,err
	}
	fmt.Printf("mint time:%d\n",time.Now().Unix())
	timeoutSec := 30 * time.Second

	_, err = ontSdk.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		fmt.Printf("error in WaitForGenerateBlock:%s\n", err)

		return 0,0,err
	}
	//get smartcontract event by txhash
	events, err := ontSdk.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		fmt.Printf("error in GetSmartContractEvent:%s\n", err)
		return 0,0,err
	}
	fmt.Printf("event is %v\n", events)
	//State = 0 means transaction failed
	if events.State == 0 {
		fmt.Printf("error in events.State is 0 failed.\n")

		return 0,0,err
	}
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		//you check the notify here
		fmt.Printf("%+v\n", notify)
		if len(notify.States.([]interface{})) > 0 {
			m := notify.States.([]interface{})[0].(string)
			if m == "BondMint" {
				id,err := strconv.Atoi( notify.States.([]interface{})[2].(string))
				if err != nil {
					return 0,0,err
				}
				total,err := strconv.Atoi(notify.States.([]interface{})[5].(string))
				return uint64(id),uint64(total),nil
			}
		}
	}
	return 0,0,fmt.Errorf("no bondmint event")
}


func deployAdmin(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool)(common.Address,error){
	fmt.Println("=============deployAdmin============")
	wasmfile := "./admin_opt.wasm"
	return deployWasm(ontSdk,signer,local,wasmfile)
}

func deployBond(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool)(common.Address,error){
	fmt.Println("=============deployBond============")
	wasmfile := "./bond_opt.wasm"
	return deployWasm(ontSdk,signer,local,wasmfile)
}

func deployCollateralVault(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool) (common.Address,error) {
	fmt.Println("=============deployCollateralVault============")
	wasmfile := "./collateral_vault_opt.wasm"
	return deployWasm(ontSdk,signer,local,wasmfile)
}
func deployQuoteCurrencylVault(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool) (common.Address,error) {
	fmt.Println("=============deployQuoteCurrencylVault============")
	wasmfile := "./quote_currency_vault_opt.wasm"
	return deployWasm(ontSdk,signer,local,wasmfile)
}

func deployOracle(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool) (common.Address,error) {
	fmt.Println("=============deployOracle============")
	wasmfile := "./oracle_opt.wasm"
	return deployWasm(ontSdk,signer,local,wasmfile)
}

func deployUSDT(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool)(common.Address,error){
	fmt.Println("=============deployUSDT============")
	return deployNeoContract(ontSdk,signer,bill_pool.USDT_CONTRACT,local)
}

func deployBTC(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool)(common.Address,error){
	fmt.Println("=============deployBTC============")
	return deployNeoContract(ontSdk,signer,bill_pool.BTC_CONTRACT,local)
}

func deployETH(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool)(common.Address,error){
	fmt.Println("=============deployETH============")
	return deployNeoContract(ontSdk,signer,bill_pool.ETH_CONTRACT,local)
}

func deployWingOracle(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool) (common.Address,error) {
	fmt.Println("=============deployWingOracle============")
	wasmfile := "./wing_oracle_opt.wasm"
	return deployWasm(ontSdk,signer,local,wasmfile)
}

func deployExchange(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool) (common.Address,error) {
	fmt.Println("=============deployExchange============")
	wasmfile := "./exchange_opt.wasm"
	return deployWasm(ontSdk,signer,local,wasmfile)
}

func deployWasm(ontSdk *sdk.OntologySdk,signer *sdk.Account,local bool,wasmfile string) (common.Address,error) {
	//get a compiled wasm file from ont_cpp

	//set timeout
	timeoutSec := 30 * time.Second
	//address1 := "AX8opZCQBpEpYsFPKpZHNguWz2s3xpT7Wk"

	// read wasm file and get the Hex fmt string
	code, err := ioutil.ReadFile(wasmfile)
	if err != nil {
		fmt.Printf("error in ReadFile:%s\n", err)

		return common.Address{},err
	}

	codeHash := common.ToHexString(code)
	//calculate the contract address from code
	contractAddr, err := utils.GetContractAddress(codeHash)
	if err != nil {
		fmt.Printf("error in GetContractAddress:%s\n", err)

		return common.Address{},err

	}
	fmt.Printf("the contractAddr is %s\n", contractAddr.ToBase58())
	fmt.Printf("the contractAddr is %s\n", contractAddr.ToHexString())
	fmt.Printf("the revert contractAddr is %s\n", hex.EncodeToString(contractAddr[:]))

	//===========================================
	//gasprice := uint64(2500)
	//if local{
	//	gasprice = uint64(0)
	//}
	//invokegaslimit := uint64(200000)
	deploygaslimit := uint64(2000000000)
	// deploy the wasm contract
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

		return common.Address{},err

	}
	_, err = ontSdk.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		fmt.Printf("error in WaitForGenerateBlock:%s\n", err)

		return common.Address{},err
	}
	fmt.Printf("the deploy contract txhash is %s\n", txHash.ToHexString())

	fmt.Println("============Done===============")

	return contractAddr,nil
}

func invokeWasm(ontSdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address,method string,param []interface{})error{


	txHash, err := ontSdk.WasmVM.InvokeWasmVMSmartContract(
		gasprice, uint64(200000), nil, signer, contractAddr, method, param)
	if err != nil {
		fmt.Printf("error in InvokeWasmVMSmartContract:%s\n", err)
		return err
	}
	timeoutSec := 30 * time.Second

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
	fmt.Printf("events.Notify:%v\n", events.Notify)
	for _, notify := range events.Notify {
		//you check the notify here
		fmt.Printf("%+v\n", notify)
	}
	return nil
}

func invokeNeo(ontSdk *sdk.OntologySdk,signer *sdk.Account,contractAddr common.Address,method string,param []interface{})error{
	txHash,err := ontSdk.NeoVM.InvokeNeoVMContract(gasprice,uint64(20000),signer,signer,contractAddr,[]interface{}{method,param})
	if err != nil {
		fmt.Printf("error in InvokeNeoVMContract:%s\n", err)
		return err
	}
	timeoutSec := 30 * time.Second

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
	fmt.Printf("events.Notify:%v\n", events.Notify)
	for _, notify := range events.Notify {
		//you check the notify here
		fmt.Printf("%+v\n", notify)
	}
	return nil
}


func deployNeoContract(ontSdk *sdk.OntologySdk,signer *sdk.Account,contract string,local bool) (common.Address,error) {

	gasprice := uint64(2500)

	if local{
		gasprice = uint64(0)
	}
	deploygaslimit := uint64(200000000)

	txHash, err := ontSdk.NeoVM.DeployNeoVMSmartContract(
		gasprice,
		deploygaslimit,
		signer,
		true,
		contract,
		"neovm",
		"1.0",
		"author",
		"email",
		"desc",
	)
	if err != nil {
		fmt.Printf("error in DeployNeoVMSmartContract:%s\n", err)

		return common.Address{},err
	}
	_, err = ontSdk.WaitForGenerateBlock(60*time.Second)
	if err != nil {
		fmt.Printf("error in WaitForGenerateBlock:%s\n", err)

		return common.Address{},err
	}
	fmt.Printf("the deploy contract txhash is %s\n", txHash.ToHexString())

	//calculate the contract address from code
	contractAddr, err := utils.GetContractAddress(contract)
	if err != nil {
		fmt.Printf("error in GetContractAddress:%s\n", err)

		return common.Address{},err
	}
	fmt.Printf("the contractAddr is %s\n", contractAddr.ToBase58())
	fmt.Printf("the contractAddr is %s\n", contractAddr.ToHexString())
	fmt.Printf("the reversed contractAddr is %s\n",hex.EncodeToString(contractAddr[:]))

	fmt.Println("============Done===============")
	return contractAddr,nil
}