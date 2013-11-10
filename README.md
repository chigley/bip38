bip38
=====

Partial [BIP 0038](https://en.bitcoin.it/wiki/BIP_0038) implementation and cracker written in Go. Exported functions are `DecryptWithPassphrase` and `Brute`. The former decrypts a BIP 0038-encrypted private key given the passphrase, returning the private key in a hex string, or the empty string on failure. `Brute` attempts to brute force a Bitcoin private key that has been passphrase-protected with BIP 0038 (currently hardcoded for a passphrase length of 3 characters - see TODO below). Designed to attempt puzzles such as:

* http://www.reddit.com/r/Bitcoin/comments/1q5wu7/this_paper_wallet_contains_01125_btc_and_is_bip/
* https://bitcointalk.org/index.php?topic=128699.0

I'm writing this primarily as an exercise to learn Go, and also to learn more about the inner workings of Bitcoin.

TODO
----

* Proper support for brute forcing over a given character set, rather than a hardcoded value
* Support for variable-length passphrases, or to search a given length range
* Catch panic case when all goroutines return without finding the passphrase (i.e. brute force failure)
* Implement decryption when EC multiply mode not used (and write tests)
* Add proper doc comments to code
* Optimise searchRange when not starting at 0
* Remove gocoin dependency (currently only used for base58 decoding, and passpoint & address calculation)

Note about cracking
-------------------

The first argument to `Brute` is an `int` number of goroutines to call while attempting the brute force search. For best results, set this to the number of cores you have available on your machine. Also make sure you set the `GOMAXPROCS` environment variable to this value before compiling! Otherwise your program will only run on a single core (by default), regardless of what you set `routines` to.

On my quad-core Intel i5-2500K, I get speeds of about 10 keys/second when brute forcing a 3 character passphrase over the printable ASCII characters.

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
