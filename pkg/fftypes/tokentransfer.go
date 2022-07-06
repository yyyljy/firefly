// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fftypes

type TokenTransferType = FFEnum

var (
	TokenTransferTypeMint     = ffEnum("tokentransfertype", "mint")
	TokenTransferTypeBurn     = ffEnum("tokentransfertype", "burn")
	TokenTransferTypeTransfer = ffEnum("tokentransfertype", "transfer")
)

type TokenTransfer struct {
	Type            TokenTransferType `ffstruct:"TokenTransfer" json:"type" ffenum:"tokentransfertype" ffexcludeinput:"true"`
	LocalID         *UUID             `ffstruct:"TokenTransfer" json:"localId,omitempty" ffexcludeinput:"true"`
	Pool            *UUID             `ffstruct:"TokenTransfer" json:"pool,omitempty"`
	TokenIndex      string            `ffstruct:"TokenTransfer" json:"tokenIndex,omitempty"`
	URI             string            `ffstruct:"TokenTransfer" json:"uri,omitempty"`
	Connector       string            `ffstruct:"TokenTransfer" json:"connector,omitempty" ffexcludeinput:"true"`
	Namespace       string            `ffstruct:"TokenTransfer" json:"namespace,omitempty" ffexcludeinput:"true"`
	Key             string            `ffstruct:"TokenTransfer" json:"key,omitempty"`
	From            string            `ffstruct:"TokenTransfer" json:"from,omitempty" ffexcludeinput:"postTokenMint"`
	To              string            `ffstruct:"TokenTransfer" json:"to,omitempty" ffexcludeinput:"postTokenBurn"`
	Amount          FFBigInt          `ffstruct:"TokenTransfer" json:"amount"`
	ProtocolID      string            `ffstruct:"TokenTransfer" json:"protocolId,omitempty" ffexcludeinput:"true"`
	Message         *UUID             `ffstruct:"TokenTransfer" json:"message,omitempty"`
	MessageHash     *Bytes32          `ffstruct:"TokenTransfer" json:"messageHash,omitempty" ffexcludeinput:"true"`
	Created         *FFTime           `ffstruct:"TokenTransfer" json:"created,omitempty" ffexcludeinput:"true"`
	TX              TransactionRef    `ffstruct:"TokenTransfer" json:"tx" ffexcludeinput:"true"`
	BlockchainEvent *UUID             `ffstruct:"TokenTransfer" json:"blockchainEvent,omitempty" ffexcludeinput:"true"`
	Config          JSONObject        `ffstruct:"TokenTransfer" json:"config,omitempty" ffexcludeoutput:"true"` // for REST calls only (not stored)
}

type TokenTransferInput struct {
	TokenTransfer
	Message *MessageInOut `ffstruct:"TokenTransferInput" json:"message,omitempty"`
	Pool    string        `ffstruct:"TokenTransferInput" json:"pool,omitempty"`
}
