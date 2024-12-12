package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	com "github.com/memoio/contractsv2/common"

	inst "github.com/memoio/contractsv2/go_contracts/instance"
	"github.com/rockiecn/createDID/database"
	"github.com/rockiecn/createDID/go-contracts/proxy"

	"encoding/hex"
)

var (

	//EcdsaSecp256k1VerificationKey2019  :  pubkey - eth public key
	//EcdsaSecp256k1RecoveryMethod2020:  pubkey - eth address -> bytes 20位
	//Ed25519VerificationKey2018： solana,ton:  pubkey - public key 32位
	// chain related info
	methodType = "EcdsaSecp256k1RecoveryMethod2020"

	// admin to send tx
	AdminSK   = "6ec7e0cdda802a466401c912b0dac5ff6116e4372746a455bb43c61294e6f01f"
	AdminAddr = "0xe0DA2108c52C799F27B02D5AF049374B294f291e"
)

func main() {
	inputeth := flag.String("eth", "dev", "eth api Address;") //dev test or product
	psk := flag.String("sk", "", "signature for sending transaction")
	pstart := flag.Uint64("start", 1, "start id")

	flag.Parse()

	chain := *inputeth
	sk := *psk
	start := *pstart

	// open db
	database.Open()
	defer database.DB.Close()

	var id uint64
	for id = start; id <= 467683; id++ {
		go func(id uint64) {
			// read an user from db
			fmt.Println("now reading user: ", id)
			user, err := database.ReadUserByID(id)
			if err != nil {
				panic(err)
			}

			// create did for this user
			fmt.Printf("now create did for user: %s, id: %d\n", user, id)
			err = CreateDID(chain, sk, user, id)
			if err != nil {
				fmt.Printf("create did failed for user: %s, id: %d, err: %s\n", user, id, err.Error())
			} else {
				fmt.Printf("create did ok for user: %s, id: [%d]\n", user, id)
			}
		}(id)

		time.Sleep(1 * time.Second)
	}

}

func CreateDID(chain string, sk string, user string, id uint64) error {
	// query instance
	instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
	fmt.Println()
	//fmt.Println("eth:", endPoint)
	//fmt.Println("instance address: ", instanceAddr)

	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	// defer cancel()
	ctx := context.Background()

	// get client
	client, err := ethclient.DialContext(ctx, endPoint)
	if err != nil {
		return err
	}

	// get chain id
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("chain id: ", chainID)
		chainID = big.NewInt(666)
		panic(err)
	}
	fmt.Println("chain id: ", chainID)

	// make tx auth
	txAuth, err := com.MakeAuth(chainID, sk)
	if err != nil {
		return err
	}

	callopts := &bind.CallOpts{From: com.AdminAddr}

	// new instanceIns
	instanceIns, err := inst.NewInstance(instanceAddr, client)
	if err != nil {
		return err
	}

	// get proxy addr
	proxyAddr, err := instanceIns.Instances(callopts, com.TypeDidProxy)
	if err != nil {
		return err
	}
	//fmt.Println("proxyAddr:", proxyAddr.Hex())

	// get proxy instance
	proxyIns, err := proxy.NewProxy(proxyAddr, client)
	if err != nil {
		return err
	}

	// base58 decode addr to pubkey
	//pubkey := base58.Decode(userAddr)

	// string to [20]byte

	pubkey := common.HexToAddress(user).Bytes()
	//fmt.Printf("pubkey(userAddr): %x\n", pubkey)
	//fmt.Println("length pubkey: ", len(pubkey))

	// hash pubkey to get did
	did := crypto.Keccak256(pubkey)
	hexDID := hex.EncodeToString(did[:])

	//fmt.Println("did: ", hexDID)
	//fmt.Println("length did: ", len(hexDID))
	//fmt.Println("methodType: ", methodType)

	// call create did
	fmt.Println("call proxy.CreateDIDByAdmin")
	tx, err := proxyIns.CreateDIDByAdmin(txAuth, hexDID, methodType, pubkey)
	if err != nil {
		return err
	}

	// wait for create did tx
	fmt.Println("waiting tx")
	err = com.CheckTx(endPoint, tx.Hash(), "createDid")
	if err != nil {
		return err
	}

	return nil
}
