package main

import (
	"fmt"
	"os"

	"github.com/maestro-org/go-sdk/client"
)

func main() {
	maestroClient := client.NewClient("AMXzQsaXXuUPwfV4IWz8q3tMTmO6x40d", "mainnet")
	blockInfo, err := maestroClient.BlockInfo("9005859")
	if err != nil {
		fmt.Printf("Failed to retrieve block info: %s", err)
		os.Exit(1)
	}
	if blockInfo == nil {
		fmt.Printf("Block info empty!")
		os.Exit(1)
	}
	fmt.Println(blockInfo.Data.AbsoluteSlot)
}
