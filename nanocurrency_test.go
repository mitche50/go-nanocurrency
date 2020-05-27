package nanocurrency

import (
	"fmt"
	"sort"
	"testing"
)

//Negative Test Cases
func TestInvalidAccount(t *testing.T) {

	rpc := NanoRPC{Host: "http://localhost", Port: "7076"}
	testAccount := "thisisaninvalidaccount"
	_, err := AccountBalance(rpc, testAccount)
	if err != nil {
		want := "Bad account number"
		got := err.Error()

		if want != got {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
}

func TestInvalidCommand(t *testing.T) {
	rpc := NanoRPC{Host: "http://localhost", Port: "7076"}
	data := map[string]string{"action": "invalid_action"}
	_, err := NodePost(rpc.Host, rpc.Port, &data)
	if err != nil {
		want := "Unknown command"
		got := err.Error()

		if want != got {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
}

//Block Test Cases
func TestBlockCount(t *testing.T) {

	rpc := NanoRPC{Host: "http://localhost", Port: "7076"}
	blockMap, err := BlockCount(rpc)
	if err != nil {
		t.Errorf(err.Error())
	}

	var gotKeys []string
	wantKeys := make([]string, len(blockMap))
	wantKeys[0] = "count"
	wantKeys[1] = "unchecked"

	for k := range blockMap {
		fmt.Println("gotKeys: ", gotKeys)
		fmt.Println("k: ", k)
		gotKeys = append(gotKeys, k)
	}

	sort.Sort(sort.StringSlice(gotKeys))

	for i, v := range gotKeys {
		if v != wantKeys[i] {
			t.Errorf("got '%s' want '%s'", gotKeys, wantKeys)
		}
	}
}

//Account Test Cases
func TestAccountBalance(t *testing.T) {

	rpc := NanoRPC{Host: "http://localhost", Port: "7076"}
	testAccount := "xrb_376w1ygoeqh3c1qrighfmea7zxrzczke6ewwjqu5f54ptienkcomh14xm3fh"
	accountMap, err := AccountBalance(rpc, testAccount)
	if err != nil {
		t.Errorf(err.Error())
	}

	var gotKeys []string
	wantKeys := make([]string, len(accountMap))
	wantKeys[0] = "balance"
	wantKeys[1] = "pending"

	for k := range accountMap {
		gotKeys = append(gotKeys, k)
	}

	sort.Sort(sort.StringSlice(gotKeys))

	for i, v := range gotKeys {
		if v != wantKeys[i] {
			t.Errorf("got '%s' want '%s'", gotKeys, wantKeys)
		}
	}
}

func TestAccountBlocks(t *testing.T) {
	rpc := NanoRPC{Host: "http://localhost", Port: "7076"}
	testAccount := "xrb_376w1ygoeqh3c1qrighfmea7zxrzczke6ewwjqu5f54ptienkcomh14xm3fh"
	accountBlocks, err := AccountBlocks(rpc, testAccount)
	if err != nil {
		t.Errorf(err.Error())
	}

	var gotKeys []string
	wantKeys := make([]string, len(accountBlocks))

	if len(accountBlocks) > 0 {
		wantKeys[0] = "block_count"
	} else {
		t.Errorf("Account Blocks return is nil.")
	}

	for k := range accountBlocks {
		gotKeys = append(gotKeys, k)
	}

	sort.Sort(sort.StringSlice(gotKeys))

	if len(gotKeys) == len(wantKeys) {

		for i, v := range gotKeys {
			if v != wantKeys[i] {
				t.Errorf("got '%s' want '%s'", gotKeys, wantKeys)
			}
		}
	} else {
		t.Errorf("got '%s' want '%s'", gotKeys, wantKeys)
	}

}

func TestAccountInformation(t *testing.T) {
	rpc := NanoRPC{Host: "http://localhost", Port: "7076"}
	testAccount := "xrb_376w1ygoeqh3c1qrighfmea7zxrzczke6ewwjqu5f54ptienkcomh14xm3fh"

	optionalInfo := map[string]string{"weight": "true", "representative": "true", "pending": "true"}

	accountInfo, err := AccountInformation(rpc, testAccount, optionalInfo)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(accountInfo) > 0 {
		wantKeys := []string{"account_version", "balance", "block_count", "frontier", "modified_timestamp", "open_block", "pending", "representative", "representative_block", "weight"}
		var gotKeys []string

		for k := range accountInfo {
			gotKeys = append(gotKeys, k)
		}

		sort.Sort(sort.StringSlice(gotKeys))

		if len(gotKeys) == len(wantKeys) {
			for i, v := range gotKeys {
				if v != wantKeys[i] {
					t.Errorf("got '%s' want '%s'", gotKeys, wantKeys)
				}
			}
		} else {
			t.Errorf("got '%s' want '%s'", gotKeys, wantKeys)
		}
	} else {
		t.Errorf("Account Information return is nil.")
	}

}

func TestAccountCreate(t *testing.T) {
	rpc := NanoRPC{Host: "http://localhost", Port: "7076"}
	testWallet := "71BA065A2B1C394F25786862EEAB88CE38B351909A744A2470F7591478C9E386"

	optionalCreate := map[string]string{"work": "false", "index": "0"}

	accountCreate, err := AccountCreate(rpc, testWallet, optionalCreate)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		if len(accountCreate) > 0 {
			wantKeys := []string{"account"}
			var gotKeys []string

			for k := range accountCreate {
				gotKeys = append(gotKeys, k)
			}

			sort.Sort(sort.StringSlice(gotKeys))

			if len(gotKeys) == len(wantKeys) {
				for i, v := range gotKeys {
					if v != wantKeys[i] {
						t.Errorf("got '%s' want '%s'", gotKeys, wantKeys)
					}
				}
			} else {
				t.Errorf("got '%s' want '%s'", gotKeys, wantKeys)
			}
		} else {
			t.Errorf("Account Create return is nil.")
		}
	}
}

func TestAccountGet(t *testing.T) {
	rpc := NanoRPC{Host: "http://localhost", Port: "7076"}
	testKey := "949C079D565DE1502F8839ED9B105FF71F57E4C2339C8DF6368C56D419492AB3"

	accountGet, err := AccountGet(rpc, testKey)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		if len(accountGet) > 0 {
			wantKeys := []string{"account"}
			var gotKeys []string

			for k := range accountGet {
				gotKeys = append(gotKeys, k)
			}

			sort.Sort(sort.StringSlice(gotKeys))

			if len(gotKeys) == len(wantKeys) {
				for i, v := range gotKeys {
					if v != wantKeys[i] {
						t.Errorf("got '%s' want '%s'", gotKeys, wantKeys)
					}
				}
			} else {
				t.Errorf("got '%s' want '%s'", gotKeys, wantKeys)
			}
		} else {
			t.Errorf("Account Get return is nil.")
		}
	}
}
