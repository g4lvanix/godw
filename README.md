# godw - A diceware passhphrase generator written in Go

This command line tool can generate secure but memorable 
passhphrases based on the [diceware technique](http://world.std.com/~reinhold/diceware.html).
The Electronic Frontier Foundation (EFF) has a great article on [how this technique works](https://www.eff.org/dice).

Besides this tool you need a word list, either the [original diceware list](http://world.std.com/~reinhold/diceware.html) created by Arnold G. Reinhold 
or one of the [lists published by the EFF](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases).

`godw` uses the cryptographically secure pseudorandom number generator provided by the 
Go standard library in [`crypto/rand`](https://golang.org/pkg/crypto/rand/) to pick 
random words from the word list. 
The words are chosen according to a unifom distribution (every word is equally likely to be chosen).

## Installation

After [installing go](https://golang.org/doc/install), simply run

```
$ go get https://github.com/g4lvanix/godw.git
```

Additionally, please obtain one of the wordlists from the sources mentioned above.

## Usage 

```
$ ./godw FILE LENGTH
```

Where `FILE` is one of the wordlists mentioned above and `LENGTH` is the 
number of words in the passhprase. 

To generate a 5 word long passhphrase with words chosen from EFFs large wordlist:
```
$ ./godw eff_large_wordlist.txt 5
```

## Notes 

[Relevant xkcd](https://www.xkcd.com/936/)
