package mfalgorithm

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestMF(t *testing.T) {
	// open file
	f39, err := os.Open("wordlist39")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f39.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f39)
	scanner.Split(bufio.ScanWords)
	var list39 []string
	for scanner.Scan() {
		list39 = append(list39, scanner.Text())
	}

	f600, err := os.Open("wordlist600k")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f600.Close()

	// read the file word by word using scanner
	scanner2 := bufio.NewScanner(f600)
	scanner2.Split(bufio.ScanWords)
	var wordListRandom []string
	for scanner2.Scan() {
		wordListRandom = append(wordListRandom, scanner.Text())
	}
	seed1 := GetSeed(list39, randomMnemonic)
	randomMnemonic := []string{"enjoy", "lunar", "follow", "dismiss", "gentle", "old", "manage", "lyrics", "feature", "combine", "skate", "reunion"}
	mnemonicSet := []string{"want", "to", "eat", "hamburgers", "which", "i", "think", "are", "delicious"}
	mappingFile := SetMnemonic(randomMnemonic, mnemonicSet, wordListRandom, list39)
	seed2 := GetSeed(mappingFile, mnemonicSet)
	fmt.Println(seed1)
	fmt.Println(seed2)
}
