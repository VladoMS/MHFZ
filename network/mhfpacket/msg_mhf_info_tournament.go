package mhfpacket

import ( 
 "errors" 

 	"erupe-ce/network/clientctx"
	"erupe-ce/network"
	"erupe-ce/common/byteframe"
)

// MsgMhfInfoTournament represents the MSG_MHF_INFO_TOURNAMENT
type MsgMhfInfoTournament struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfInfoTournament) Opcode() network.PacketID {
	return network.MSG_MHF_INFO_TOURNAMENT
}

// Parse parses the packet from binary
func (m *MsgMhfInfoTournament) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfInfoTournament) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
