# zorkmid

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![Tests](https://github.com/ZorkNetwork/zorkmid/actions/workflows/tests.yaml/badge.svg)](https://github.com/ZorkNetwork/zorkmid/actions/workflows/tests.yaml)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/ZorkNetwork/zorkmid)

Zorkmid is the reference node implementaiton for Zork crypto currency written in Go.

Zorkmid also known as:

- ZorkCoin™
- Zorkmid
- Zork
- Zorks™
- MilliZorks
- MicroZorks
- KiloZorks
- MegaZorks
- ZORK™

## What is Zorkmid? aka Zorks?

Zorks are a proof-of-work cryptocurrency forked from Kaspa. It is based on [the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a generalization of Nakamoto consensus. The ZorkCoin currency is named after the epic game series
of Zork. ZorkCoins are generically referenced as Zorks and the ticker symbol is
ZORK.

We are happy to have forked our reference node from Kaspa and maintain it as a branch of the original source. This allows for easier maintenance of the future
of Zorkmid since we can continue to update from our orinator, as well as provide bugfixes back to our origin. Though this does create some confusion when looking at some of the source since the name 'Kaspa' is not replaced through the entire source tree, instead the source is only modified where required to do so for the changes implemented in Zorks.

Note our transparency of any and all changes to the Kaspa code.

This Golang code is simply the start of the Zork Network.

Any & All Contributors are referenced as members of the Zork Underground.

## Installation

### Download & Use a Release

- As there are no releases yet there's not really a write-up to do it yet, but
this placeholder exists to help you envision that there will be releases in the
future!

### Build from Source

## Requirements

Tested with Go 1.23.4. Later should work, earlier might work.

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
go version
```

- Run the following commands to obtain and install the node including all dependencies:

```bash
git clone https://github.com/ZorkNetwork/zorkmid
cd zorkmid
go install . ./cmd/...
```

- The node and related utilities should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.

## Getting Started

The node has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
zorkmid
```

## Discord

Join the [Zork Underground](https://discord.gg/gmpSzpqCDh).

## Issue Tracker

Since zorkmid was forked from the kaspanet/kaspad repository, many of the issues
as related to that repository may also exist in this one.
The [integrated github issue tracker for kaspad](https://github.com/kaspanet/kaspad/issues)
can be used for issues that exists in both.

At this time zorkmid doesn't yet have a tracking system configured, so at this
time use discord, or email [Barbazzo Fernap](196495312+Barbazzo-Fernap@users.noreply.github.com)

## Documentation

Zork Network documentation is still in development. Join now to help guide the future of the netwrok!

The [Kaspa documentation](https://github.com/kaspanet/docs) can be used for portions of this project.

## License

Zorkmid is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
since it is forked from Kaspad.
