// Copyright (c) 2017-2021 Ivan Jelincic <parazyd@dyne.org>
//
// This file is part of tordam
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package tordam

import (
	"crypto/ed25519"
)

// Peer is the base struct for any peer in the network.
type Peer struct {
	Pubkey     ed25519.PublicKey `json:"pubkey"`     // Peer's ed25519 public key
	Portmap    []string          `json:"portmap"`    // Peer's port map in Tor
	Nonce      string            `json:"nonce"`      // The nonce to be signed after announce init
	SelfRevoke string            `json:"selfrevoke"` // Our revoke key we use to update our data
	PeerRevoke string            `json:"peerrevoke"` // Peer's revoke key if they wish to update their data
	LastSeen   int64             `json:"lastseen"`   // Timestamp of last announce
	Trusted    int               `json:"trusted"`    // Trusted is int because of possible levels of trust
}
