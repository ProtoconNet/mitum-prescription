package types

import (
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/pkg/errors"
)

func (p *PrescriptionInfo) encode(
	ht hint.Hint,
	prescriptionHash string,
	prescribeDate, prepareDate, endDate uint64,
	status, hospital, pharmacy string,
) error {
	p.BaseHinter = hint.NewBaseHinter(ht)
	p.prescriptionHash = prescriptionHash
	p.prescribeDate = prescribeDate
	p.prepareDate = prepareDate
	p.endDate = endDate
	v, found := PrescriptionStatusUnmarshaller[status]
	if !found {
		return errors.Errorf("failed to unmarshall prescription status")
	}
	p.status = v
	p.hospital = hospital
	p.pharmacy = pharmacy

	return nil
}
