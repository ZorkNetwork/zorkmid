// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package util_test

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/crypto/blake2b"

	"github.com/ZorkNetwork/zorkmid/util"
)

func TestAddresses(t *testing.T) {
	tests := []struct {
		name           string
		addr           string
		encoded        string
		valid          bool
		result         util.Address
		f              func() (util.Address, error)
		passedPrefix   util.Bech32Prefix
		expectedPrefix util.Bech32Prefix
	}{
		// Positive P2PK tests.
		{
			name:    "mainnet p2pk",
			addr:    "zork:qrlcu5hc3f4tcfpyz5tq27h2gyr2m4at0wuate4q30jmyysfgswhghuqvu2na",
			encoded: "zork:qrlcu5hc3f4tcfpyz5tq27h2gyr2m4at0wuate4q30jmyysfgswhghuqvu2na",
			valid:   true,
			result: util.TstAddressPubKey(
				util.Bech32PrefixKaspa,
				[util.PublicKeySize]byte{
					//6d8e52f88a6abc24241516057aea4106add7ab7bb9d5e6a08be5b21209441d74
					0xFF, 0x8e, 0x52, 0xf8, 0x8a, 0x6a, 0xbc, 0x24,
					0x24, 0x15, 0x16, 0x05, 0x7a, 0xea, 0x41, 0x06,
					0xad, 0xd7, 0xab, 0x7b, 0xb9, 0xd5, 0xe6, 0xa0,
					0x8b, 0xe5, 0xb2, 0x12, 0x09, 0x44, 0x1d, 0x74,
				}),
			f: func() (util.Address, error) {
				publicKey := []byte{
					0xFF, 0x8e, 0x52, 0xf8, 0x8a, 0x6a, 0xbc, 0x24,
					0x24, 0x15, 0x16, 0x05, 0x7a, 0xea, 0x41, 0x06,
					0xad, 0xd7, 0xab, 0x7b, 0xb9, 0xd5, 0xe6, 0xa0,
					0x8b, 0xe5, 0xb2, 0x12, 0x09, 0x44, 0x1d, 0x74,
				}
				return util.NewAddressPublicKey(publicKey, util.Bech32PrefixKaspa)
			},
			passedPrefix:   util.Bech32PrefixUnknown,
			expectedPrefix: util.Bech32PrefixKaspa,
		},
		{
			name:    "mainnet p2pk 2",
			addr:    "zork:qq80qvqs0lfxuzmt7sz3909ze6camq9d4t35ennsep3hxfe7ln35cn3lpjcaw",
			encoded: "zork:qq80qvqs0lfxuzmt7sz3909ze6camq9d4t35ennsep3hxfe7ln35cn3lpjcaw",
			valid:   true,
			result: util.TstAddressPubKey(
				util.Bech32PrefixKaspa,
				[util.PublicKeySize]byte{
					0x0e, 0xf0, 0x30, 0x10, 0x7f, 0xd2, 0x6e, 0x0b, 0x6b, 0xf4,
					0x05, 0x12, 0xbc, 0xa2, 0xce, 0xb1, 0xdd, 0x80, 0xad, 0xaa,
					0xe3, 0x4c, 0xce, 0x70, 0xc8, 0x63, 0x73, 0x27, 0x3e, 0xfc,
					0xe3, 0x4c,
				}),
			f: func() (util.Address, error) {
				publicKey := []byte{
					0x0e, 0xf0, 0x30, 0x10, 0x7f, 0xd2, 0x6e, 0x0b, 0x6b, 0xf4,
					0x05, 0x12, 0xbc, 0xa2, 0xce, 0xb1, 0xdd, 0x80, 0xad, 0xaa,
					0xe3, 0x4c, 0xce, 0x70, 0xc8, 0x63, 0x73, 0x27, 0x3e, 0xfc,
					0xe3, 0x4c,
				}
				return util.NewAddressPublicKey(publicKey, util.Bech32PrefixKaspa)
			},
			passedPrefix:   util.Bech32PrefixKaspa,
			expectedPrefix: util.Bech32PrefixKaspa,
		},
		{
			name:    "testnet p2pk",
			addr:    "zorktest:qputx94qseratdmjs0j395mq8u03er0x3l35ennsep3hxfe7ln35cn9gka0d6",
			encoded: "zorktest:qputx94qseratdmjs0j395mq8u03er0x3l35ennsep3hxfe7ln35cn9gka0d6",
			valid:   true,
			result: util.TstAddressPubKey(
				util.Bech32PrefixKaspaTest,
				[util.PublicKeySize]byte{
					0x78, 0xb3, 0x16, 0xa0, 0x86, 0x47, 0xd5, 0xb7, 0x72, 0x83,
					0xe5, 0x12, 0xd3, 0x60, 0x3f, 0x1f, 0x1c, 0x8d, 0xe6, 0x8f,
					0xe3, 0x4c, 0xce, 0x70, 0xc8, 0x63, 0x73, 0x27, 0x3e, 0xfc,
					0xe3, 0x4c,
				}),
			f: func() (util.Address, error) {
				publicKey := []byte{
					0x78, 0xb3, 0x16, 0xa0, 0x86, 0x47, 0xd5, 0xb7, 0x72, 0x83,
					0xe5, 0x12, 0xd3, 0x60, 0x3f, 0x1f, 0x1c, 0x8d, 0xe6, 0x8f,
					0xe3, 0x4c, 0xce, 0x70, 0xc8, 0x63, 0x73, 0x27, 0x3e, 0xfc,
					0xe3, 0x4c,
				}
				return util.NewAddressPublicKey(publicKey, util.Bech32PrefixKaspaTest)
			},
			passedPrefix:   util.Bech32PrefixKaspaTest,
			expectedPrefix: util.Bech32PrefixKaspaTest,
		},

		// ECDSA P2PK tests.
		{
			name:    "mainnet ecdsa p2pk",
			addr:    "zork:qyp0r5mcq4rd5grj3652ra09u5dcgwqq9ntuswp247nama5quyj40eqnpwn4yz2",
			encoded: "zork:qyp0r5mcq4rd5grj3652ra09u5dcgwqq9ntuswp247nama5quyj40eqnpwn4yz2",
			valid:   true,
			result: util.TstAddressPubKeyECDSA(
				util.Bech32PrefixKaspa,
				[util.PublicKeySizeECDSA]byte{
					0x02, 0xf1, 0xd3, 0x78, 0x05, 0x46, 0xda, 0x20, 0x72, 0x8e, 0xa8, 0xa1, 0xf5, 0xe5, 0xe5, 0x1b, 0x84, 0x38, 0x00, 0x2c, 0xd7, 0xc8, 0x38, 0x2a, 0xaf, 0xa7, 0xdd, 0xf6, 0x80, 0xe1, 0x25, 0x57, 0xe4,
				}),
			f: func() (util.Address, error) {
				publicKey := []byte{
					0x02, 0xf1, 0xd3, 0x78, 0x05, 0x46, 0xda, 0x20, 0x72, 0x8e, 0xa8, 0xa1, 0xf5, 0xe5, 0xe5, 0x1b, 0x84, 0x38, 0x00, 0x2c, 0xd7, 0xc8, 0x38, 0x2a, 0xaf, 0xa7, 0xdd, 0xf6, 0x80, 0xe1, 0x25, 0x57, 0xe4,
				}
				return util.NewAddressPublicKeyECDSA(publicKey, util.Bech32PrefixKaspa)
			},
			passedPrefix:   util.Bech32PrefixUnknown,
			expectedPrefix: util.Bech32PrefixKaspa,
		},

		// Negative P2PK tests.
		{
			name:  "p2pk wrong public key length",
			addr:  "",
			valid: false,
			f: func() (util.Address, error) {
				publicKey := []byte{
					0x00, 0x0e, 0xf0, 0x30, 0x10, 0x7f, 0xd2, 0x6e, 0x0b, 0x6b,
					0xf4, 0x05, 0x12, 0xbc, 0xa2, 0xce, 0xb1, 0xdd, 0x80, 0xad,
					0xaa}
				return util.NewAddressPublicKey(publicKey, util.Bech32PrefixKaspa)
			},
			passedPrefix:   util.Bech32PrefixKaspa,
			expectedPrefix: util.Bech32PrefixKaspa,
		},
		{
			name:           "p2pk bad checksum",
			addr:           "zork:qr35ennsep3hxfe7lnz5ee7j5jgmkjswss74as46gx",
			valid:          false,
			passedPrefix:   util.Bech32PrefixKaspa,
			expectedPrefix: util.Bech32PrefixKaspa,
		},

		// Positive P2SH tests.
		{
			name:    "mainnet p2sh",
			addr:    "zork:prq20q4qd9ulr044cauyy9wtpeupqpjv67pn2vyc6acly7xqkrjdzy0c2zvku",
			encoded: "zork:prq20q4qd9ulr044cauyy9wtpeupqpjv67pn2vyc6acly7xqkrjdzy0c2zvku",
			valid:   true,
			result: util.TstAddressScriptHash(
				util.Bech32PrefixKaspa,
				[blake2b.Size256]byte{
					0xc0, 0xa7, 0x82, 0xa0, 0x69, 0x79, 0xf1, 0xbe,
					0xb5, 0xc7, 0x78, 0x42, 0x15, 0xcb, 0x0e, 0x78,
					0x10, 0x06, 0x4c, 0xd7, 0x83, 0x35, 0x30, 0x98,
					0xd7, 0x71, 0xf2, 0x78, 0xc0, 0xb0, 0xe4, 0xd1,
				}),
			f: func() (util.Address, error) {
				script := []byte{
					0x52, 0x41, 0x04, 0x91, 0xbb, 0xa2, 0x51, 0x09, 0x12, 0xa5,
					0xbd, 0x37, 0xda, 0x1f, 0xb5, 0xb1, 0x67, 0x30, 0x10, 0xe4,
					0x3d, 0x2c, 0x6d, 0x81, 0x2c, 0x51, 0x4e, 0x91, 0xbf, 0xa9,
					0xf2, 0xeb, 0x12, 0x9e, 0x1c, 0x18, 0x33, 0x29, 0xdb, 0x55,
					0xbd, 0x86, 0x8e, 0x20, 0x9a, 0xac, 0x2f, 0xbc, 0x02, 0xcb,
					0x33, 0xd9, 0x8f, 0xe7, 0x4b, 0xf2, 0x3f, 0x0c, 0x23, 0x5d,
					0x61, 0x26, 0xb1, 0xd8, 0x33, 0x4f, 0x86, 0x41, 0x04, 0x86,
					0x5c, 0x40, 0x29, 0x3a, 0x68, 0x0c, 0xb9, 0xc0, 0x20, 0xe7,
					0xb1, 0xe1, 0x06, 0xd8, 0xc1, 0x91, 0x6d, 0x3c, 0xef, 0x99,
					0xaa, 0x43, 0x1a, 0x56, 0xd2, 0x53, 0xe6, 0x92, 0x56, 0xda,
					0xc0, 0x9e, 0xf1, 0x22, 0xb1, 0xa9, 0x86, 0x81, 0x8a, 0x7c,
					0xb6, 0x24, 0x53, 0x2f, 0x06, 0x2c, 0x1d, 0x1f, 0x87, 0x22,
					0x08, 0x48, 0x61, 0xc5, 0xc3, 0x29, 0x1c, 0xcf, 0xfe, 0xf4,
					0xec, 0x68, 0x74, 0x41, 0x04, 0x8d, 0x24, 0x55, 0xd2, 0x40,
					0x3e, 0x08, 0x70, 0x8f, 0xc1, 0xf5, 0x56, 0x00, 0x2f, 0x1b,
					0x6c, 0xd8, 0x3f, 0x99, 0x2d, 0x08, 0x50, 0x97, 0xf9, 0x97,
					0x4a, 0xb0, 0x8a, 0x28, 0x83, 0x8f, 0x07, 0x89, 0x6f, 0xba,
					0xb0, 0x8f, 0x39, 0x49, 0x5e, 0x15, 0xfa, 0x6f, 0xad, 0x6e,
					0xdb, 0xfb, 0x1e, 0x75, 0x4e, 0x35, 0xfa, 0x1c, 0x78, 0x44,
					0xc4, 0x1f, 0x32, 0x2a, 0x18, 0x63, 0xd4, 0x62, 0x13, 0x53,
					0xae}
				return util.NewAddressScriptHash(script, util.Bech32PrefixKaspa)
			},
			passedPrefix:   util.Bech32PrefixKaspa,
			expectedPrefix: util.Bech32PrefixKaspa,
		},
		{
			name:    "mainnet p2sh 2",
			addr:    "zork:pzgqclhvzgt9dt9er7a7vvd7ske0njgdl88hmjrd85czpnut6n2qv58p5czjp",
			encoded: "zork:pzgqclhvzgt9dt9er7a7vvd7ske0njgdl88hmjrd85czpnut6n2qv58p5czjp",
			valid:   true,
			result: util.TstAddressScriptHash(
				util.Bech32PrefixKaspa,
				[blake2b.Size256]byte{
					//900c7eec121656acb91fbbe631be85b2f9c90df9cf7dc86d3d3020cf8bd4d406
					0x90, 0x0c, 0x7e, 0xec, 0x12, 0x16, 0x56, 0xac,
					0xb9, 0x1f, 0xbb, 0xe6, 0x31, 0xbe, 0x85, 0xb2,
					0xf9, 0xc9, 0x0d, 0xf9, 0xcf, 0x7d, 0xc8, 0x6d,
					0x3d, 0x30, 0x20, 0xcf, 0x8b, 0xd4, 0xd4, 0x06,
				}),
			f: func() (util.Address, error) {
				hash := []byte{
					0x90, 0x0c, 0x7e, 0xec, 0x12, 0x16, 0x56, 0xac,
					0xb9, 0x1f, 0xbb, 0xe6, 0x31, 0xbe, 0x85, 0xb2,
					0xf9, 0xc9, 0x0d, 0xf9, 0xcf, 0x7d, 0xc8, 0x6d,
					0x3d, 0x30, 0x20, 0xcf, 0x8b, 0xd4, 0xd4, 0x06,
				}
				return util.NewAddressScriptHashFromHash(hash, util.Bech32PrefixKaspa)
			},
			passedPrefix:   util.Bech32PrefixKaspa,
			expectedPrefix: util.Bech32PrefixKaspa,
		},
		{
			name:    "testnet p2sh",
			addr:    "zorktest:pp9y9l5ztx2vt26yw40t6g55p3t5arfh6dm9n0qf2kha6wttumj6k0q0qjmfj",
			encoded: "zorktest:pp9y9l5ztx2vt26yw40t6g55p3t5arfh6dm9n0qf2kha6wttumj6k0q0qjmfj",
			valid:   true,
			result: util.TstAddressScriptHash(
				util.Bech32PrefixKaspaTest,
				[blake2b.Size256]byte{
					//4a42fe825994c5ab44755ebd22940c574e8d37d37659bc0955afdd396be6e5ab
					0x4a, 0x42, 0xfe, 0x82, 0x59, 0x94, 0xc5, 0xab,
					0x44, 0x75, 0x5e, 0xbd, 0x22, 0x94, 0x0c, 0x57,
					0x4e, 0x8d, 0x37, 0xd3, 0x76, 0x59, 0xbc, 0x09,
					0x55, 0xaf, 0xdd, 0x39, 0x6b, 0xe6, 0xe5, 0xab,
				}),
			f: func() (util.Address, error) {
				hash := []byte{
					0x4a, 0x42, 0xfe, 0x82, 0x59, 0x94, 0xc5, 0xab,
					0x44, 0x75, 0x5e, 0xbd, 0x22, 0x94, 0x0c, 0x57,
					0x4e, 0x8d, 0x37, 0xd3, 0x76, 0x59, 0xbc, 0x09,
					0x55, 0xaf, 0xdd, 0x39, 0x6b, 0xe6, 0xe5, 0xab,
				}
				return util.NewAddressScriptHashFromHash(hash, util.Bech32PrefixKaspaTest)
			},
			passedPrefix:   util.Bech32PrefixKaspaTest,
			expectedPrefix: util.Bech32PrefixKaspaTest,
		},

		// Negative P2SH tests.
		{
			name:  "p2sh wrong hash length",
			addr:  "",
			valid: false,
			f: func() (util.Address, error) {
				hash := []byte{
					0x00, 0xf8, 0x15, 0xb0, 0x36, 0xd9, 0xbb, 0xbc,
					0xe5, 0xe9, 0xf2, 0xa0, 0x0a, 0xbd, 0x1b, 0xf3,
					0xdc, 0x91, 0xe9, 0x55, 0x10}
				return util.NewAddressScriptHashFromHash(hash, util.Bech32PrefixKaspa)
			},
			passedPrefix:   util.Bech32PrefixKaspa,
			expectedPrefix: util.Bech32PrefixKaspa,
		},
	}

	for _, test := range tests {
		// Decode addr and compare error against valid.
		decoded, err := util.DecodeAddress(test.addr, test.passedPrefix)
		if (err == nil) != test.valid {
			t.Errorf("%v: decoding test failed: %v", test.name, err)
			return
		}

		if err == nil {
			// Ensure the stringer returns the same address as the
			// original.
			if decodedStringer, ok := decoded.(fmt.Stringer); ok {
				addr := test.addr

				if addr != decodedStringer.String() {
					t.Errorf("%v: String on decoded value does not match expected value: %v != %v",
						test.name, test.addr, decodedStringer.String())
					return
				}
			}

			// Encode again and compare against the original.
			encoded := decoded.EncodeAddress()
			if test.encoded != encoded {
				t.Errorf("%v: decoding and encoding produced different addressess: %v != %v",
					test.name, test.encoded, encoded)
				return
			}

			// Perform type-specific calculations.
			var saddr []byte
			switch decoded.(type) {
			case *util.AddressPublicKey:
				saddr = util.TstAddressSAddrP2PK(encoded)

			case *util.AddressPublicKeyECDSA:
				saddr = util.TstAddressSAddrP2PKECDSA(encoded)

			case *util.AddressScriptHash:
				saddr = util.TstAddressSAddrP2SH(encoded)
			}

			// Check script address, as well as the HashBlake2b method for P2SH addresses.
			if !bytes.Equal(saddr, decoded.ScriptAddress()) {
				t.Errorf("%v: script addresses do not match:\n%x != \n%x",
					test.name, saddr, decoded.ScriptAddress())
				return
			}
			switch a := decoded.(type) {
			case *util.AddressPublicKey:
				if h := a.ScriptAddress()[:]; !bytes.Equal(saddr, h) {
					t.Errorf("%v: hashes do not match:\n%x != \n%x",
						test.name, saddr, h)
					return
				}

			case *util.AddressScriptHash:
				if h := a.HashBlake2b()[:]; !bytes.Equal(saddr, h) {
					t.Errorf("%v: hashes do not match:\n%x != \n%x",
						test.name, saddr, h)
					return
				}
			}

			// Ensure the address is for the expected network.
			if !decoded.IsForPrefix(test.expectedPrefix) {
				t.Errorf("%v: calculated network does not match expected",
					test.name)
				return
			}
		}

		if !test.valid {
			// If address is invalid, but a creation function exists,
			// verify that it returns a nil addr and non-nil error.
			if test.f != nil {
				_, err := test.f()
				if err == nil {
					t.Errorf("%v: address is invalid but creating new address succeeded",
						test.name)
					return
				}
			}
			continue
		}

		// Valid test, compare address created with f against expected result.
		addr, err := test.f()
		if err != nil {
			t.Errorf("%v: address is valid but creating new address failed with error %v",
				test.name, err)
			return
		}

		if !reflect.DeepEqual(addr, test.result) {
			t.Errorf("%v: created address does not match expected result",
				test.name)
			return
		}

		if !reflect.DeepEqual(addr, decoded) {
			t.Errorf("%v: created address does not match the decoded address",
				test.name)
			return
		}

		if !reflect.DeepEqual(addr, decoded) {
			t.Errorf("%v: created address does not match the decoded address",
				test.name)
			return
		}
	}
}

