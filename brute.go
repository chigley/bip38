package bip38

import (
	"log"
	"math"
)

var totalTried = 0

func searchRange(start int, finish int, encryptedKey string, charset string, c chan string) {
	i := 0
	for _, rune1 := range charset {
		for _, rune2 := range charset {
			for _, rune3 := range charset {
				if start <= i {
					guess := string(rune1) + string(rune2) + string(rune3)

					privKey := DecryptWithPassphrase(encryptedKey, guess)
					if privKey != "" {
						c <- privKey + " (" + guess + ")"
					}

					if totalTried%10 == 0 {
						log.Printf("%d passphrases tried", totalTried)
					}

					totalTried++
				} else if i == finish {
					return
				}

				i++
			}
		}
	}
}

func Brute(routines int, encryptedKey string) string {
	length := 3 // Length of passphrase, hardcoded for now...

	if encryptedKey == "" {
		log.Fatal("encryptedKey required")
	}

	if routines < 1 {
		log.Fatal("routines must be >= 1")
	}

	if length < 1 {
		log.Fatal("length must be >= 1")
	}

	// Extended ASCII
	//charset := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~€‚ƒ„…†‡ˆ‰Š‹ŒŽ‘’“”•–—˜™š›œžŸ ¡¢£¤¥¦§¨©ª«¬­®¯°±²³´µ¶·¸¹º»¼½¾¿ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿ"

	// Printable ASCII
	charset := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	spaceSize := int(math.Pow(float64(len(charset)), float64(length)))
	blockSize := spaceSize / routines

	c := make(chan string)

	for i := 0; i < routines; i++ {
		var finish int
		if i == routines-1 {
			// Last block needs to go right to the end of the search space
			finish = spaceSize
		} else {
			finish = i*blockSize + blockSize
		}
		go searchRange(i*blockSize, finish, encryptedKey, charset, c)
	}

	return <-c
}
