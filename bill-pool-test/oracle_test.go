package bill_pool_test

import (
	"fmt"
	"testing"
)

func Test_Oracle(t *testing.T){
	ontsdk,wallet,err := InitEnv(IsLocal)
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

	oraAddr,err := deployOracle(ontsdk,signer,true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("oraAddr address:%s\n",oraAddr.ToHexString())

	err = invokeWasm(ontsdk,signer,oraAddr,"init",[]interface{}{signer.Address})
	if err!= nil {
		//panic(err)
		fmt.Println("oracle already inited")
	}

	//deploy USDT
	fmt.Println("========deployUSDT===========")

	usdtaddr,err := deployUSDT(ontsdk,signer,true)
	if err!= nil {
		panic(err)
	}
	_,err = ontsdk.NeoVM.InvokeNeoVMContract(uint64(0),uint64(200000),signer,signer,usdtaddr,[]interface{}{"init",[]interface{}{}})
	if err!= nil {
		panic(err)
	}

	//err = invokeWasm(ontsdk,signer,oraAddr,"updateAssetMeta",[]interface{}{[]interface{}{usdtaddr},
	//	[]interface{}{[]interface{}{"USDT",18}}})
	//if err!= nil {
	//	panic(err)
	//}
	err = invokeWasm(ontsdk,signer,oraAddr,"updateAssetDetail",[]interface{}{usdtaddr,"USDT",18})
	if err!= nil {
		panic(err)
	}

	fmt.Println("=====getAssetMeta===")
	r,err := ontsdk.WasmVM.PreExecInvokeWasmVMContract(oraAddr,"getAssetMeta",[]interface{}{usdtaddr})
	if err!= nil {
		panic(err)
	}

	fmt.Println(r.Result.ToByteArray())


}
