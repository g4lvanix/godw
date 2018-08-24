package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func getWords(fname string) ([]string, error) {
	var words []string
	file, err := os.Open(fname)

	if err != nil {
		return words, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := strings.Fields(scanner.Text())
		words = append(words, l[1])
	}

	return words, nil
}

func chooseWords(words []string, n int) []string {
	var out []string
	for i := 0; i < n; i++ {
		k, _ := rand.Int(rand.Reader,
			big.NewInt(int64(len(words))))
		out = append(out, words[k.Int64()])
	}
	return out
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Print("\nDiceware password generator\n\n" +
			"Usage: ./godw FILE LENGTH\n\n" +
			"FILE   The diceware word file\n" +
			"LENGTH Number of words in the passphrase\n")
		os.Exit(0)
	}

	filename := args[1]
	nwords, err := strconv.Atoi(args[2])

	if err != nil {
		panic(err)
	} else if nwords < 1 {
		fmt.Print("Password must contain 1 word")
		os.Exit(1)
	}

	words, err := getWords(filename)

	if err != nil {
		panic(err)
	}

	passphrase := strings.Join(chooseWords(words, nwords), " ")

	fmt.Print(passphrase + "\n")
}
