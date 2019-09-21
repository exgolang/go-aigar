// Copyright 2016 The go-aigar Authors
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

import (
	"math/big"

	"github.com/exgolang/go-aigar/common"
)

// DAOForkBlockExtra is the block header extra-data field to set for the DAO fork
// point and a number of consecutive blocks to allow fast/light syncers to correctly
// pick the side they want  ("dao-hard-fork").
var DAOForkBlockExtra = common.FromHex("0x64616f2d686172642d666f726b")

// DAOForkExtraRange is the number of consecutive blocks from the DAO fork point
// to override the extra-data in to prevent no-fork attacks.
var DAOForkExtraRange = big.NewInt(10)

// DAORefundContract is the address of the refund contract to send DAO balances to.
var DAORefundContract = common.HexToAddress("0x1E4e3461e0873B6F6d21dAcc51705c4Aecb0A9FD")

// DAODrainList is the list of accounts whose full balances will be moved into a
// refund contract at the beginning of the dao-fork block.
func DAODrainList() []common.Address {
	return []common.Address{
		common.HexToAddress("0xb1080cDd0fF5B0dE78F4B5B470A6b0F4232FC377"),
		common.HexToAddress("0xB7331D9E194A45a45b375fDB9F856FD48551b6eF"),
		common.HexToAddress("0xA3dD8140cb29cAbEB03a8E47254a57656e35fBA7"),
		common.HexToAddress("0xACA35cB8d3dC454816eA2103F6e49532AefA84c0"),
		common.HexToAddress("0x670c89806Cf9C1f1Da9C022977c9dF5eA3baBd17"),
		common.HexToAddress("0xEe3C23dDE5F90441AA1008D8391fc6CDf63F17b2"),
		common.HexToAddress("0x8A4b509FA755419a4962f73D5cAdcA84Fe9e4cb6"),
		common.HexToAddress("0x55f4ac7a4d29928c367035486b87E7E67547c900"),
		common.HexToAddress("0xa52BAF602CDEd73585f17DDF06b621bD0351d691"),
		common.HexToAddress("0x7677079725f513e84917FCBa6b4a32b25B40A26F"),
		common.HexToAddress("0x2eeDC61F5841a3734481708eb4c847d35D184233"),
		common.HexToAddress("0xe93C7fF346822A4284ac2487b09BfF50C5E9da49"),
		common.HexToAddress("0x7702dE477B1120C84C3f7ea87280e0EC5f430dFC"),
		common.HexToAddress("0x485c8C02227cAFDB92f8d4762aA2185652eE5138"),
		common.HexToAddress("0x2D67e63B2bFC0Ce6DDD5C087e046F276323f7Ea9"),
		common.HexToAddress("0xD89D21F9b40c0D91B5112ecbB297FA8EF5ae7042"),
		common.HexToAddress("0x6F7151ec4E934A79A932A159A4d554C8ca8Cb406"),
		common.HexToAddress("0xd3Ba9236767BE744b79EE8145Ed103cF0af46f50"),
		common.HexToAddress("0xE81DD44242CD32E1b3a24D1654F890cdBaCB48F3"),
		common.HexToAddress("0x3a0f2a30d6636afB1E80ce20220C316A6B4473d3"),
		common.HexToAddress("0x960499062081E0723465B451FE4F824A59940b66"),
		common.HexToAddress("0x9bb9167D35AA42Bb814AC7946A138Cb7f2Bebb8f"),
		common.HexToAddress("0x60eB76BE9a9b5c342708891AE23231448DB174B0"),
		common.HexToAddress("0x5E1B7158e29eAC2bDb032170Ad6B3288a640fD41"),
		common.HexToAddress("0xc59Cf2C2E2D142CeD9765B2cc31D871287198DB1"),
		common.HexToAddress("0x7ABfD5C857ee8242c29C8939052baCA93Abb5179"),
	}
}
