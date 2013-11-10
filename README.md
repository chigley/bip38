bip38
=====

Incomplete [BIP 0038](https://en.bitcoin.it/wiki/BIP_0038) implementation and cracker written in Go. Attempts to brute force Bitcoin private keys that have been passphrase-protected with BIP 0038. Currently hardcoded for a passphrase length of 3 characters. Written to attempt puzzles such as:

* http://www.reddit.com/r/Bitcoin/comments/1q5wu7/this_paper_wallet_contains_01125_btc_and_is_bip/
* https://bitcointalk.org/index.php?topic=128699.0

I'm writing this primarily as an exercise to learn Go, and also to learn more about the inner workings of Bitcoin.

TODO
----

* Remove gocoin dependency (currently only used for base58 decoding and passpoint calculation)
* Write tests for BIP 0038 implementation at the very least
* Catch panic case when all goroutines return without finding the passphrase
* Optimise searchRange when not starting at 0, and add in support for variable-length passphrases
* Implement decryption when EC multiply mode not used
* Add a note about GOMAXPROCS and the routines parameter

Example
-------

    package main
    
    import (
    	"fmt"
    	"github.com/chigley/bip38"
    )
    
    func main() {
    	privKey := bip38.DecryptWithPassphrase("6PfQu77ygVyJLZjfvMLyhLMQbYnu5uguoJJ4kMCLqWwPEdfpwANVS76gTX", "TestingOneTwoThree")
    	fmt.Printf("%s", privKey)
    	// a43a940577f4e97f5c4d39eb14ff083a98187c64ea7c99ef7ce460833959a519
    
    	result := bip38.Brute(4, "6PfYWMgNPboK3PQ1D5mNqaKxRmD9j7Wooi2wLmjVM2Ze776Qx3tyMR26pq")
    	fmt.Printf("%s", result)
    	// 6d0507ca41e06bd1a9297e82d1a1f6ee9a824b14bb5170873ef7ca8296701f7f (  !)
    }

See also
--------

* https://en.bitcoin.it/wiki/BIP_0038
* https://github.com/casascius/Bitcoin-Address-Utility/blob/dcfc3b99a3df1427fc19fcfbe18c1bfedfdad4eb/Model/Bip38KeyPair.cs#L79
* https://github.com/notespace/bip38-cracker/
* https://gist.github.com/laanwj/7372573
