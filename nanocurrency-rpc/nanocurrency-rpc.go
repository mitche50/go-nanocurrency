package nanocurrencyrpc

import (
	"fmt"
	nano "nanocurrency"
)

func nanocurrencyrpc() {
	rpc := nano.NanoRPC{Host: "http://[::1]", Port: "55000"}
	testWallet := "764D0C665E14B1675F4E6E394DE4FE23791ADA303199660DF2F9E0CA57A85319"
	testAccount := "nano_3q3ku818j8764mjatogd6t8pfoymixhn5rg94n31ypiqwfzm6gfycu7cknon"

	blockCount, blockError := nano.BlockCount(rpc)
	if blockError != nil {
		fmt.Println("Error: ", blockError)
	} else {
		fmt.Println("\nBlock Count Return\n")
		fmt.Println("checked blocks: ", blockCount["count"])
		fmt.Println("unchecked blocks: ", blockCount["unchecked"])
	}

	accountBalance, accountError := nano.AccountBalance(rpc, testAccount)
	if accountError != nil {
		fmt.Println("Account Balance Error: ", accountError)
	} else {
		fmt.Println("\nAccount Balance Return\n")
		fmt.Println("pending: ", accountBalance["pending"])
		fmt.Println("balance: ", accountBalance["balance"])
	}

	optionalInfo := map[string]string{"weight": "true", "representative": "true", "pending": "true"}
	accountInfo, infoError := nano.AccountInformation(rpc, testAccount, optionalInfo)
	if infoError != nil {
		fmt.Println("Account Balance Error: ", infoError)
	} else {
		fmt.Println("\nAccount Information Return\n")
		for k, v := range accountInfo {
			fmt.Println(k, v)
		}
	}

	optionalCreate := map[string]string{"work": "false", "index": "0"}

	accountCreate, createError := nano.AccountCreate(rpc, testWallet, optionalCreate)
	if createError != nil {
		fmt.Println("Account Create error: ", createError)
	} else {
		fmt.Println("\nAccount Create Return\n")
		for k, v := range accountCreate {
			fmt.Println(k, v)
		}
	}

	//optionalHistory := map[string]string{"raw": "true"}

	accountHistory, historyError := nano.AccountHistory(rpc, testAccount, "3", nil)
	if historyError != nil {
		fmt.Println("Account History error: ", historyError)
	} else {
		fmt.Println("\nAccount History Return\n")
		for k, v := range accountHistory {
			switch valueType := v.(type) {
			case []interface{}:
				for i, u := range valueType {
					fmt.Printf("\nBlock %v\n", i)
					for k2, v2 := range u.(map[string]interface{}) {
						fmt.Println(k2, v2)
					}
				}
			case interface{}:
				fmt.Println(k, v)
			default:
				fmt.Printf("unexpected type %T", valueType)
			}
		}
	}

	for i := 0; i <= 1000; i++ {
		blockCount, blockError := nano.BlockCount(rpc)
		if blockError != nil {
			fmt.Println("Error: ", blockError)
		} else {
			fmt.Printf("\n%v: Block Count Return\n\n", i)
			fmt.Println("checked blocks: ", blockCount["count"])
			fmt.Println("unchecked blocks: ", blockCount["unchecked"])
		}
	}
}
