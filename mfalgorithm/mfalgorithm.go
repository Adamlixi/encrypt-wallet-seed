package mfalgorithm

import (
	"crypto/sha256"
	"fmt"
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
		var hashByte []byte
		for _, v := range hashRes {
			s2 := fmt.Sprintf("%08b", int64(v))
			hashByte = append(hashByte, []byte(s2)...)
		}
		n := len(seed) % mnemonicsLen
		seedAddLen := mnemonicsLen - n
		seed = append(seed, hashByte[:seedAddLen]...)
	}
	segmentLen := len(seed) / len(mnemonics)
	start := 0
	segments := make([][]byte, 0)
	for i := segmentLen; i <= len(seed); i += segmentLen {
		segments = append(segments, seed[start:i])
		start += segmentLen
	}
	for index, mnemonic := range mnemonics {
		for pos, word := range wordListRandom {
			if mnemonic == word {
				ss := string(segments[index])
				in, err := strconv.ParseInt(ss, 2, 64)
				if err != nil {
					panic(err)
				}
				wordListRandom[pos], wordListRandom[in] = wordListRandom[in], wordListRandom[pos]
				break
			}
		}
	}
	return wordListRandom
}

func GetSeed(mnemonics []string, list []string) []byte {
	var seed []byte
	mnemonicsLen := len(mnemonics)
	seedAddLen := 0
	if 128%mnemonicsLen != 0 {
		n := 128 % mnemonicsLen
		seedAddLen = mnemonicsLen - n
	}
	lenPerSeg := (128 + seedAddLen) / len(mnemonics)
	formatStr := "%0" + strconv.FormatInt(int64(lenPerSeg), 10) + "b"
	for _, word := range mnemonics {
		for i, v := range list {
			if v == word {
				s2 := fmt.Sprintf(formatStr, int64(i))
				seed = append(seed, []byte(s2)...)
				break
			}
		}
	}
	seedLen := len(seed) - seedAddLen
	return seed[:seedLen]
}
