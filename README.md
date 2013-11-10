bip38-cracker
=============

Incomplete [BIP 0038](https://en.bitcoin.it/wiki/BIP_0038) cracker written in Go. Attempts to brute force Bitcoin private keys that have been passphrase-protected with BIP 0038. Currently hardcoded for a passphrase length of 3 characters. Written to attempt puzzles such as:

* http://www.reddit.com/r/Bitcoin/comments/1q5wu7/this_paper_wallet_contains_01125_btc_and_is_bip/
* https://bitcointalk.org/index.php?topic=128699.0

I'm writing this primarily as an exercise to learn Go, and also to learn more about the inner workings of Bitcoin.

TODO
----

* Remove gocoin dependency (currently only used for base58 decoding and passpoint calculation)
* Move to package of its own rather than being in main package
* Write tests for BIP 0038 implementation at the very least

See also
--------

* https://en.bitcoin.it/wiki/BIP_0038
* https://github.com/casascius/Bitcoin-Address-Utility/blob/dcfc3b99a3df1427fc19fcfbe18c1bfedfdad4eb/Model/Bip38KeyPair.cs#L79
* https://github.com/notespace/bip38-cracker/
* https://gist.github.com/laanwj/7372573

