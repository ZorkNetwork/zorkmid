// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"math/big"

	"github.com/kaspanet/go-muhash"
	"github.com/kaspanet/kaspad/domain/consensus/model/externalapi"
	"github.com/kaspanet/kaspad/domain/consensus/utils/blockheader"
	"github.com/kaspanet/kaspad/domain/consensus/utils/subnetworks"
	"github.com/kaspanet/kaspad/domain/consensus/utils/transactionhelper"
)

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, //script version
	0x01, // Varint
	0x00, // OP-FALSE
	0x49, 0x6E, 0x20, 0x46, 0x72, 0x6F, 0x62, 0x73,
	0x20, 0x57, 0x65, 0x20, 0x54, 0x72, 0x75, 0x73,
	0x74, 0x0A, // In Frobs We Trust
	// 092B 094D 0930 0949 092C 094D 0938 0020 092E 0947 0902 0020 0939 092E 0020 092D 0930 094B 0938 093E 0020 0915 0930 0924 0947 0020 0939 0948 0902 000A // hindi
	// 041C 044B 0020 0434 043E 0432 0435 0440 044F 0435 043C 0020 0424 0440 043E 0431 0441 0443 000A // russian
	// 6211 5011 4FE1 8CF4 0020 0046 006C 006F 0062 0073 000A  // traditional chinese
	// 0641 064A 0020 0641 0631 0648 0628 0632 0020 0646 062D 0646 0020 0646 062B 0642 // arabic
	// 0046 0072 006F 0062 0073 0020 0645 06CC 06BA 0020 06C1 0645 0020 0627 0639 062A 0645 0627 062F 0020 06A9 0631 062A 06D2 0020 06C1 06CC 06BA 000A // urdu

	0xc7, 0x81, 0x58, 0x9f, 0xf0, 0x83, 0xf7, 0x2f, // Kaspa block hash c781589ff083f72f33bc7f07d41e9b9a92806a2d02a14abfefd12ca4e4d5bdf2
	0x33, 0xbc, 0x7f, 0x07, 0xd4, 0x1e, 0x9b, 0x9a,
	0x92, 0x80, 0x6a, 0x2d, 0x02, 0xa1, 0x4a, 0xbf,
	0xef, 0xd1, 0x2c, 0xa4, 0xe4, 0xd5, 0xbd, 0xf2,
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0, []*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	//0d4078ff9033c71dd6683a61e5ac49355898af7df68843c7ed872ddb14a2ebc6
	0x0d, 0x40, 0x78, 0xff, 0x90, 0x33, 0xc7, 0x1d,
	0xd6, 0x68, 0x3a, 0x61, 0xe5, 0xac, 0x49, 0x35,
	0x58, 0x98, 0xaf, 0x7d, 0xf6, 0x88, 0x43, 0xc7,
	0xed, 0x87, 0x2d, 0xdb, 0x14, 0xa2, 0xeb, 0xc6,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	//e8b0598ab7bd4ebee332d35ad7a28e41051d7f89aefc25b86a7f3e5f38b81886
	0xe8, 0xb0, 0x59, 0x8a, 0xb7, 0xbd, 0x4e, 0xbe,
	0xe3, 0x32, 0xd3, 0x5a, 0xd7, 0xa2, 0x8e, 0x41,
	0x05, 0x1d, 0x7f, 0x89, 0xae, 0xfc, 0x25, 0xb8,
	0x6a, 0x7f, 0x3e, 0x5f, 0x38, 0xb8, 0x18, 0x86,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
			0x71, 0x0f, 0x27, 0xdf, 0x42, 0x3e, 0x63, 0xaa, 0x6c, 0xdb, 0x72, 0xb8, 0x9e, 0xa5, 0xa0, 0x6c, 0xff, 0xa3, 0x99, 0xd6, 0x6f, 0x16, 0x77, 0x04, 0x45, 0x5b, 0x5a, 0xf5, 0x9d, 0xef, 0x8e, 0x20,
		}),
		1637609671037,
		486722099,
		0x3392c,
		1312860, // Checkpoint DAA score
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                             // Varint
	0x00,                                                             // OP-FALSE
	0x7A, 0x6F, 0x72, 0x6B, 0x2D, 0x64, 0x65, 0x76, 0x6E, 0x65, 0x74, // zork-devnet
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	//7c8a1a9b1325fdccc1e78246e49f26bd3ab5034c2d2817e918bdc6fab44d3723
	0x7c, 0x8a, 0x1a, 0x9b, 0x13, 0x25, 0xfd, 0xcc,
	0xc1, 0xe7, 0x82, 0x46, 0xe4, 0x9f, 0x26, 0xbd,
	0x3a, 0xb5, 0x03, 0x4c, 0x2d, 0x28, 0x17, 0xe9,
	0x18, 0xbd, 0xc6, 0xfa, 0xb4, 0x4d, 0x37, 0x23,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	//a1c54d7c2ffb3665c88010272af5329abe4b89db74ff8aae15a8cb253b65b875
	0xa1, 0xc5, 0x4d, 0x7c, 0x2f, 0xfb, 0x36, 0x65,
	0xc8, 0x80, 0x10, 0x27, 0x2a, 0xf5, 0x32, 0x9a,
	0xbe, 0x4b, 0x89, 0xdb, 0x74, 0xff, 0x8a, 0xae,
	0x15, 0xa8, 0xcb, 0x25, 0x3b, 0x65, 0xb8, 0x75,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x11e9db49828,
		525264379,
		0x48e5e,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                             // Varint
	0x00,                                                             // OP-FALSE
	0x7A, 0x6F, 0x72, 0x6B, 0x2D, 0x73, 0x69, 0x6D, 0x6E, 0x65, 0x74, // zork-simnet
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	//c15a6ff2685edf3032e1b3e36bb9d74dd8e273743e97d93f9f889b85e248a3e1
	0xc1, 0x5a, 0x6f, 0xf2, 0x68, 0x5e, 0xdf, 0x30,
	0x32, 0xe1, 0xb3, 0xe3, 0x6b, 0xb9, 0xd7, 0x4d,
	0xd8, 0xe2, 0x73, 0x74, 0x3e, 0x97, 0xd9, 0x3f,
	0x9f, 0x88, 0x9b, 0x85, 0xe2, 0x48, 0xa3, 0xe1,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	//9db8689f02de17829c4ab4a7fc0d9a2b591a3245f67ee51422bab8780cf245a
	0x9d, 0xb8, 0x68, 0x9f, 0x02, 0xde, 0x17, 0x82,
	0x9c, 0x4a, 0xb4, 0xa7, 0xfc, 0x0d, 0x9a, 0x2b,
	0x59, 0x1a, 0x32, 0x45, 0xf6, 0x7e, 0xe5, 0x11,
	0x42, 0x2b, 0xab, 0x87, 0x80, 0xcf, 0x24, 0x5a,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x17c5f62fbb6,
		0x207fffff,
		0x2,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x7A, 0x6F, 0x72, 0x6B, 0x2D, 0x74, 0x65, 0x73, 0x74, 0x6E, 0x65, 0x74, // zork-testnet
	0x49, 0x6E, 0x20, 0x46, 0x72, 0x6F, 0x62, 0x73,
	0x20, 0x57, 0x65, 0x20, 0x54, 0x72, 0x75, 0x73,
	0x74, 0x0A, // In Frobs We Trust
	0xc7, 0x81, 0x58, 0x9f, 0xf0, 0x83, 0xf7, 0x2f, // Kaspa block hash c781589ff083f72f33bc7f07d41e9b9a92806a2d02a14abfefd12ca4e4d5bdf2
	0x33, 0xbc, 0x7f, 0x07, 0xd4, 0x1e, 0x9b, 0x9a,
	0x92, 0x80, 0x6a, 0x2d, 0x02, 0xa1, 0x4a, 0xbf,
	0xef, 0xd1, 0x2c, 0xa4, 0xe4, 0xd5, 0xbd, 0xf2,
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	//07f354606c98d573621bdafa7f265e88fb02a122ce9c8754c965f52431f9c787
	0x07, 0xf3, 0x54, 0x60, 0x6c, 0x98, 0xd5, 0x73,
	0x62, 0x1b, 0xda, 0xfa, 0x7f, 0x26, 0x5e, 0x88,
	0xfb, 0x02, 0xa1, 0x22, 0xce, 0x9c, 0x87, 0x54,
	0xc9, 0x65, 0xf5, 0x24, 0x31, 0xf9, 0xc7, 0x87,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	//d785b5e7f0dfe96b491a8a25353bab5460ca638aa0afd5d79c4a1d4e8f1e48c8
	0xd7, 0x85, 0xb5, 0xe7, 0xf0, 0xdf, 0xe9, 0x6b,
	0x49, 0x1a, 0x8a, 0x25, 0x35, 0x3b, 0xab, 0x54,
	0x60, 0xca, 0x63, 0x8a, 0xa0, 0xaf, 0xd5, 0xd7,
	0x9c, 0x4a, 0x1d, 0x4e, 0x8f, 0x1e, 0x48, 0xc8,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x17c5f62fbb6,
		0x1e7fffff,
		0x14582,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
