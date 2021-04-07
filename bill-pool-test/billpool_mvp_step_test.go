package bill_pool_test

import (
	"fmt"
	"github.com/ontio/ontology/common"
	"testing"
)
var (
	IsLocal = false
)
//step 1.
func Test_DeployAdmin(t *testing.T){
	fmt.Println("====deploy admin======")
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
	adminAddr,err := deployAdmin(ontsdk,signer,IsLocal)
	if err != nil {
		panic(err)
	}
	fmt.Println("==================init admin========================")
	err = invokeWasm(ontsdk,signer,adminAddr,"init",[]interface{}{signer.Address})
	if err!= nil {
		//panic(err)
		fmt.Println("admin already inited")
	}
}

//step 2.
func Test_DeployCollateral(t *testing.T){
	fmt.Println("====deploy admin======")
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
	cvcAddr,err := deployCollateralVault(ontsdk,signer,IsLocal)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cvcAddr address:%s\n",cvcAddr.ToHexString())
}

//step 3.
func Test_DeployQuoteVault(t *testing.T){
	fmt.Println("====deploy QuoteVault======")
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
	qcvcAddr,err := deployQuoteCurrencylVault(ontsdk,signer,IsLocal)
	if err != nil {
		panic(err)
	}
	fmt.Printf("qcvcAddr address:%s\n",qcvcAddr.ToHexString())
}

//step 4.
func Test_DeployOracle(t *testing.T){
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
	wingOra ,err := deployWingOracle(ontsdk,signer,IsLocal)
	if err != nil {
		panic(err)
	}

	oraAddr,err := deployOracle(ontsdk,signer,IsLocal)
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
}

//step 5.
func Test_DepolyToken(t *testing.T){
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
	//deploy USDT
	usdtaddr,err := deployUSDT(ontsdk,signer,IsLocal)
	if err!= nil {
		panic(err)
	}

	ethaddr,err := deployETH(ontsdk,signer,IsLocal)
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
}

//step 6.
func Test_UpdateOracle(t *testing.T)  {
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

	oraAddr,_ := common.AddressFromHexString("897ff0e53a3411674bf7cb185eb672cb963a5128")
	usdtaddr,_ := common.AddressFromHexString("325efb00f9b3e48fe18811ccde010f78ced83d1e")
	ethaddr,_ := common.AddressFromHexString("b67d6445c950b0e9bd621589671190d119b0f145")

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

}

