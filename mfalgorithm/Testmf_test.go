package mfalgorithm

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMF(t *testing.T) {
	// read BIP39 word list(2048 words)
	f39, err := os.Open("wordlist39")
	if err != nil {
		panic(err)
	}
	defer f39.Close()

	scanner := bufio.NewScanner(f39)
	scanner.Split(bufio.ScanWords)
	var list39 []string
	for scanner.Scan() {
		list39 = append(list39, scanner.Text())
	}
	// read MF Algorithm word list(larger than 65535)
	f600, err := os.Open("wordlist600k")
	if err != nil {
		log.Fatal(err)
	}
	defer f600.Close()

	scanner2 := bufio.NewScanner(f600)
	scanner2.Split(bufio.ScanWords)
	var wordListRandom []string
	for scanner2.Scan() {
		wordListRandom = append(wordListRandom, scanner2.Text())
	}
	// shuffle the word list
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(wordListRandom), func(i, j int) {
		wordListRandom[i], wordListRandom[j] = wordListRandom[j], wordListRandom[i]
	})
	// import seed phrase from Metamask
	randomMnemonic := []string{"world", "table", "follow", "urban", "uphold", "usage", "manage", "useless", "feature", "whale", "skate", "yard"}
	// set customized seed phrase
	mnemonicSet := []string{"want", "to", "eat", "hamburgers", "which", "i", "think", "are", "delicious"}
	seed1 := GetSeed(randomMnemonic, list39)
	// get mapping file, you can store it in your cloud
	mappingFile := SetMnemonic(randomMnemonic, mnemonicSet, wordListRandom, list39)
	// use mapping file and customized seed phrase to get back seed
	seed2 := GetSeed(mnemonicSet, mappingFile)
	s1 := string(seed1)
	s2 := string(seed2)
	if s1 != s2 {
		t.Fatal("fail")
	}
}