func TestDecodeAddressErrorConditions(t *testing.T) {
	tests := []struct {
		address      string
		prefix       util.Bech32Prefix
		errorMessage string
	}{
		{
			"bitcoincash:qpzry9x8gf2tvdw0s3jn54khce6mua7lcw20ayyn",
			util.Bech32PrefixUnknown,
			"decoded address's prefix could not be parsed",
		},
		{
			"zorksim:raskzctpv9skzctpv9skzctpv9skzctpvy37ct7zafpv9skzctpvyrghfh3nm",
			util.Bech32PrefixKaspaSim,
			"unknown address type",
		},
		{
			"zorksim:raskzcggd400kmx",
			util.Bech32PrefixKaspaSim,
			"unknown address type",
		},
		{
			"zorktest:qpxl0r7v5xcux7zqpn02r50z7g9a7xa89t76ff62j9eq445etege5aqcflzq0",
			util.Bech32PrefixKaspa,
			"decoded address is of wrong network",
		},
	}

	for _, test := range tests {
		_, err := util.DecodeAddress(test.address, test.prefix)
		if err == nil {
			t.Errorf("decodeAddress unexpectedly succeeded")
		} else if !strings.Contains(err.Error(), test.errorMessage) {
			t.Errorf("received mismatched error. Expected '%s' but got '%s'",
				test.errorMessage, err)
		}
	}
}