//step 7.
func Test_DeployBond(t *testing.T){
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
	contractAddr,err := deployBond(ontsdk,signer,true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bond address:%s\n",contractAddr.ToHexString())

}

//step 8.
func Test_InitCollateraVault(t *testing.T){
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

	cvcAddr ,_:=common.AddressFromHexString("93b615581137a2dbc5604fcdb598db66fa2b3aad")
	//adminAddr ,_:=common.AddressFromHexString("6256fee226ef81919ec629a118111bcbfa9bab7a")
	bondAddr ,_:=common.AddressFromHexString("068df9e7f7132e126c37977d890698fad2fce06a")

	//fmt.Println("==================init collateral vault========================")
	//err = invokeWasm(ontsdk,signer,cvcAddr,"init",[]interface{}{signer.Address,bondAddr,adminAddr})
	//if err != nil {
	//	fmt.Println("cvcAddr already initialized")
	//}

	err = invokeWasm(ontsdk,signer,cvcAddr,"updateBondContract",[]interface{}{signer.Address,bondAddr})
	if err != nil {
		panic(err)
	}

}

//step 9.
func Test_SetAdminParam(t *testing.T){
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
	adminAddr ,_:=common.AddressFromHexString("6256fee226ef81919ec629a118111bcbfa9bab7a")
	usdtaddr ,_:=common.AddressFromHexString("325efb00f9b3e48fe18811ccde010f78ced83d1e")
	ethaddr ,_:=common.AddressFromHexString("b67d6445c950b0e9bd621589671190d119b0f145")

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

}

//step 10.
func Test_InitQuoteCurrency(t *testing.T){
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

	qcvcAddr ,_:=common.AddressFromHexString("aeff551998bee4de0ffc964ec930cddd90c72f1a")
	adminAddr ,_:=common.AddressFromHexString("6256fee226ef81919ec629a118111bcbfa9bab7a")
	bondAddr ,_:=common.AddressFromHexString("068df9e7f7132e126c37977d890698fad2fce06a")

	fmt.Println("==================init collateral vault========================")
	err = invokeWasm(ontsdk,signer,qcvcAddr,"init",[]interface{}{signer.Address,bondAddr,adminAddr})
	if err != nil {
		fmt.Println("qcvcAddr already initialized")
	}
	err = invokeWasm(ontsdk,signer,qcvcAddr,"updateBondContract",[]interface{}{signer.Address,bondAddr})
	if err != nil {
		panic(err)
	}
}

//step 11.
func Test_SetBondParams(t *testing.T){
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

	qcvcAddr ,_:=common.AddressFromHexString("aeff551998bee4de0ffc964ec930cddd90c72f1a")
	cvcAddr ,_:=common.AddressFromHexString("93b615581137a2dbc5604fcdb598db66fa2b3aad")
	adminAddr ,_:=common.AddressFromHexString("6256fee226ef81919ec629a118111bcbfa9bab7a")
	bondAddr ,_:=common.AddressFromHexString("068df9e7f7132e126c37977d890698fad2fce06a")
	oraAddr ,_:=common.AddressFromHexString("897ff0e53a3411674bf7cb185eb672cb963a5128")

	fmt.Println("========setCollareralVault===========")
	err = invokeWasm(ontsdk,signer,bondAddr,"setCollareralVault",[]interface{}{cvcAddr})
	if err != nil {
		panic(err)
	}
	fmt.Println("========setQuoteCurrencyValut===========")
	err = invokeWasm(ontsdk,signer,bondAddr,"setQuoteCurrencyValut",[]interface{}{qcvcAddr})
	if err != nil {
		panic(err)
	}

	fmt.Println("========setOracle===========")
	err = invokeWasm(ontsdk,signer,bondAddr,"setOracle",[]interface{}{oraAddr})
	if err != nil {
		panic(err)
	}
	fmt.Println("========setAdmin===========")
	err = invokeWasm(ontsdk,signer,bondAddr,"setAdmin",[]interface{}{adminAddr})
	if err != nil {
		panic(err)
	}
}

//step 12.
func Test_Approve_CollateralVault(t *testing.T){
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

	cvcAddr ,_:=common.AddressFromHexString("93b615581137a2dbc5604fcdb598db66fa2b3aad")
	usdtaddr ,_:=common.AddressFromHexString("325efb00f9b3e48fe18811ccde010f78ced83d1e")
	ethaddr ,_:=common.AddressFromHexString("b67d6445c950b0e9bd621589671190d119b0f145")

	err = invokeNeo(ontsdk,signer,usdtaddr,"approve",[]interface{}{signer.Address,cvcAddr,100000})

	if err!= nil {
		panic(err)
	}
	r,err := ontsdk.NeoVM.PreExecInvokeNeoVMContract(usdtaddr,[]interface{}{"allowance",[]interface{}{signer.Address,cvcAddr}})
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
}

//step 13.
func Test_Mint(t *testing.T){

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
	bondAddr ,_:=common.AddressFromHexString("068df9e7f7132e126c37977d890698fad2fce06a")
	usdtaddr ,_:=common.AddressFromHexString("325efb00f9b3e48fe18811ccde010f78ced83d1e")

	bondid,totalsupply,err := mint(ontsdk,signer,true,bondAddr,usdtaddr,1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("========mint Done bondid:%d,totalsupply:%d===========\n",bondid,totalsupply)
}
//step 14.
func Test_issue(t *testing.T){
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
	bondAddr ,_:=common.AddressFromHexString("068df9e7f7132e126c37977d890698fad2fce06a")
	bondid:= uint64(1)

	err = issueBond(ontsdk,signer,bondAddr,bondid)
	if err != nil {
		panic(err)
	}
}

//step 15.
func Test_AddWhiteList(t *testing.T){
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

	acct2,_ := wallet.GetAccountByAddress("AHoTem8EKxJhsCSMwR5k977vN7t2UWgtyh",[]byte("123456"))

	bondAddr ,_:=common.AddressFromHexString("068df9e7f7132e126c37977d890698fad2fce06a")
	//usdtaddr ,_:=common.AddressFromHexString("325efb00f9b3e48fe18811ccde010f78ced83d1e")
	//qcvcAddr ,_:=common.AddressFromHexString("aeff551998bee4de0ffc964ec930cddd90c72f1a")

	bondid:= uint64(1)
	totalSupply := uint64(100)


	err = invokeWasm(ontsdk,signer,bondAddr,"addSingleWhiteList",[]interface{}{bondid,acct2.Address,1000000,totalSupply})
	if err != nil {
		panic(err)
	}

}

//step 16.
func Test_Buybond(t *testing.T){
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

	acct2,_ := wallet.GetAccountByAddress("AHoTem8EKxJhsCSMwR5k977vN7t2UWgtyh",[]byte("123456"))

	bondAddr ,_:=common.AddressFromHexString("068df9e7f7132e126c37977d890698fad2fce06a")
	usdtaddr ,_:=common.AddressFromHexString("325efb00f9b3e48fe18811ccde010f78ced83d1e")
	qcvcAddr ,_:=common.AddressFromHexString("aeff551998bee4de0ffc964ec930cddd90c72f1a")

	bondid:= uint64(1)
	totalSupply := uint64(100)

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
	//r, err = ontsdk.NeoVM.PreExecInvokeNeoVMContract(usdtaddr,[]interface{}{"allowance",[]interface{}{acct2.Address,qcvcAddr}})
	//if err != nil {
	//	panic(err)
	//}
	//uallowance,_ := r.Result.ToInteger()
	//fmt.Printf("acct2:%s,bond:%s\n",acct2.Address.ToHexString(),qcvcAddr.ToHexString())
	//fmt.Printf("usdt allowance:%d\n",uallowance)


	err = invokeWasm(ontsdk,acct2,bondAddr,"transferByDealWhiteList",[]interface{}{bondid,acct2.Address,totalSupply})
	if err != nil {
		panic(err)
	}

}

//step 17.
func Test_confirm(t *testing.T){
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
	bondAddr ,_:=common.AddressFromHexString("068df9e7f7132e126c37977d890698fad2fce06a")
	bondid:= uint64(1)

	err = invokeWasm(ontsdk,signer,bondAddr,"clearIssuance",[]interface{}{bondid})
	if err != nil {
		panic(err)
	}
}
//step 18.
func Test_RetrieveRaisedCurrency(t *testing.T){
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
	bondAddr ,_:=common.AddressFromHexString("068df9e7f7132e126c37977d890698fad2fce06a")
	bondid:= uint64(1)

	err = invokeWasm(ontsdk,signer,bondAddr,"retrieveRaisedCurrency",[]interface{}{bondid})
	if err != nil {
		panic(err)
	}

}
//step 19.
func Test_IssuerExecute(t *testing.T){
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
	bondAddr ,_:=common.AddressFromHexString("068df9e7f7132e126c37977d890698fad2fce06a")
	bondid:= uint64(1)
	usdtaddr ,_:=common.AddressFromHexString("325efb00f9b3e48fe18811ccde010f78ced83d1e")
	qcvcAddr ,_:=common.AddressFromHexString("aeff551998bee4de0ffc964ec930cddd90c72f1a")
	usdtAmount := 100000000

	err = invokeNeo(ontsdk,signer,usdtaddr,"approve",[]interface{}{signer.Address,qcvcAddr,usdtAmount})
	if err != nil {
		panic(err)
	}


	err = invokeWasm(ontsdk,signer,bondAddr,"issuerExecute",[]interface{}{bondid})
	if err != nil {
		panic(err)
	}
}