package mhfpacket

import ( 
 "errors" 

 	"erupe-ce/network/clientctx"
	"erupe-ce/network"
	"erupe-ce/common/byteframe"
)

// MsgMhfOperationInvGuild represents the MSG_MHF_OPERATION_INV_GUILD
type MsgMhfOperationInvGuild struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfOperationInvGuild) Opcode() network.PacketID {
	return network.MSG_MHF_OPERATION_INV_GUILD
}

// Parse parses the packet from binary
func (m *MsgMhfOperationInvGuild) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfOperationInvGuild) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
