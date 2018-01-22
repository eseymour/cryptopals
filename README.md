# Cryptopals Crypto Challenge Solutions

[![GoDoc](https://godoc.org/github.com/eseymour/cryptopals?status.svg)](https://godoc.org/github.com/eseymour/cryptopals/)
[![Build Status](https://travis-ci.org/eseymour/cryptopals.svg?branch=master)](https://travis-ci.org/eseymour/cryptopals)

This repository contains programs written in [Go](https://golang.org)
to solve the [Crytopals Crypto Challenges](https://cryptopals.com).

## Package Structure

The project is divided into sets containing challenges containing a descriptive
package name for the implementation of the solution. For example, the solution
to challenge 1 is in the package
`github.com/eseymour/cryptopals/set1/challenge01/hexToBase64`. The progress
section also lists the name of the relevant package within the
`github.com/eseymour/cryptopals/set#/challenge##` package.

## Progress

### Set 1

| #   | Challenge                        | Completed        | Package             |
| --- | -------------------------------- | ---------------- | ------------------- |
| 1   | [Convert hex to base64][1]       | January 19, 2018 | `hexToBase64`       |
| 2   | [Fixed XOR][2]                   | January 19, 2018 | `test/fixedXOR`     |
| 3   | [Single-byte XOR cipher][3]      | January 21, 2018 | `test/crackByteXOR` |
| 4   | [Detect single-character XOR][4] |                  |                     |
| 5   | [Implement repeating-key XOR][5] | January 20, 1018 | `test/repeatingXOR` |
| 6   | [Break repeating-key XOR][6]     |                  |                     |
| 7   | [AES in ECB mode][7]             |                  |                     |
| 8   | [Detect AES in ECB mode][8]      |                  |                     |

[1]: https://cryptopals.com/sets/1/challenges/1
[2]: https://cryptopals.com/sets/1/challenges/2
[3]: https://cryptopals.com/sets/1/challenges/3
[4]: https://cryptopals.com/sets/1/challenges/4
[5]: https://cryptopals.com/sets/1/challenges/5
[6]: https://cryptopals.com/sets/1/challenges/6
[7]: https://cryptopals.com/sets/1/challenges/7
[8]: https://cryptopals.com/sets/1/challenges/8
