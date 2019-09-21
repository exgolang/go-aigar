// Copyright 2019 The go-aigar Authors
// This file is part of the go-aigar library.
//
// The go-aigar library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-aigar library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-aigar library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Aigar Foundation Go Bootnodes
	"enode://59a9ca420a9cdab5cc35b58797ce59a6756725d23e34273fce11050206e19e0f47e7c3e1ec0c47a36b8310b15ccf9ce40796a1a7e6b0f1074cbaceaa3fe8fd77@178.32.112.221:30301",  // London
	"enode://e648a63eab88bfc158539c1a4aa61b6615791d3191e0023ea6b0b2fdbce0193c318af3a453a9d1c3ffcd8900b1364ccc55d974cd3389d1389a51e05220578c47@185.227.108.100:30301", // Frankfurt
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes []string

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes []string

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes []string

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes []string
