package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

var start = flag.Int64("s", 1, "ethereum hd index start ")
var end = flag.Int64("e", 1, "ethereum hd index end ")

const ETHEREUM_HD_PATH = "m/44'/60'/0'/0/"

func main() {
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter entropy: ")
	entropy, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(entropy) < 16 {
		fmt.Println("seed length is too short")
		return
	}

	mnemonic, err := bip39.NewMnemonic([]byte(entropy)[:16])
	if err != nil {
		fmt.Printf("NewMnemonic error %s\n", err)
		return
	}

	fmt.Printf("entropy: %s \nmnemonic: %s\n", []byte(entropy)[:16], mnemonic)

	hdWallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		fmt.Printf("NewFromMnemonic err:%s", err.Error())
		return
	}

	for i := *start; i <= *end; i++ {
		path := ETHEREUM_HD_PATH + fmt.Sprintf("%d", i)
		drPath, err := hdwallet.ParseDerivationPath(path)
		if err != nil {
			panic(err)
		}

		account, err := hdWallet.Derive(drPath, false)
		if err != nil {
			panic(err)
		}
		priv, err := hdWallet.PrivateKeyHex(account)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s: %s 0x%s\n", path, strings.ToLower(account.Address.String()), priv)
	}
}
