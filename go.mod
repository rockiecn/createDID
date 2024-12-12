module github.com/rockiecn/createDID

go 1.22.2

replace (
	github.com/memoio/contractsv2 => ../memov2-contractsv2
	github.com/memoio/did-solidity => ../did-solidity
)

require (
	github.com/ethereum/go-ethereum v1.10.16
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/memoio/contractsv2 v0.0.0-00010101000000-000000000000
)

require (
	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/go-ole/go-ole v1.2.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/rjeczalik/notify v0.9.1 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	github.com/tklauser/numcpus v0.2.2 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
)
