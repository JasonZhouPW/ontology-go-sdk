package bill_pool_test

import (
	"fmt"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"testing"
	"time"
)

func Test_MVP(t *testing.T){
	fmt.Println("mvp test start...")
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
	//fmt.Println("========setExchange===========")
	//err = invokeWasm(ontsdk,signer,contractAddr,"setExchange",[]interface{}{exchangeAddr})
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println("========init exchange===========")
	//err = invokeWasm(ontsdk,signer,exchangeAddr,"init",[]interface{}{signer.Address,contractAddr})
	//if err != nil {
	//	fmt.Printf("already initialized")
	//}

	fmt.Println("========setAdmin===========")
	err = invokeWasm(ontsdk,signer,contractAddr,"setAdmin",[]interface{}{adminAddr})
	if err != nil {
		panic(err)
	}

	fmt.Println("========mint===========")
	tmp := uint64(time.Now().Unix())
	fmt.Printf("%d\n",tmp)
	//fmt.Printf("usdt addr:%s\n",usdtaddr.ToHexString())
	//bondid,totalsupply,err := mint(ontsdk,signer,true,contractAddr,usdtaddr,1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("========mint Done bondid:%d,totalsupply:%d===========\n",bondid,totalsupply)
	bondid2,totalsupply2,err := mint(ontsdk,signer,true,contractAddr,usdtaddr,1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("========mint Done bondid2:%d,totalsupply2:%d===========\n",bondid2,totalsupply2)
	//testCancel := true
	//if testCancel {
	//	fmt.Println("========cancel mint==========")
	//
	//	err = cancelMint(ontsdk,signer,contractAddr,bondid)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	fmt.Println("========issue==========")
	err = issueBond(ontsdk,signer,contractAddr,bondid2)
	if err != nil {
		panic(err)
	}
	fmt.Println("========issue done==========")
	fmt.Println("========add white list of acc2 ==========")
	acct2,_ := wallet.GetAccountByAddress("AHoTem8EKxJhsCSMwR5k977vN7t2UWgtyh",[]byte("123456"))
	err = invokeWasm(ontsdk,signer,contractAddr,"addSingleWhiteList",[]interface{}{bondid2,acct2.Address,1000000,totalsupply2})
	if err != nil {
		panic(err)
	}
	fmt.Println("========transfer some assets from signer ==========")
	usdtamt := 1000000*1000000
	err = invokeNeo(ontsdk,signer,usdtaddr,"transfer",[]interface{}{signer.Address,acct2.Address,10*usdtamt})
	if err != nil {
		panic(err)
	}
	fmt.Println("======== bond by acct2 ==========")
	err = invokeNeo(ontsdk,acct2,usdtaddr,"approve",[]interface{}{acct2.Address,qcvcAddr,usdtamt})
	if err != nil {
		panic(err)
	}
	r, err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(usdtaddr,[]interface{}{"allowance",[]interface{}{acct2.Address,qcvcAddr}})
	if err != nil {
		panic(err)
	}
	uallowance,_ := r.Result.ToInteger()
	fmt.Printf("acct2:%s,bond:%s\n",acct2.Address.ToHexString(),qcvcAddr.ToHexString())
	fmt.Printf("usdt allowance:%d\n",uallowance)


	err = invokeWasm(ontsdk,acct2,contractAddr,"transferByDealWhiteList",[]interface{}{bondid2,acct2.Address,totalsupply2})
	if err != nil {
		panic(err)
	}

	err = invokeWasm(ontsdk,signer,contractAddr,"clearIssuance",[]interface{}{bondid2})
	if err != nil {
		panic(err)
	}

	err = invokeWasm(ontsdk,signer,contractAddr,"retrieveRaisedCurrency",[]interface{}{bondid2})
	if err != nil {
		panic(err)
	}

	time.Sleep(4 *time.Minute)
	//err = invokeNeo(ontsdk,signer,usdtaddr,"approve",[]interface{}{signer.Address,qcvcAddr,usdtamt})
	//if err != nil {
	//	panic(err)
	//}
	//err = invokeWasm(ontsdk,signer,contractAddr,"issuerExecute",[]interface{}{bondid2})
	//if err != nil {
	//	panic(err)
	//}
	//err = invokeWasm(ontsdk,acct2,contractAddr,"retrievePrincipal",[]interface{}{acct2.Address,bondid2})
	//if err != nil {
	//	panic(err)
	//}

	err = invokeWasm(ontsdk,acct2,contractAddr,"liquidate",[]interface{}{bondid2,acct2.Address})
	if err != nil {
		panic(err)
	}

}

func issueBond(ontSdk *ontology_go_sdk.OntologySdk,signer *sdk.Account,bondaddr common.Address,   bondid uint64)error  {
	txHash, err := ontSdk.WasmVM.InvokeWasmVMSmartContract(
		gasprice, uint64(200000), nil, signer, bondaddr, "issue", []interface{}{
			bondid,
		})
	if err != nil {
		fmt.Printf("error in InvokeWasmVMSmartContract:%s\n", err)
		return err
	}
	fmt.Printf("mint time:%d\n",time.Now().Unix())
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
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		//you check the notify here
		fmt.Printf("%+v\n", notify)

	}
	return nil
}

func cancelMint(ontSdk *ontology_go_sdk.OntologySdk,signer *sdk.Account,bondaddr common.Address,   bondid uint64)error  {
	txHash, err := ontSdk.WasmVM.InvokeWasmVMSmartContract(
		gasprice, uint64(200000), nil, signer, bondaddr, "cancelMint", []interface{}{
			bondid,
		})
	if err != nil {
		fmt.Printf("error in InvokeWasmVMSmartContract:%s\n", err)
		return err
	}
	fmt.Printf("mint time:%d\n",time.Now().Unix())
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
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		//you check the notify here
		fmt.Printf("%+v\n", notify)

	}
	return nil
}
