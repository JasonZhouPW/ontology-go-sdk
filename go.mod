module github.com/ontio/ontology-go-sdk

go 1.12

require (
	github.com/FactomProject/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/cmars/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/gorilla/websocket v1.4.1
	github.com/itchyny/base58-go v0.1.0
	github.com/kr/pretty v0.2.0 // indirect
	github.com/ontio/go-bip32 v0.0.0-20190520025953-d3cea6894a2b
	github.com/ontio/ontology v1.11.0
	github.com/ontio/ontology-crypto v1.0.9
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.4.0
	github.com/tyler-smith/go-bip39 v1.0.2
	golang.org/x/crypto v0.0.0-20200311171314-f7b00557c8c4
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)

replace (
	labix.org/v2/mgo => github.com/go-mgo/mgo v0.0.0-20160801194620-b6121c6199b7
	launchpad.net/gocheck => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
)
