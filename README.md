bip38-cracker
=============

Incomplete [BIP 0038](https://en.bitcoin.it/wiki/BIP_0038) cracker written in Go. This will eventually attempt to brute force Bitcoin private keys that have been passphrase-protected with BIP 0038. Written to attempt puzzles such as:

* http://www.reddit.com/r/Bitcoin/comments/1q5wu7/this_paper_wallet_contains_01125_btc_and_is_bip/
* https://bitcointalk.org/index.php?topic=128699.0

I'm writing this primarily as an exercise to learn Go, and also to learn more about the inner workings of Bitcoin.

TODO
----

* Complete basic implementation
* Remove gocoin dependency (currently only used for base58 decoding and passpoint calculation)
* Multi-threaded brute forcer

See also
--------

* https://en.bitcoin.it/wiki/BIP_0038
* https://github.com/casascius/Bitcoin-Address-Utility/blob/dcfc3b99a3df1427fc19fcfbe18c1bfedfdad4eb/Model/Bip38KeyPair.cs#L79
* https://github.com/notespace/bip38-cracker/