func TestParsePrefix(t *testing.T) {
	tests := []struct {
		prefixStr      string
		expectedPrefix util.Bech32Prefix
		expectedError  bool
	}{
		{"zork", util.Bech32PrefixKaspa, false},
		{"zorktest", util.Bech32PrefixKaspaTest, false},
		{"zorksim", util.Bech32PrefixKaspaSim, false},
		{"blabla", util.Bech32PrefixUnknown, true},
		{"unknown", util.Bech32PrefixUnknown, true},
		{"", util.Bech32PrefixUnknown, true},
	}

	for _, test := range tests {
		result, err := util.ParsePrefix(test.prefixStr)
		if (err != nil) != test.expectedError {
			t.Errorf("TestParsePrefix: %s: expected error status: %t, but got %t",
				test.prefixStr, test.expectedError, err != nil)
		}

		if result != test.expectedPrefix {
			t.Errorf("TestParsePrefix: %s: expected prefix: %d, but got %d",
				test.prefixStr, test.expectedPrefix, result)
		}
	}
}

func TestPrefixToString(t *testing.T) {
	tests := []struct {
		prefix            util.Bech32Prefix
		expectedPrefixStr string
	}{
		{util.Bech32PrefixKaspa, "zork"},
		{util.Bech32PrefixKaspaTest, "zorktest"},
		{util.Bech32PrefixKaspaSim, "zorksim"},
		{util.Bech32PrefixUnknown, ""},
	}

	for _, test := range tests {
		result := test.prefix.String()

		if result != test.expectedPrefixStr {
			t.Errorf("TestPrefixToString: %s: expected string: %s, but got %s",
				test.prefix, test.expectedPrefixStr, result)
		}
	}
}
