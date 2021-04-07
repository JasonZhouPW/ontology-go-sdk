package bill_pool_test

import (
	"fmt"
	"github.com/ontio/ontology/common"
	"testing"
)

func Test_issueExecute(t *testing.T){
	fmt.Println("==================start testing....========================")
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
	bondAddr,_ := common.AddressFromHexString("165c0bd93f98e6f61d31f9861f6463ee37e9a77d")

	err = invokeWasm(ontsdk,signer,bondAddr,"issuerExecute",[]interface{}{2})
	if err != nil {
		panic(err)
	}

}


func Test_Exchange(t *testing.T){
	fmt.Println("==================start testing....========================")
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
	exchangeAddress,err := deployExchange(ontsdk,signer,true)
	if err != nil {
		panic(err)
	}
	bondAddress,_ := common.AddressFromHexString("92d52fdd47d8aa2caaeec2d7fe1d3f23da968bbc")
	//exchangeAddress,_:= common.AddressFromHexString("986c383fb47420f05fc97a50b54847969cb63eb7")
	usdtAddress,_:= common.AddressFromHexString("325efb00f9b3e48fe18811ccde010f78ced83d1e")
	bondid := 21
	totalSupply := 100
	err = invokeWasm(ontsdk,signer,bondAddress,"approve",[]interface{}{signer.Address,exchangeAddress,bondid,totalSupply})
	if err != nil {
		panic(err)
	}
	fmt.Println("======== 2. make order to exchange ==========")
	err = invokeWasm(ontsdk,signer,exchangeAddress,"makeOrder",[]interface{}{signer.Address,
		fmt.Sprintf("%d",bondid),
		bondAddress,
		bondid,
		totalSupply,
		10000000,
		20000000,
		3600,
		0,
		usdtAddress,
	})
	if err != nil {
		panic(err)
	}
}

func Test_bond(t *testing.T){
	fmt.Println("==================start testing....========================")
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

	bondAddr,_:= common.AddressFromHexString("ed4de8105b9d96d92e35440a502c4f75aeba3051")
	r,err := ontsdk.WasmVM.PreExecInvokeWasmVMContract(bondAddr,"getBond",[]interface{}{1})
	if err != nil {
		panic(err)
	}
	b,_:=r.Result.ToByteArray()
	fmt.Printf("%v\n",b)

}