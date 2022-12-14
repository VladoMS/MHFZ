package mhfpacket

import ( 
 "errors" 

 	"erupe-ce/network/clientctx"
	"erupe-ce/network"
	"erupe-ce/common/byteframe"
)

// MsgMhfUpdateWarehouse represents the MSG_MHF_UPDATE_WAREHOUSE
type MsgMhfUpdateWarehouse struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateWarehouse) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_WAREHOUSE
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateWarehouse) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateWarehouse) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
