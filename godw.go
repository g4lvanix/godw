package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var filename = flag.String("f", "eff_large_wordlist.txt", "name of the diceware wordlist file")
var pin = flag.Bool("p", false, "generate a numeric pin")
var length = flag.Int("l", 0, "length of the generated passphrase or pin")

func getWords(fname string) ([]string, error) {
	var words []string

	file, err := os.Open(fname)
	if err != nil {
		return words, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := strings.Fields(scanner.Text())
		words = append(words, l[1])
	}
	return words, nil
}

func getPasshphrase(words []string, n int) string {
	var out []string
	for i := 0; i < n; i++ {
		k, _ := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
		out = append(out, words[k.Int64()])
	}
	return strings.Join(out, " ")
}

func getPin(n int) string {
	var out []string
	for i := 0; i < n; i++ {
		k, _ := rand.Int(rand.Reader, big.NewInt(10))
		out = append(out, strconv.Itoa(int(k.Int64())))
	}
	return strings.Join(out, "")
}

func main() {
	flag.Parse()

	if *length == 0 {
		fmt.Print("The passphrase/pin length must be greater than 0\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if *pin {
		fmt.Println(getPin(*length))
	} else {
		words, err := getWords(*filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(getPasshphrase(words, *length))
	}
}
