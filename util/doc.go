/*
Package util provides kaspa-specific convenience functions and types.

# Block Overview

A Block defines a kaspa block that provides easier and more efficient
manipulation of raw blocks. It also memoizes hashes for the
block and its transactions on their first access so subsequent accesses don't
have to repeat the relatively expensive hashing operations.

# Tx Overview

A Tx defines a kaspa transaction that provides more efficient manipulation of
raw transactions. It memoizes the hash for the transaction on its
first access so subsequent accesses don't have to repeat the relatively
expensive hashing operations.

# Address Overview

The Address interface provides an abstraction for a kaspa address. While the
most common type is a pay-to-pubkey, kaspa already supports others and
may well support more in the future. This package currently provides
implementations for the pay-to-pubkey, and pay-to-script-hash address
types.

To decode/encode an address:

	addrString := "zork:qzvkdlhlc5s6cdu6mayfm3kyh4z2mwag8ragvukcps4su34e3rgrs7e535nhj"
	defaultPrefix := util.Bech32PrefixKaspa
	addr, err := util.DecodeAddress(addrString, defaultPrefix)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(addr.EncodeAddress())
*/
package util
