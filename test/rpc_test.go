package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/rjman-self/go-polkadot-rpc-client/client"
	"testing"
)

const url = "ws://127.0.0.1:8087"

func Test_GetBlockByNumber(t *testing.T) {
	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPrefix(ss58.ChainXPrefix)
	//expand.SetSerDeOptions(false)
	resp, err := c.GetBlockByNumber(872)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(resp)
	fmt.Println(string(d))
}

func Test_GetAccountInfo(t *testing.T) {
	c, err := client.New(url)
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
