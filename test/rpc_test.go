package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/rjman-self/go-polkadot-rpc-client/client"
	"testing"
)

func Test_GetBlockByNumber(t *testing.T) {
	c, err := client.New("ws://127.0.0.1:9944")
	if err != nil {
		t.Fatal(err)
	}

	c.SetPrefix(ss58.PolkadotPrefix)
	//expand.SetSerDeOptions(false)
	resp, err := c.GetBlockByNumber(58492)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(resp)
	fmt.Println(string(d))
}

func Test_GetAccountInfo(t *testing.T) {
	c, err := client.New("ws://127.0.0.1:9944")
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.PolkadotPrefix)
	ai, err := c.GetAccountInfo("15oF4uVJwmo4TdGW7VfQxNLavjCXviqxT9S1MgbjMNHr6Sp5")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(ai)
	fmt.Println(string(d))
	fmt.Println(ai.Data.Free.String())
}
