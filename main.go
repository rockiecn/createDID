package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/rockiecn/createDID/database"
)

var (
	AdminSK   = "6ec7e0cdda802a466401c912b0dac5ff6116e4372746a455bb43c61294e6f01f"
	AdminAddr = "0xe0DA2108c52C799F27B02D5AF049374B294f291e"
)

func main() {
	//        inputeth := flag.String("eth", "dev", "eth api Address;") //dev test or product
	//        psk := flag.String("sk", "", "signature for sending transaction")
	pstart := flag.Uint64("start", 1, "start id")
	pbalance := flag.String("balance", "0", "minimum balance")

	flag.Parse()

	start := *pstart
	bal := *pbalance

	// open db
	database.Open()
	defer database.DB.Close()

	var id uint64
	for id = start; id <= 467683; id++ {
		go func(id uint64) {
			// read an user from db
			fmt.Println("now reading user: ", id)
			user, balance, err := database.ReadUserBalByID(id)
			if err != nil {
				panic(err)
			}
			fmt.Printf("user: %s, balance: %s, startBal: %s\n", user, balance, bal)

			actBal, _ := new(big.Int).SetString(balance, 10)
			startBal, _ := new(big.Int).SetString(bal, 10)
			// check balance
			if actBal.Cmp(startBal) < 0 {
				fmt.Println("balance too low, skip")
			} else {
				// create did for this user
				fmt.Printf("now create did for user: %s, id: %d\n", user, id)
				err = CreateDID(user)
				if err != nil {
					fmt.Printf("create did failed for user: %s, id: %d, err: %s\n", user, id, err.Error())
				} else {
					fmt.Printf("create did ok for user: %s, id: [%d]\n", user, id)
				}
			}
		}(id)

		time.Sleep(1 * time.Second)
	}

}

func CreateDID(user string) error {
	url := "https://didapi.memolabs.org/did/create"

	// 使用user地址创建请求数据
	data := map[string]string{"address": user, "sig": "sig"} // 这里假设你有签名，可以替换 "sig"
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return err
	}

	// 创建POST请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}

	// 打印响应内容
	fmt.Println("Response:", string(body))
	return nil
}
