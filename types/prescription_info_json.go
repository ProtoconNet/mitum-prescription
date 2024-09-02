package types

import (
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/ProtoconNet/mitum2/util/hint"
)

type PrescriptionInfoJSONMarshaler struct {
	hint.BaseHinter
	PrescriptionHash string `json:"prescription_hash"`
	PrescribeDate    uint64 `json:"prescribe_date"`
	PrepareDate      uint64 `json:"prepare_date"`
	EndDate          uint64 `json:"end_date"`
	Status           string `json:"status"`
	Hospital         string `json:"hospital"`
	Pharmacy         string `json:"pharmacy"`
}

func (p PrescriptionInfo) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(PrescriptionInfoJSONMarshaler{
		BaseHinter:       p.BaseHinter,
		PrescriptionHash: p.prescriptionHash,
		PrescribeDate:    p.prescribeDate,
		PrepareDate:      p.prepareDate,
		EndDate:          p.endDate,
		Status:           p.status.String(),
		Hospital:         p.hospital,
		Pharmacy:         p.pharmacy,
	})
}

type PrescriptionInfoJSONUnmarshaler struct {
	Hint             hint.Hint `json:"_hint"`
	PrescriptionHash string    `json:"prescription_hash"`
	PrescribeDate    uint64    `json:"prescribe_date"`
	PrepareDate      uint64    `json:"prepare_date"`
	EndDate          uint64    `json:"end_date"`
	Status           string    `json:"status"`
	Hospital         string    `json:"hospital"`
	Pharmacy         string    `json:"pharmacy"`
}

func (p *PrescriptionInfo) DecodeJSON(b []byte, enc encoder.Encoder) error {
	e := util.StringError("failed to decode json of PrescriptionInfo")

	var u PrescriptionInfoJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	return p.encode(u.Hint, u.PrescriptionHash, u.PrescribeDate, u.PrepareDate, u.EndDate, u.Status, u.Hospital, u.Pharmacy)
}
