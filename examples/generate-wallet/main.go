package main

import (
	"fmt"
	"github.com/glossd/btc/netchain"
	"github.com/glossd/btc/wallet"
)

func main() {
	compress := false
	format := wallet.GetFormat(compress)
	net := netchain.TestNet
	wif, address, privateKey := wallet.New(net, compress)
	fmt.Printf(
		"New wallet generated.\n\n"+
			"Net: %s\n"+
			"WIF: %s\n"+
			"Bitcoin address: %s\n"+
			"Private Key: %s\n\n"+
			"WIF and Bitcoin address are in %s format!\n",
		net, wif, address, privateKey, format)
}
