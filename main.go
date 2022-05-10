package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

func main() {
	entropy, err := getPassword()
	if err != nil {
		panic(err)
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

	fmt.Println(mnemonic)
}

func getPassword() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter entropy: ")
	entropy, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(entropy), nil
}
