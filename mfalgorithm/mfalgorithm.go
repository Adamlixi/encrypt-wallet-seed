package mfalgorithm

import (
	"crypto/sha256"
	"strconv"
)

func SetMnemonic(mnemonicsRandom []string, mnemonics []string, wordListRandom []string, list39 []string) []string {
	if len(wordListRandom) < 65535 {
		return nil
	}
	if len(mnemonics) < 8 || len(mnemonics) > 20 {
		return nil
	}
	seed := GetSeed(mnemonicsRandom, list39)
	mnemonicsLen := len(mnemonics)
	if len(seed)%mnemonicsLen != 0 {
		hashRes := sha256.Sum256(seed)

		n := len(seed) % mnemonicsLen
		m := mnemonicsLen - n
		seedAddLen := n
		if n > m {
			seedAddLen = m
		}
		seed = append(seed, hashRes[:seedAddLen]...)
	}
	segmentLen := len(seed) / len(mnemonics)
	start := 0
	segments := make([][]byte, 0)
	for i := segmentLen; i < len(seed); i += segmentLen {
		segments = append(segments, seed[start:i])
		start += segmentLen
	}
	for index, mnemonic := range mnemonics {
		for _, word := range wordListRandom {
			if mnemonic == word {
				in, err := strconv.ParseInt(string(segments[index]), 2, 32)
				if err != nil {
					panic(err)
				}
				wordListRandom[index], wordListRandom[in] = wordListRandom[in], wordListRandom[index]
			}
		}
	}
	return wordListRandom
}

func GetSeed(mnemonicsRandom []string, list []string) []byte {
	var seed []byte
	for _, word := range mnemonicsRandom {
		for i, v := range list {
			if v == word {
				s2 := strconv.FormatInt(int64(i), 2)
				seed = append(seed, []byte(s2)...)
				break
			}
		}
	}
	return seed
}
