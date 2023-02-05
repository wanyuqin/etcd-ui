package etcdv3

import (
	"fmt"
	"log"
	"testing"
)

func TestEtcdCli_ListKeys(t *testing.T) {
	err := NewClientV3("127.0.0.1:31368")
	if err != nil {
		log.Fatalln(err)
	}

	keys, err := Cli.ListKeys()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v\n", keys)
}
