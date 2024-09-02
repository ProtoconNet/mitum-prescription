package types

import "github.com/ProtoconNet/mitum2/util"

type PrescriptionStatus uint8

const (
	Null PrescriptionStatus = iota
	Registered
	Used
)

var PrescriptionStatusUnmarshaller = map[string]PrescriptionStatus{
	"Null":       Null,
	"Registered": Registered,
	"Used":       Used,
}

func (p PrescriptionStatus) String() string {
	switch p {
	case Registered:
		return "Registered"
	case Used:
		return "Used"
	default:
		return "Null"
	}
}

func (p PrescriptionStatus) Bytes() []byte {
	return util.Uint8ToBytes(uint8(8))
}
