package main

import (
	"github.com/kirinlabs/HttpRequest"
	"strings"
)

const baseUrl = "https://tokenwallet.one/create/"

func generate_wallet(symbol string) {
	url := baseUrl + strings.ToLower(symbol)
	resp, err := HttpRequest.NewRequest().Get(url)
	defer resp.Close()
	if err != nil {
		println(err.Error())
	}
	body, err := resp.Body()
	if err != nil {
		println(err.Error())
	}
	println(string(body))
}

func main() {
	//supported symbols: btc eth etc heco bsc okt trx usdt_erc20 usdt_trc20 usdt_omni
	//					eos bch qtum fil bsv vet matic ltc dash doge atom xmr waves
	//					xrp bnb neo ont xlm pote club
	generate_wallet("btc")
	generate_wallet("eth")
	generate_wallet("etc")
	generate_wallet("heco")
	generate_wallet("bsc")
	generate_wallet("trx")
	generate_wallet("usdt_erc20")
	generate_wallet("usdt_trc20")
	generate_wallet("usdt_omni")
	generate_wallet("eos")
	generate_wallet("bch")
	generate_wallet("fil")
	generate_wallet("bsv")
	generate_wallet("vet")
	generate_wallet("matic")
	generate_wallet("ltc")
	generate_wallet("dash")
	generate_wallet("doge")
	generate_wallet("atom")
	generate_wallet("xmr")
	generate_wallet("waves")
	generate_wallet("xrp")
	generate_wallet("bnb")
	generate_wallet("neo")
	generate_wallet("ont")
	generate_wallet("xlm")
	generate_wallet("pote")
	generate_wallet("club")
}
